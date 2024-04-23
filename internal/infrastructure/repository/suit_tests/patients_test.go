package suit_tests

import (
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	db "booking_service/internal/pkg/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookingPatientsTestSite struct {
	suite.Suite
	Repository  *repo.BookingPatients
	CleanUpFunc func()
}

func (s *BookingPatientsTestSite) SetupSuite() {
	pgPool, _ := db.New(config.New())
	s.Repository = repo.NewBookingPatients(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *BookingPatientsTestSite) TestUserCRUD() {

}

func (s *BookingPatientsTestSite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestBookingPatientsTestSuite(t *testing.T) {
	suite.Run(t, new(BookingPatientsTestSite))
}
