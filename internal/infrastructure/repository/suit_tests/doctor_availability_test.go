package suit_tests

import (
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	db "booking_service/internal/pkg/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
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

}

func (s *DoctorAvailabilityTestSite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestDoctorAvailabilityTestSuite(t *testing.T) {
	suite.Run(t, new(DoctorAvailabilityTestSite))
}
