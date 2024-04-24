package usecase

import (
	"booking_service/internal/entity/patients"
	"context"
)

// BookedPatientUseCase -.
type BookedPatientUseCase struct {
	Repo Patient
}

// NewBookedPatient -.
func NewBookedPatient(r Patient) *BookedPatientUseCase {
	return &BookedPatientUseCase{
		Repo: r,
	}
}

func (r *BookedPatientUseCase) CreatePatient(ctx context.Context, req *patients.CreatedPatient) (*patients.Patient, error) {
	return r.Repo.CreatePatient(ctx, req)
}

func (r *BookedPatientUseCase) GetPatient(ctx context.Context, req *patients.FieldValueReq) (*patients.Patient, error) {
	return r.Repo.GetPatient(ctx, req)
}

func (r *BookedPatientUseCase) GetAllPatiens(ctx context.Context, req *patients.GetAllPatients) (*patients.PatientsType, error) {
	return r.Repo.GetAllPatiens(ctx, req)
}

func (r *BookedPatientUseCase) UpdatePatient(ctx context.Context, req *patients.UpdatePatient) (*patients.Patient, error) {
	return r.Repo.UpdatePatient(ctx, req)
}

func (r *BookedPatientUseCase) UpdatePhonePatient(ctx context.Context, req *patients.UpdatePhoneNumber) (*patients.StatusRes, error) {
	return r.Repo.UpdatePhonePatient(ctx, req)
}

func (r *BookedPatientUseCase) DeletePatient(ctx context.Context, req *patients.FieldValueReq) (*patients.StatusRes, error) {
	return r.Repo.DeletePatient(ctx, req)
}
