package services

import (
	pb "booking_service/genproto/booking_service"
	_ "booking_service/internal/delivery/grpc"
	appointment "booking_service/internal/entity/booked_appointments"
	_ "booking_service/internal/pkg/otlp"
	"booking_service/internal/usecase"
	"context"
	date "github.com/rickb777/date"
	_ "go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"time"
)

type BookingAppointments struct {
	logger                   *zap.Logger
	bookedAppointmentUseCase usecase.BookedAppointments
}

func BookingAppointmentsNewRPC(logger *zap.Logger, AppointmentUsaCase usecase.BookedAppointments) *BookingAppointments {

	return &BookingAppointments{
		logger:                   logger,
		bookedAppointmentUseCase: AppointmentUsaCase,
	}
}

func (r *BookingAppointments) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentReq) (*pb.Appointment, error) {
	Date, err := date.AutoParse(req.AppointmentDate)
	if err != nil {
		return nil, err
	}
	Time, err := time.Parse("15:04:05", req.AppointmentTime)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedAppointmentUseCase.CreateAppointment(ctx, &appointment.CreateAppointment{
		DepartmentId:    req.DepartmentId,
		DoctorId:        req.DoctorId,
		PatientId:       req.PatientId,
		AppointmentDate: Date,
		AppointmentTime: Time,
		Duration:        req.Duration,
		Key:             req.Key,
		ExpiresAt:       req.ExpiresAt,
		PatientStatus:   req.PatientStatus,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Appointment{
		Id:              res.Id,
		DepartmentId:    res.DepartmentId,
		DoctorId:        res.DoctorId,
		PatientId:       res.PatientId,
		AppointmentDate: res.AppointmentDate.String(),
		AppointmentTime: res.AppointmentTime.String(),
		Duration:        res.Duration,
		Key:             res.Key,
		ExpiresAt:       res.ExpiresAt,
		PatientStatus:   res.PatientStatus,
		CreatedAt:       res.CreatedAt.String(),
		UpdatedAt:       res.UpdatedAt.String(),
		DeletedAt:       res.DeletedAt.String(),
	}, nil
}

func (r *BookingAppointments) GetAppointment(ctx context.Context, req *pb.FieldValueReq) (*pb.Appointment, error) {
	res, err := r.bookedAppointmentUseCase.GetAppointment(ctx, &appointment.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Appointment{
		Id:              res.Id,
		DepartmentId:    res.DepartmentId,
		DoctorId:        res.DoctorId,
		PatientId:       res.PatientId,
		AppointmentDate: res.AppointmentDate.String(),
		AppointmentTime: res.AppointmentTime.String(),
		Duration:        res.Duration,
		Key:             res.Key,
		ExpiresAt:       res.ExpiresAt,
		PatientStatus:   res.PatientStatus,
		CreatedAt:       res.CreatedAt.String(),
		UpdatedAt:       res.UpdatedAt.String(),
		DeletedAt:       res.DeletedAt.String(),
	}, nil
}

func (r *BookingAppointments) GetAllAppointment(ctx context.Context, req *pb.GetAllReq) (*pb.AppointmentsType, error) {
	var appointmentsRes pb.AppointmentsType

	allAppointment, err := r.bookedAppointmentUseCase.GetAllAppointment(ctx, &appointment.GetAllAppointment{
		Page:         req.Page,
		Limit:        req.Limit,
		DeleteStatus: req.DeleteStatus,
		Field:        req.Field,
		Value:        req.Value,
		OrderBy:      req.OrderBy,
	})
	if err != nil {
		return nil, err
	}

	for _, appoint := range allAppointment.Appointments {
		var appointmentRes pb.Appointment
		appointmentRes.Id = appoint.Id
		appointmentRes.DepartmentId = appoint.DepartmentId
		appointmentRes.DoctorId = appoint.DoctorId
		appointmentRes.PatientId = appoint.PatientId
		appointmentRes.AppointmentDate = appoint.AppointmentDate.String()
		appointmentRes.AppointmentTime = appoint.AppointmentTime.String()
		appointmentRes.Duration = appoint.Duration
		appointmentRes.Key = appoint.Key
		appointmentRes.ExpiresAt = appoint.ExpiresAt
		appointmentRes.CreatedAt = appoint.CreatedAt.String()
		appointmentRes.UpdatedAt = appoint.UpdatedAt.String()
		appointmentRes.DeletedAt = appoint.DeletedAt.String()

		appointmentsRes.Appointments = append(appointmentsRes.Appointments, &appointmentRes)
	}
	appointmentsRes.Count = allAppointment.Count

	return &appointmentsRes, nil
}

func (r *BookingAppointments) UpdateAppointment(ctx context.Context, req *pb.UpdateAppointmentReq) (*pb.Appointment, error) {
	reqDate, err := date.AutoParse(req.AppointmentDate)
	if err != nil {
		return nil, err
	}

	reqTime, err := time.Parse("15:04:05", req.AppointmentTime)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedAppointmentUseCase.UpdateAppointment(ctx, &appointment.UpdateAppointment{
		Field:           req.Field,
		Value:           req.Value,
		AppointmentDate: reqDate,
		AppointmentTime: reqTime,
		Duration:        req.Duration,
		Key:             req.Key,
		ExpiresAt:       req.ExpiresAt,
		PatientStatus:   req.PatientStatus,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Appointment{
		Id:              res.Id,
		DepartmentId:    res.DepartmentId,
		DoctorId:        res.DoctorId,
		PatientId:       res.PatientId,
		AppointmentDate: res.AppointmentDate.String(),
		AppointmentTime: res.AppointmentTime.String(),
		Duration:        res.Duration,
		Key:             res.Key,
		ExpiresAt:       res.ExpiresAt,
		PatientStatus:   res.PatientStatus,
		CreatedAt:       res.CreatedAt.String(),
		UpdatedAt:       res.UpdatedAt.String(),
		DeletedAt:       res.DeletedAt.String(),
	}, nil
}

func (r *BookingAppointments) DeleteAppointment(ctx context.Context, req *pb.FieldValueReq) (*pb.StatusRes, error) {
	res, err := r.bookedAppointmentUseCase.DeleteAppointment(ctx, &appointment.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return &pb.StatusRes{Status: res.Status}, err
	}

	return &pb.StatusRes{Status: res.Status}, err
}
