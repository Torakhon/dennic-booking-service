package suit_tests

import (
	"booking_service/internal/entity/doctor_availability"
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	db "booking_service/internal/pkg/postgres"
	"context"
	"github.com/rickb777/date"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type DoctorAvailabilityTestSite struct {
	suite.Suite
	Repository  *repo.DoctorAvailability
	CleanUpFunc func()
}

func (s *DoctorAvailabilityTestSite) SetupSuite() {
	pgPool, _ := db.New(config.New())
	s.Repository = repo.NewDoctorAvailability(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *DoctorAvailabilityTestSite) TestUserCRUD() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(2))
	defer cancel()

	doctorDate, _ := date.AutoParse("1231-02-02")
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2000-01-01 14:14:14")
	endTime, _ := time.Parse("2006-01-02 15:04:05", "2000-01-01 11:11:11")

	createReq := &doctor_availability.CreateDoctorAvailability{
		DepartmentId: "8bde0f01-33d9-4d29-9e0f-0133d9cd29d8",
		DoctorId:     "b78c39b4-c038-4e6b-8c39-b4c0384e6bf9",
		DoctorDate:   doctorDate,
		StartTime:    startTime,
		EndTime:      endTime,
		Status:       "available",
	}

	createRes, err := s.Repository.CreateDoctorAvailability(ctx, createReq)
	s.Suite.NoError(err)
	s.Suite.NotNil(createRes)
	s.Suite.Equal(createRes.DepartmentId, createReq.DepartmentId)
	s.Suite.Equal(createRes.DoctorId, createReq.DoctorId)
	s.Suite.Equal(createRes.DoctorDate, createReq.DoctorDate)
	s.Suite.Equal(createRes.StartTime, createReq.StartTime)
	s.Suite.Equal(createRes.EndTime, createReq.EndTime)
	s.Suite.Equal(createRes.Status, createReq.Status)

	hardDelRes, err := s.Repository.DeleteDoctorAvailability(ctx, &doctor_availability.FieldValueReq{
		Field:        "doctor_id",
		Value:        createReq.DoctorId,
		DeleteStatus: true,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(hardDelRes)
	s.Suite.Equal(hardDelRes.Status, true)

}

func (s *DoctorAvailabilityTestSite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestDoctorAvailabilityTestSuite(t *testing.T) {
	suite.Run(t, new(DoctorAvailabilityTestSite))
}
