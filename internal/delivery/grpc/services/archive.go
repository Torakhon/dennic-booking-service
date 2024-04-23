package services

import (
	pb "booking_service/genproto/booking_service"
	"booking_service/internal/entity/archive"
	"booking_service/internal/usecase"
	"context"
	"go.uber.org/zap"

	"time"
)

type BookingArchive struct {
	logger               *zap.Logger
	bookedArchiveUseCase usecase.Archive
}

func BookingArchiveNewRPC(
	logger *zap.Logger,
	ArchiveUsaCase usecase.Archive) *BookingArchive {

	return &BookingArchive{
		logger:               logger,
		bookedArchiveUseCase: ArchiveUsaCase,
	}
}

func (r *BookingArchive) CreateArchive(ctx context.Context, req *pb.CreatedArchive) (*pb.Archive, error) {
	startTime, err := time.Parse("15:04:05", req.StartTime)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse("15:04:05", req.EndTime)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedArchiveUseCase.CreateArchive(ctx, &archive.CreatedArchive{
		DoctorAvailabilityId: req.DoctorAvailabilityId,
		StartTime:            startTime,
		EndTime:              endTime,
		PatientProblem:       req.PatientProblem,
		Status:               req.Status,
		PaymentType:          req.PaymentType,
		PaymentAmount:        float64(req.PaymentAmount),
	})

	if err != nil {
		return nil, err
	}

	return &pb.Archive{
		Id:                   res.Id,
		DoctorAvailabilityId: res.DoctorAvailabilityId,
		StartTime:            res.StartTime.String(),
		EndTime:              res.EndTime.String(),
		PatientProblem:       res.PatientProblem,
		Status:               res.Status,
		PaymentType:          res.PaymentType,
		PaymentAmount:        float32(res.PaymentAmount),
		CreatedAt:            res.CreatedAt.String(),
		UpdatedAt:            res.UpdatedAt.String(),
		DeletedAt:            res.DeletedAt.String(),
	}, nil

}

func (r *BookingArchive) GetArchive(ctx context.Context, req *pb.FieldValueReq) (*pb.Archive, error) {
	res, err := r.bookedArchiveUseCase.GetArchive(ctx, &archive.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Archive{
		Id:                   res.Id,
		DoctorAvailabilityId: res.DoctorAvailabilityId,
		StartTime:            res.StartTime.String(),
		EndTime:              res.EndTime.String(),
		PatientProblem:       res.PatientProblem,
		Status:               res.Status,
		PaymentType:          res.PaymentType,
		PaymentAmount:        float32(res.PaymentAmount),
		CreatedAt:            res.CreatedAt.String(),
		UpdatedAt:            res.UpdatedAt.String(),
		DeletedAt:            res.DeletedAt.String(),
	}, nil
}

func (r *BookingArchive) GetAllArchives(ctx context.Context, req *pb.GetAllReq) (*pb.ArchivesType, error) {
	var archives pb.ArchivesType

	archivesRes, err := r.bookedArchiveUseCase.GetAllArchive(ctx, &archive.GetAllArchives{
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

	for _, archiveRes := range archivesRes.Archives {
		var archiv pb.Archive
		archiv.Id = archiveRes.Id
		archiv.DoctorAvailabilityId = archiveRes.DoctorAvailabilityId
		archiv.StartTime = archiveRes.StartTime.String()
		archiv.EndTime = archiveRes.EndTime.String()
		archiv.PatientProblem = archiveRes.PatientProblem
		archiv.Status = archiveRes.Status
		archiv.PaymentType = archiveRes.PaymentType
		archiv.PaymentAmount = float32(archiveRes.PaymentAmount)
		archiv.CreatedAt = archiveRes.CreatedAt.String()
		archiv.UpdatedAt = archiveRes.UpdatedAt.String()
		archiv.DeletedAt = archiveRes.DeletedAt.String()
		archives.Archives = append(archives.Archives, &archiv)
	}
	archives.Count = archivesRes.Count

	return &archives, nil
}

func (r *BookingArchive) UpdateArchive(ctx context.Context, req *pb.UpdateArchiveReq) (*pb.Archive, error) {
	startTime, err := time.Parse("15:04:05", req.StartTime)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse("15:04:05", req.EndTime)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedArchiveUseCase.UpdateArchive(ctx, &archive.UpdateArchive{
		Field:                req.Field,
		Value:                req.Value,
		DoctorAvailabilityId: req.DoctorAvailabilityId,
		StartTime:            startTime,
		EndTime:              endTime,
		PatientProblem:       req.PatientProblem,
		Status:               req.Status,
		PaymentType:          req.PaymentType,
		PaymentAmount:        float64(req.PaymentAmount),
	})

	if err != nil {
		return nil, err
	}

	return &pb.Archive{
		Id:                   res.Id,
		DoctorAvailabilityId: res.DoctorAvailabilityId,
		StartTime:            res.StartTime.String(),
		EndTime:              res.EndTime.String(),
		PatientProblem:       res.PatientProblem,
		Status:               res.Status,
		PaymentType:          res.PaymentType,
		PaymentAmount:        float32(res.PaymentAmount),
		CreatedAt:            res.CreatedAt.String(),
		UpdatedAt:            res.UpdatedAt.String(),
		DeletedAt:            res.DeletedAt.String(),
	}, nil
}

func (r *BookingArchive) DeleteArchive(ctx context.Context, req *pb.FieldValueReq) (*pb.StatusRes, error) {
	res, err := r.bookedArchiveUseCase.DeleteArchive(ctx, &archive.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.GetDeleteStatus(),
	})

	if err != nil {
		return &pb.StatusRes{Status: res.Status}, err
	}

	return &pb.StatusRes{Status: res.Status}, nil
}
