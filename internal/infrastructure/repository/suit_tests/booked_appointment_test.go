package suit_tests

import (
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	db "booking_service/internal/pkg/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookingAppointmentTestSite struct {
	suite.Suite
	Repository  *repo.BookingAppointment
	CleanUpFunc func()
}

func (s *BookingAppointmentTestSite) SetupSuite() {
	pgPool, _ := db.New(config.New())
	s.Repository = repo.NewBookingAppointment(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *BookingAppointmentTestSite) TestUserCRUD() {

}

func (s *BookingAppointmentTestSite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestBookingAppointmentTestSuite(t *testing.T) {
	suite.Run(t, new(BookingAppointmentTestSite))
}
