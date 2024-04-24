package repo

import (
	"booking_service/internal/entity/doctor_notes"
	"booking_service/internal/pkg/postgres"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DoctorNotes struct {
	db *postgres.PostgresDB
}

func NewDoctorNotes(db *postgres.PostgresDB) *DoctorNotes {
	return &DoctorNotes{
		db: db,
	}
}

const tableNameDoctorNotes = "doctor_notes"

func tableColumNotes() string {
	return `id,
			appointment_id,
			doctor_id,
			patient_id,
			prescription,
			created_at,
			updated_at,
			deleted_at`
}

func (r *DoctorNotes) CreateDoctorNotes(ctx context.Context, req *doctor_notes.CreatedDoctorNote) (*doctor_notes.DoctorNote, error) {
	var (
		note    doctor_notes.DoctorNote
		upTime  sql.NullTime
		delTime sql.NullTime
	)
	toSql, args, err := r.db.Sq.Builder.
		Insert(tableNameDoctorNotes).
		Columns(`appointment_id,
						doctor_id,
						patient_id,
						prescription`).
		Values(
			req.AppointmentId,
			req.DoctorId,
			req.PatientId,
			req.Prescription,
		).
		Suffix(fmt.Sprintf("RETURNING %s", tableColumNotes())).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err = r.db.QueryRow(ctx, toSql, args...).Scan(
		&note.Id,
		&note.AppointmentId,
		&note.DoctorId,
		&note.PatientId,
		&note.Prescription,
		&note.CreatedAt,
		&upTime,
		&delTime,
	); err != nil {
		return nil, err
	}

	if upTime.Valid {
		note.UpdatedAt = upTime.Time
	}

	if delTime.Valid {
		note.DeletedAt = delTime.Time
	}

	return &note, nil
}

func (r *DoctorNotes) GetDoctorNotes(ctx context.Context, req *doctor_notes.FieldValueReq) (*doctor_notes.DoctorNote, error) {
	var (
		note    doctor_notes.DoctorNote
		upTime  sql.NullTime
		delTime sql.NullTime
	)

	toSql := r.db.Sq.Builder.
		Select(tableColumNotes()).
		From(tableNameDoctorNotes).
		Where(r.db.Sq.Equal(req.Field, req.Value))

	if !req.DeleteStatus {
		toSql = toSql.Where(r.db.Sq.Equal("deleted_at", nil))
	}

	toSqls, args, err := toSql.ToSql()

	if err != nil {
		return nil, err
	}

	if err = r.db.QueryRow(ctx, toSqls, args...).Scan(
		&note.Id,
		&note.AppointmentId,
		&note.DoctorId,
		&note.PatientId,
		&note.Prescription,
		&note.CreatedAt,
		&upTime,
		&delTime,
	); err != nil {
		return nil, err
	}

	if upTime.Valid {
		note.UpdatedAt = upTime.Time
	}

	if delTime.Valid {
		note.DeletedAt = delTime.Time
	}

	return &note, nil
}

func (r *DoctorNotes) GetAllDoctorNotes(ctx context.Context, req *doctor_notes.GetAllNotes) (*doctor_notes.DoctorNotesType, error) {
	var (
		notes   doctor_notes.DoctorNotesType
		upTime  sql.NullTime
		delTime sql.NullTime
	)

	toSql := r.db.Sq.Builder.
		Select(tableColumNotes()).
		From(tableNameDoctorNotes)

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
		var note doctor_notes.DoctorNote
		if err = rows.Scan(
			&note.Id,
			&note.AppointmentId,
			&note.DoctorId,
			&note.PatientId,
			&note.Prescription,
			&note.CreatedAt,
			&upTime,
			&delTime,
		); err != nil {
			return nil, err
		}

		if upTime.Valid {
			note.UpdatedAt = upTime.Time
		}

		if delTime.Valid {
			note.DeletedAt = delTime.Time
		}

		notes.DoctorNotes = append(notes.DoctorNotes, &note)

		notes.Count += 1
	}
	return &notes, nil
}

func (r *DoctorNotes) UpdateDoctorNotes(ctx context.Context, req *doctor_notes.UpdateDoctorNoteReq) (*doctor_notes.DoctorNote, error) {
	var (
		note    doctor_notes.DoctorNote
		upTime  sql.NullTime
		delTime sql.NullTime
	)
	toSql, args, err := r.db.Sq.Builder.
		Update(tableNameDoctorNotes).
		SetMap(map[string]interface{}{
			"appointment_id": req.AppointmentId,
			"doctor_id":      req.DoctorId,
			"patient_id":     req.PatientId,
			"prescription":   req.Prescription,
			"updated_at":     time.Now(),
		}).
		Where(r.db.Sq.Equal(req.Field, req.Value)).
		Suffix(fmt.Sprintf("RETURNING %s", tableColumNotes())).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err = r.db.QueryRow(ctx, toSql, args...).Scan(
		&note.Id,
		&note.AppointmentId,
		&note.DoctorId,
		&note.PatientId,
		&note.Prescription,
		&note.CreatedAt,
		&upTime,
		&delTime,
	); err != nil {
		return nil, err
	}

	if upTime.Valid {
		note.UpdatedAt = upTime.Time
	}

	if delTime.Valid {
		note.DeletedAt = delTime.Time
	}

	return &note, nil

}

func (r *DoctorNotes) DeleteDoctorNotes(ctx context.Context, req *doctor_notes.FieldValueReq) (*doctor_notes.StatusRes, error) {
	if !req.DeleteStatus {
		toSql, args, err := r.db.Sq.Builder.
			Update(tableNameDoctorNotes).
			Set("deleted_at", time.Now()).
			Where(r.db.Sq.EqualMany(map[string]interface{}{
				"deleted_at": nil,
				req.Field:    req.Value,
			})).
			ToSql()
		if err != nil {
			return &doctor_notes.StatusRes{Status: false}, err
		}

		_, err = r.db.Exec(ctx, toSql, args...)

		if err != nil {
			return &doctor_notes.StatusRes{Status: false}, err
		}
		return &doctor_notes.StatusRes{Status: true}, nil

	} else {
		toSql, args, err := r.db.Sq.Builder.
			Delete(tableNameDoctorNotes).
			Where(r.db.Sq.Equal(req.Field, req.Value)).
			ToSql()

		if err != nil {
			return &doctor_notes.StatusRes{Status: false}, err
		}

		_, err = r.db.Exec(ctx, toSql, args...)

		if err != nil {
			return &doctor_notes.StatusRes{Status: false}, err
		}
		return &doctor_notes.StatusRes{Status: true}, nil
	}
}
