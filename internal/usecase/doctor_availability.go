package usecase

import (
	"booking_service/internal/entity/doctor_availability"
	"context"
)

// BookedDoctorAvailabilityUseCase -.
type BookedDoctorAvailabilityUseCase struct {
	Repo DoctorAvailability
}

// NewBookedDoctorAvailability -.
func NewBookedDoctorAvailability(r DoctorAvailability) *BookedDoctorAvailabilityUseCase {
	return &BookedDoctorAvailabilityUseCase{
		Repo: r,
	}
}

func (r *BookedDoctorAvailabilityUseCase) CreateDoctorAvailability(ctx context.Context, req *doctor_availability.CreateDoctorAvailability) (*doctor_availability.DoctorAvailability, error) {
	return r.Repo.CreateDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) GetDoctorAvailability(ctx context.Context, req *doctor_availability.FieldValueReq) (*doctor_availability.DoctorAvailability, error) {
	return r.Repo.GetDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) GetAllDoctorAvailability(ctx context.Context, req *doctor_availability.GetAllReq) (*doctor_availability.DoctorAvailabilityType, error) {
	return r.Repo.GetAllDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) UpdateDoctorAvailability(ctx context.Context, req *doctor_availability.UpdateDoctorAvailability) (*doctor_availability.DoctorAvailability, error) {
	return r.Repo.UpdateDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) DeleteDoctorAvailability(ctx context.Context, req *doctor_availability.FieldValueReq) (*doctor_availability.StatusRes, error) {
	return r.Repo.DeleteDoctorAvailability(ctx, req)
}
