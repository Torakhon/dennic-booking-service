package usecase

import (
	_ "booking_service/genproto/booking_service"
	appointment "booking_service/internal/entity/booked_appointments"
	"context"
)

// BookedAppointmentsUseCase -.
type BookedAppointmentsUseCase struct {
	repo BookedAppointments
}

// NewBookedAppointments -.
func NewBookedAppointments(r BookedAppointments) *BookedAppointmentsUseCase {
	return &BookedAppointmentsUseCase{
		repo: r,
	}
}

func (r *BookedAppointmentsUseCase) CreateAppointment(ctx context.Context, req *appointment.CreateAppointment) (*appointment.Appointment, error) {

	return r.repo.CreateAppointment(ctx, req)
}

func (r *BookedAppointmentsUseCase) GetAppointment(ctx context.Context, req *appointment.FieldValueReq) (*appointment.Appointment, error) {
	return r.repo.GetAppointment(ctx, req)
}

func (r *BookedAppointmentsUseCase) GetAllAppointment(ctx context.Context, req *appointment.GetAllAppointment) (*appointment.AppointmentsType, error) {
	return r.repo.GetAllAppointment(ctx, req)
}

func (r *BookedAppointmentsUseCase) UpdateAppointment(ctx context.Context, req *appointment.UpdateAppointment) (*appointment.Appointment, error) {
	return r.repo.UpdateAppointment(ctx, req)
}

func (r *BookedAppointmentsUseCase) DeleteAppointment(ctx context.Context, req *appointment.FieldValueReq) (*appointment.StatusRes, error) {
	return r.repo.DeleteAppointment(ctx, req)
}
