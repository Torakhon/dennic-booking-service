package suit_tests

import (
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	db "booking_service/internal/pkg/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookingArchiveTestSite struct {
	suite.Suite
	Repository  *repo.BookingArchive
	CleanUpFunc func()
}

func (s *BookingArchiveTestSite) SetupSuite() {
	pgPool, _ := db.New(config.New())
	s.Repository = repo.NewBookingArchive(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *BookingArchiveTestSite) TestUserCRUD() {

}

func (s *BookingArchiveTestSite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestUBookingArchiveTestSuite(t *testing.T) {
	suite.Run(t, new(BookingArchiveTestSite))
}
