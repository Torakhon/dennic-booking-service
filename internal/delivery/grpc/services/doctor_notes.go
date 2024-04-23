package services

import (
	pb "booking_service/genproto/booking_service"
	"booking_service/internal/entity/doctor_notes"
	"booking_service/internal/usecase"
	"context"
	"go.uber.org/zap"
)

type BookingDoctorNotes struct {
	logger                          *zap.Logger
	bookedAppointmentUseCase        usecase.BookedAppointments
	bookedPatientUseCase            usecase.Patient
	bookedDoctorNotesUseCase        usecase.DoctorNotes
	bookedArchiveUseCase            usecase.Archive
	bookedDoctorAvailabilityUseCase usecase.DoctorAvailability
}

func BookingDoctorNotesNewRPC(
	logger *zap.Logger, DoctorNotesUseCase usecase.DoctorNotes) *BookingDoctorNotes {
	return &BookingDoctorNotes{
		logger:                   logger,
		bookedDoctorNotesUseCase: DoctorNotesUseCase,
	}
}

func (r *BookingDoctorNotes) CreateDoctorNote(ctx context.Context, req *pb.CreatedDoctorNote) (*pb.DoctorNote, error) {
	res, err := r.bookedDoctorNotesUseCase.CreateDoctorNotes(ctx, &doctor_notes.CreatedDoctorNote{
		AppointmentId: req.AppointmentId,
		DoctorId:      req.DoctorId,
		PatientId:     req.PatientId,
		Prescription:  req.Prescription,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DoctorNote{
		Id:            res.Id,
		AppointmentId: res.AppointmentId,
		DoctorId:      res.DoctorId,
		PatientId:     res.PatientId,
		Prescription:  res.Prescription,
		CreatedAt:     res.CreatedAt.String(),
		UpdatedAt:     res.UpdatedAt.String(),
		DeletedAt:     res.DeletedAt.String(),
	}, nil
}

func (r *BookingDoctorNotes) GetDoctorNote(ctx context.Context, req *pb.FieldValueReq) (*pb.DoctorNote, error) {
	res, err := r.bookedDoctorNotesUseCase.GetDoctorNotes(ctx, &doctor_notes.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DoctorNote{
		Id:            res.Id,
		AppointmentId: res.AppointmentId,
		DoctorId:      res.DoctorId,
		PatientId:     res.PatientId,
		Prescription:  res.Prescription,
		CreatedAt:     res.CreatedAt.String(),
		UpdatedAt:     res.UpdatedAt.String(),
		DeletedAt:     res.DeletedAt.String(),
	}, nil
}

func (r *BookingDoctorNotes) GetAllNotes(ctx context.Context, req *pb.GetAllReq) (*pb.DoctorNotesType, error) {
	var notesRes pb.DoctorNotesType

	allNotes, err := r.bookedDoctorNotesUseCase.GetAllDoctorNotes(ctx, &doctor_notes.GetAllNotes{
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

	for _, note := range allNotes.DoctorNotes {
		var noteRes pb.DoctorNote
		noteRes.Id = note.Id
		noteRes.AppointmentId = note.AppointmentId
		noteRes.DoctorId = note.DoctorId
		noteRes.PatientId = note.PatientId
		noteRes.Prescription = note.Prescription
		noteRes.CreatedAt = note.CreatedAt.String()
		noteRes.UpdatedAt = note.UpdatedAt.String()
		noteRes.DeletedAt = note.DeletedAt.String()
		notesRes.DoctorNotes = append(notesRes.DoctorNotes, &noteRes)
	}
	notesRes.Count = allNotes.Count

	return &notesRes, nil
}

func (r *BookingDoctorNotes) UpdateDoctorNote(ctx context.Context, req *pb.UpdateDoctorNoteReq) (*pb.DoctorNote, error) {
	res, err := r.bookedDoctorNotesUseCase.UpdateDoctorNotes(ctx, &doctor_notes.UpdateDoctorNoteReq{
		Field:         req.Field,
		Value:         req.Value,
		AppointmentId: req.AppointmentId,
		DoctorId:      req.DoctorId,
		PatientId:     req.PatientId,
		Prescription:  req.Prescription,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DoctorNote{
		Id:            res.Id,
		AppointmentId: res.AppointmentId,
		DoctorId:      res.DoctorId,
		PatientId:     res.PatientId,
		Prescription:  res.Prescription,
		CreatedAt:     res.CreatedAt.String(),
		UpdatedAt:     res.UpdatedAt.String(),
		DeletedAt:     res.DeletedAt.String(),
	}, nil
}

func (r *BookingDoctorNotes) DeleteDoctorNote(ctx context.Context, req *pb.FieldValueReq) (*pb.StatusRes, error) {
	res, err := r.bookedDoctorNotesUseCase.DeleteDoctorNotes(ctx, &doctor_notes.FieldValueReq{
		Field:        req.Field,
		Value:        req.Value,
		DeleteStatus: req.DeleteStatus,
	})

	if err != nil {
		return &pb.StatusRes{Status: res.Status}, err
	}

	return &pb.StatusRes{Status: res.Status}, err
}
