package repo

import (
	appointment "booking_service/internal/entity/booked_appointments"
	"booking_service/internal/pkg/postgres"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type BookingAppointment struct {
	db *postgres.PostgresDB
}

func NewBookingAppointment(db *postgres.PostgresDB) *BookingAppointment {
	return &BookingAppointment{
		db: db,
	}
}

const tableNameAppointment = "booked_appointments"

func tableColums() string {
	return `id, 
			department_id, 
			doctor_id, 
			patient_id, 
			appointment_date, 
			appointment_time, 
			duration, 
			key, 
			expires_at, 
			patient_status, 
			created_at, 
			updated_at, 
			deleted_at`
}

func (r *BookingAppointment) CreateAppointment(ctx context.Context, req *appointment.CreateAppointment) (*appointment.Appointment, error) {
	var (
		response appointment.Appointment
		upAt     sql.NullTime
		delAt    sql.NullTime
	)
	toSql, args, err := r.db.Sq.Builder.
		Insert(tableNameAppointment).
		Columns(` 
			department_id, 
			doctor_id, 
			patient_id, 
			appointment_date, 
			appointment_time, 
			duration, 
			key, 
			expires_at, 
			patient_status`).
		Values(
			req.DepartmentId,
			req.DoctorId,
			req.PatientId,
			req.AppointmentDate.String(),
			req.AppointmentTime,
			req.Duration,
			req.Key,
			req.ExpiresAt,
			req.PatientStatus).
		Suffix(fmt.Sprintf("RETURNING %s", tableColums())).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err = r.db.QueryRow(ctx, toSql, args...).Scan(
		&response.Id,
		&response.DepartmentId,
		&response.DoctorId,
		&response.PatientId,
		&response.AppointmentDate,
		&response.AppointmentTime,
		&response.Duration,
		&response.Key,
		&response.ExpiresAt,
		&response.PatientStatus,
		&response.CreatedAt,
		&upAt,
		&delAt,
	); err != nil {
		return nil, err
	}

	if upAt.Valid {
		response.UpdatedAt = upAt.Time
	}

	if delAt.Valid {
		response.DeletedAt = delAt.Time
	}

	return &response, nil
}

func (r *BookingAppointment) GetAppointment(ctx context.Context, req *appointment.FieldValueReq) (*appointment.Appointment, error) {
	var (
		response appointment.Appointment
		upAt     sql.NullTime
		delAt    sql.NullTime
	)

	toSql := r.db.Sq.Builder.
		Select(tableColums()).
		From(tableNameAppointment)

	if !req.DeleteStatus {
		toSql = toSql.Where(r.db.Sq.Equal("deleted_at", nil))
	}
	toSqls, args, err := toSql.Where(r.db.Sq.Equal(req.Field, req.Value)).ToSql()

	if err != nil {
		return nil, err
	}
	if err = r.db.QueryRow(ctx, toSqls, args...).Scan(
		&response.Id,
		&response.DepartmentId,
		&response.DoctorId,
		&response.PatientId,
		&response.AppointmentDate,
		&response.AppointmentTime,
		&response.Duration,
		&response.Key,
		&response.ExpiresAt,
		&response.PatientStatus,
		&response.CreatedAt,
		&upAt,
		&delAt,
	); err != nil {
		return nil, err
	}

	if upAt.Valid {
		response.UpdatedAt = upAt.Time
	}

	if delAt.Valid {
		response.DeletedAt = delAt.Time
	}

	return &response, nil
}

func (r *BookingAppointment) GetAllAppointment(ctx context.Context, req *appointment.GetAllAppointment) (*appointment.AppointmentsType, error) {
	var (
		response appointment.AppointmentsType
		upAt     sql.NullTime
		delAt    sql.NullTime
	)

	toSql := r.db.Sq.Builder.
		Select(tableColums()).
		From(tableNameAppointment)

	if req.Page >= 1 && req.Limit >= 1 {
		toSql = toSql.
			Limit(req.Limit).
			Offset(req.Limit * (req.Page - 1))
	}

	if req.Value != "" {
		toSql = toSql.Where(r.db.Sq.ILike(req.Field, req.Value+"%"))
	}
	if req.OrderBy != "" {
		toSql = toSql.OrderBy(req.OrderBy)
	}
	if !req.DeleteStatus {
		toSql = toSql.Where(r.db.Sq.Equal("deleted_at", nil))
	}

	toSqls, args, err := toSql.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, toSqls, args...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var res appointment.Appointment
		if err := rows.Scan(
			&res.Id,
			&res.DepartmentId,
			&res.DoctorId,
			&res.PatientId,
			&res.AppointmentDate,
			&res.AppointmentTime,
			&res.Duration,
			&res.Key,
			&res.ExpiresAt,
			&res.PatientStatus,
			&res.CreatedAt,
			&upAt,
			&delAt,
		); err != nil {
			return nil, err
		}

		if upAt.Valid {
			res.UpdatedAt = upAt.Time
		}

		if delAt.Valid {
			res.DeletedAt = delAt.Time
		}

		response.Count += 1

		response.Appointments = append(response.Appointments, &res)
	}
	return &response, nil
}

func (r *BookingAppointment) UpdateAppointment(ctx context.Context, req *appointment.UpdateAppointment) (*appointment.Appointment, error) {
	var (
		response appointment.Appointment
		upAt     sql.NullTime
		delAt    sql.NullTime
	)
	toSql, args, err := r.db.Sq.Builder.
		Update(tableNameAppointment).
		SetMap(map[string]interface{}{
			"appointment_date": req.AppointmentDate.String(),
			"appointment_time": req.AppointmentTime,
			"duration":         req.Duration,
			"key":              req.Key,
			"expires_at":       req.ExpiresAt,
			"patient_status":   req.PatientStatus,
			"updated_at":       time.Now(),
		}).
		Where(r.db.Sq.Equal(req.Field, req.Value)).
		Suffix(fmt.Sprintf("RETURNING %s", tableColums())).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err = r.db.QueryRow(ctx, toSql, args...).Scan(
		&response.Id,
		&response.DepartmentId,
		&response.DoctorId,
		&response.PatientId,
		&response.AppointmentDate,
		&response.AppointmentTime,
		&response.Duration,
		&response.Key,
		&response.ExpiresAt,
		&response.PatientStatus,
		&response.CreatedAt,
		&upAt,
		&delAt,
	); err != nil {
		return nil, err
	}

	if upAt.Valid {
		response.UpdatedAt = upAt.Time
	}

	if delAt.Valid {
		response.DeletedAt = delAt.Time
	}

	return &response, nil
}

func (r *BookingAppointment) DeleteAppointment(ctx context.Context, req *appointment.FieldValueReq) (*appointment.StatusRes, error) {
	if !req.DeleteStatus {
		toSql, args, err := r.db.Sq.Builder.
			Update(tableNameAppointment).
			Set("deleted_at", time.Now()).
			Where(r.db.Sq.EqualMany(map[string]interface{}{
				"deleted_at": nil,
				req.Field:    req.Value,
			})).
			ToSql()
		if err != nil {
			return &appointment.StatusRes{Status: false}, err
		}

		_, err = r.db.Exec(ctx, toSql, args...)

		if err != nil {
			return &appointment.StatusRes{Status: false}, err
		}
		return &appointment.StatusRes{Status: true}, nil

	} else {
		toSql, args, err := r.db.Sq.Builder.
			Delete(tableNameAppointment).
			Where(r.db.Sq.Equal(req.Field, req.Value)).
			ToSql()

		if err != nil {
			return &appointment.StatusRes{Status: false}, err
		}

		_, err = r.db.Exec(ctx, toSql, args...)

		if err != nil {
			return &appointment.StatusRes{Status: false}, err
		}
		return &appointment.StatusRes{Status: true}, nil
	}
}
