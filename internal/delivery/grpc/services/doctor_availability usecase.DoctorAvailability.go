package services

import (
	pb "booking_service/genproto/booking_service"
	"booking_service/internal/entity/doctor_availability"
	"booking_service/internal/usecase"
	"context"
	"github.com/rickb777/date"
	"go.uber.org/zap"
	"time"
)

type BookingDoctorAvailability struct {
	logger                          *zap.Logger
	bookedDoctorAvailabilityUseCase usecase.DoctorAvailability
}

func BookingDoctorAvailabilityNewRPC(
	logger *zap.Logger, DoctorAvailability usecase.DoctorAvailability) *BookingDoctorAvailability {
	return &BookingDoctorAvailability{
		logger:                          logger,
		bookedDoctorAvailabilityUseCase: DoctorAvailability,
	}
}

func (r *BookingDoctorAvailability) CreateDoctorAvailability(ctx context.Context, req *pb.CreatedDoctorAvailability) (*pb.DoctorAvailability, error) {
	reqDate, err := date.AutoParse(req.DoctorDate)
	if err != nil {
		return nil, err
	}
	reqStartTime, err := time.Parse("15:04:05", req.StartTime)
	if err != nil {
		return nil, err
	}
	reqEndTime, err := time.Parse("15:04:05", req.EndTime)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedDoctorAvailabilityUseCase.CreateDoctorAvailability(ctx, &doctor_availability.CreateDoctorAvailability{
		DepartmentId: req.DepartmentId,
		DoctorId:     req.DoctorId,
		DoctorDate:   reqDate,
		StartTime:    reqStartTime,
		EndTime:      reqEndTime,
		Status:       req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DoctorAvailability{
		Id:           res.Id,
		DepartmentId: res.DepartmentId,
		DoctorId:     res.DoctorId,
		DoctorDate:   res.DoctorDate.String(),
		StartTime:    res.StartTime.String(),
		EndTime:      res.EndTime.String(),
		Status:       res.Status,
		CreatedAt:    res.CreatedAt.String(),
		UpdatedAt:    res.UpdatedAt.String(),
		DeletedAt:    res.DeletedAt.String(),
	}, nil
}

func (r *BookingDoctorAvailability) GetDoctorAvailability(ctx context.Context, req *pb.FieldValueReq) (*pb.DoctorAvailability, error) {
	res, err := r.bookedDoctorAvailabilityUseCase.GetDoctorAvailability(ctx, &doctor_availability.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DoctorAvailability{
		Id:           res.Id,
		DepartmentId: res.DepartmentId,
		DoctorId:     res.DoctorId,
		DoctorDate:   res.DoctorDate.String(),
		StartTime:    res.StartTime.String(),
		EndTime:      res.EndTime.String(),
		Status:       res.Status,
		CreatedAt:    res.CreatedAt.String(),
		UpdatedAt:    res.UpdatedAt.String(),
		DeletedAt:    res.DeletedAt.String(),
	}, nil
}

func (r *BookingDoctorAvailability) GetAllDoctorAvailabilitys(ctx context.Context, req *pb.GetAllReq) (*pb.DoctorAvailabilitysType, error) {
	var docAvails pb.DoctorAvailabilitysType

	allDocAvails, err := r.bookedDoctorAvailabilityUseCase.GetAllDoctorAvailability(ctx, &doctor_availability.GetAllReq{
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

	for _, availability := range allDocAvails.DoctorAvailabilitys {
		var docAvail pb.DoctorAvailability
		docAvail.Id = availability.Id
		docAvail.DepartmentId = availability.DepartmentId
		docAvail.DoctorId = availability.DoctorId
		docAvail.DoctorDate = availability.DoctorDate.String()
		docAvail.StartTime = availability.StartTime.String()
		docAvail.EndTime = availability.EndTime.String()
		docAvail.Status = availability.Status
		docAvail.CreatedAt = availability.CreatedAt.String()
		docAvail.UpdatedAt = availability.UpdatedAt.String()
		docAvail.DeletedAt = availability.DeletedAt.String()
		docAvails.Doctor_Availabilitys = append(docAvails.Doctor_Availabilitys, &docAvail)
	}

	docAvails.Count = allDocAvails.Count

	return &docAvails, nil
}

func (r *BookingDoctorAvailability) UpdateDoctorAvailability(ctx context.Context, req *pb.UpdateDoctorAvailabilityReq) (*pb.DoctorAvailability, error) {
	reqDate, err := date.AutoParse(req.DoctorDate)
	if err != nil {
		return nil, err
	}
	reqStartTime, err := time.Parse("15:04:05", req.StartTime)
	if err != nil {
		return nil, err
	}
	reqEndTime, err := time.Parse("15:04:05", req.EndTime)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedDoctorAvailabilityUseCase.UpdateDoctorAvailability(ctx, &doctor_availability.UpdateDoctorAvailability{
		Field:        req.Field,
		Value:        req.Value,
		DepartmentId: req.DepartmentId,
		DoctorId:     req.DoctorId,
		DoctorDate:   reqDate,
		StartTime:    reqStartTime,
		EndTime:      reqEndTime,
		Status:       req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DoctorAvailability{
		Id:           res.Id,
		DepartmentId: res.DepartmentId,
		DoctorId:     res.DoctorId,
		DoctorDate:   res.DoctorDate.String(),
		StartTime:    res.StartTime.String(),
		EndTime:      res.EndTime.String(),
		Status:       res.Status,
		CreatedAt:    res.CreatedAt.String(),
		UpdatedAt:    res.UpdatedAt.String(),
		DeletedAt:    res.DeletedAt.String(),
	}, nil
}

func (r *BookingDoctorAvailability) DeleteDoctorAvailability(ctx context.Context, req *pb.FieldValueReq) (*pb.StatusRes, error) {
	res, err := r.bookedDoctorAvailabilityUseCase.DeleteDoctorAvailability(ctx, &doctor_availability.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return &pb.StatusRes{Status: res.Status}, err
	}

	return &pb.StatusRes{Status: res.Status}, err
}
