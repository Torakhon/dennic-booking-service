package services

import (
	pb "booking_service/genproto/booking_service"
	"booking_service/internal/entity/patients"
	"booking_service/internal/usecase"
	"context"
	"github.com/rickb777/date"
	"go.uber.org/zap"
)

type BookingPatient struct {
	logger               *zap.Logger
	bookedPatientUseCase usecase.Patient
}

func BookingPatientNewRPC(logger *zap.Logger, PatientUsaCase usecase.Patient) *BookingPatient {
	return &BookingPatient{
		logger:               logger,
		bookedPatientUseCase: PatientUsaCase,
	}
}

func (r *BookingPatient) CreatePatient(ctx context.Context, req *pb.CreatedPatient) (*pb.Patient, error) {
	reqDate, err := date.AutoParse(req.BirthDate)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedPatientUseCase.CreatePatient(ctx, &patients.CreatedPatient{
		Id:             req.Id,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		BirthDate:      reqDate,
		Gender:         req.Gender,
		BloodGroup:     req.BloodGroup,
		PhoneNumber:    req.PhoneNumber,
		City:           req.City,
		Country:        req.Country,
		PatientProblem: req.PatientProblem,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Patient{
		Id:             res.Id,
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		BirthDate:      res.BirthDate.String(),
		Gender:         res.Gender,
		BloodGroup:     res.BloodGroup,
		PhoneNumber:    res.PhoneNumber,
		City:           res.City,
		Country:        res.Country,
		PatientProblem: res.PatientProblem,
		CreatedAt:      res.CreatedAt.String(),
		UpdatedAt:      res.UpdatedAt.String(),
		DeletedAt:      res.DeletedAt.String(),
	}, nil
}

func (r *BookingPatient) GetPatient(ctx context.Context, req *pb.FieldValueReq) (*pb.Patient, error) {
	res, err := r.bookedPatientUseCase.GetPatient(ctx, &patients.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Patient{
		Id:             res.Id,
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		BirthDate:      res.BirthDate.String(),
		Gender:         res.Gender,
		BloodGroup:     res.BloodGroup,
		PhoneNumber:    res.PhoneNumber,
		City:           res.City,
		Country:        res.Country,
		PatientProblem: res.PatientProblem,
		CreatedAt:      res.CreatedAt.String(),
		UpdatedAt:      res.UpdatedAt.String(),
		DeletedAt:      res.DeletedAt.String(),
	}, nil
}

func (r *BookingPatient) GetAllPatiens(ctx context.Context, req *pb.GetAllReq) (*pb.PatientsType, error) {
	var patentsRes pb.PatientsType
	allPatients, err := r.bookedPatientUseCase.GetAllPatiens(ctx, &patients.GetAllPatients{
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

	for _, patient := range allPatients.Patients {
		var patientRes pb.Patient
		patientRes.Id = patient.Id
		patientRes.FirstName = patient.FirstName
		patientRes.LastName = patient.LastName
		patientRes.BirthDate = patient.BirthDate.String()
		patientRes.Gender = patient.Gender
		patientRes.BloodGroup = patient.BloodGroup
		patientRes.PhoneNumber = patient.PhoneNumber
		patientRes.City = patient.City
		patientRes.Country = patient.Country
		patientRes.PatientProblem = patient.PatientProblem
		patientRes.CreatedAt = patient.CreatedAt.String()
		patientRes.UpdatedAt = patient.UpdatedAt.String()
		patientRes.DeletedAt = patient.DeletedAt.String()
		patentsRes.Patients = append(patentsRes.Patients, &patientRes)
	}
	patentsRes.Count = allPatients.Count

	return &patentsRes, nil
}

func (r *BookingPatient) UpdatePatient(ctx context.Context, req *pb.UpdatePatientReq) (*pb.Patient, error) {
	reqData, err := date.AutoParse(req.BirthDate)
	if err != nil {
		return nil, err
	}

	res, err := r.bookedPatientUseCase.UpdatePatient(ctx, &patients.UpdatePatient{
		Field:          req.Field,
		Value:          req.Value,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		BirthDate:      reqData,
		Gender:         req.Gender,
		BloodGroup:     req.BloodGroup,
		City:           req.City,
		Country:        req.Country,
		PatientProblem: req.PatientProblem,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Patient{
		Id:             res.Id,
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		BirthDate:      res.BirthDate.String(),
		Gender:         res.Gender,
		BloodGroup:     res.BloodGroup,
		PhoneNumber:    res.PhoneNumber,
		City:           res.City,
		Country:        res.Country,
		PatientProblem: res.PatientProblem,
		CreatedAt:      res.CreatedAt.String(),
		UpdatedAt:      res.UpdatedAt.String(),
		DeletedAt:      res.DeletedAt.String(),
	}, nil
}

func (r *BookingPatient) UpdatePhonePatient(ctx context.Context, req *pb.UpdatePhoneNumber) (*pb.StatusRes, error) {
	res, err := r.bookedPatientUseCase.UpdatePhonePatient(ctx, &patients.UpdatePhoneNumber{
		Field:       req.Field,
		Value:       req.Value,
		PhoneNumber: req.PhoneNumber,
	})

	if err != nil {
		return &pb.StatusRes{Status: res.Status}, err
	}

	return &pb.StatusRes{Status: res.Status}, nil
}

func (r *BookingPatient) DeletePatient(ctx context.Context, req *pb.FieldValueReq) (*pb.StatusRes, error) {
	res, err := r.bookedPatientUseCase.DeletePatient(ctx, &patients.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return &pb.StatusRes{Status: res.Status}, err
	}

	return &pb.StatusRes{Status: res.Status}, nil
}
