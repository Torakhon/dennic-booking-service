package app

import (
	pb "booking_service/genproto/booking_service"
	grpc_server "booking_service/internal/delivery/grpc/server"
	invest_grpc "booking_service/internal/delivery/grpc/services"
	"booking_service/internal/infrastructure/grpc_service_clients"
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	"booking_service/internal/pkg/logger"
	"booking_service/internal/pkg/postgres"
	"booking_service/internal/usecase"
	"fmt"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	Config         *config.Config
	Logger         *zap.Logger
	DB             *postgres.PostgresDB
	GrpcServer     *grpc.Server
	ShutdownOTLP   func() error
	ServiceClients grpc_service_clients.ServiceClients
}

func NewApp(cfg *config.Config) (*App, error) {
	// init l
	l, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		return nil, err
	}

	// otlp collector initialization
	//shutdownOTLP, err := otlp.InitOTLPProvider(cfg)
	//if err != nil {
	//	return nil, err
	//}

	// init db
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	// grpc server init
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(l),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_server.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(l),
				grpc_recovery.UnaryServerInterceptor(),
			),
			grpc_server.UnaryInterceptorData(l),
		)),
	)

	return &App{
		Config:     cfg,
		Logger:     l,
		DB:         db,
		GrpcServer: grpcServer,
	}, nil
}

func (a *App) Run() error {

	// context timeout initialization

	// Initialize Service Clients
	serviceClients, err := grpc_service_clients.New(a.Config)
	if err != nil {
		return fmt.Errorf("error during initialize service clients: %w", err)
	}
	a.ServiceClients = serviceClients

	// repositories initialization
	bookingAppointment := repo.NewBookingAppointment(a.DB)

	bookingPatients := repo.NewBookingPatients(a.DB)

	doctorNotes := repo.NewDoctorNotes(a.DB)

	bookingArchive := repo.NewBookingArchive(a.DB)

	doctorAvailability := repo.NewDoctorAvailability(a.DB)

	// usecase initialization
	appointmentsUseCase := usecase.BookedAppointments(bookingAppointment)

	patientUseCase := usecase.Patient(bookingPatients)

	doctorNotesUseCase := usecase.DoctorNotes(doctorNotes)

	archiveUseCase := usecase.Archive(bookingArchive)

	doctorAvailabilityUseCase := usecase.DoctorAvailability(doctorAvailability)

	pb.RegisterBookedAppointmentsServer(a.GrpcServer, invest_grpc.BookingAppointmentsNewRPC(a.Logger, appointmentsUseCase))

	pb.RegisterBookedPatientServer(a.GrpcServer, invest_grpc.BookingPatientNewRPC(a.Logger, patientUseCase))

	pb.RegisterBookedDoctorNotesServer(a.GrpcServer, invest_grpc.BookingDoctorNotesNewRPC(a.Logger, doctorNotesUseCase))

	pb.RegisterBookedArchiveServer(a.GrpcServer, invest_grpc.BookingArchiveNewRPC(a.Logger, archiveUseCase))

	pb.RegisterBookedDoctorAvailabilityServer(a.GrpcServer, invest_grpc.BookingDoctorAvailabilityNewRPC(a.Logger, doctorAvailabilityUseCase))
	a.Logger.Info("gRPC Server Listening", zap.String("url", a.Config.RPCPort))

	if err := grpc_server.Run(a.Config, a.GrpcServer); err != nil {
		return fmt.Errorf("gRPC fatal to serve grpc server over %s %w", a.Config.RPCPort, err)
	}

	return nil
}

func (a *App) Stop() {
	// close broker producer
	// closing client service connections
	a.ServiceClients.Close()
	// stop gRPC server
	a.GrpcServer.Stop()

	// database connection
	a.DB.Close()

	// shutdown otlp collector
	if err := a.ShutdownOTLP(); err != nil {
		a.Logger.Error("shutdown otlp collector", zap.Error(err))
	}

	// zap logger sync
	a.Logger.Sync()
}
