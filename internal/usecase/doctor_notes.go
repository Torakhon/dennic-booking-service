package usecase

import (
	"booking_service/internal/entity/doctor_notes"
	"context"
)

// BookedDoctorNotesUseCase -.
type BookedDoctorNotesUseCase struct {
	Repo DoctorNotes
}

// NewBookedDoctorNotes -.
func NewBookedDoctorNotes(r DoctorNotes) *BookedDoctorNotesUseCase {
	return &BookedDoctorNotesUseCase{
		Repo: r,
	}
}

func (r *BookedDoctorNotesUseCase) CreateDoctorNotes(ctx context.Context, req *doctor_notes.CreatedDoctorNote) (*doctor_notes.DoctorNote, error) {
	return r.Repo.CreateDoctorNotes(ctx, req)
}

func (r *BookedDoctorNotesUseCase) GetDoctorNotes(ctx context.Context, req *doctor_notes.FieldValueReq) (*doctor_notes.DoctorNote, error) {
	return r.Repo.GetDoctorNotes(ctx, req)
}

func (r *BookedDoctorNotesUseCase) GetAllDoctorNotes(ctx context.Context, req *doctor_notes.GetAllNotes) (*doctor_notes.DoctorNotesType, error) {
	return r.Repo.GetAllDoctorNotes(ctx, req)
}

func (r *BookedDoctorNotesUseCase) UpdateDoctorNotes(ctx context.Context, req *doctor_notes.UpdateDoctorNoteReq) (*doctor_notes.DoctorNote, error) {
	return r.Repo.UpdateDoctorNotes(ctx, req)
}

func (r *BookedDoctorNotesUseCase) DeleteDoctorNotes(ctx context.Context, req *doctor_notes.FieldValueReq) (*doctor_notes.StatusRes, error) {
	return r.Repo.DeleteDoctorNotes(ctx, req)
}
