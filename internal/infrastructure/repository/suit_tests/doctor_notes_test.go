package suit_tests

import (
	repo "booking_service/internal/infrastructure/repository/postgresql"
	"booking_service/internal/pkg/config"
	db "booking_service/internal/pkg/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DoctorNotesTestSite struct {
	suite.Suite
	Repository  *repo.DoctorNotes
	CleanUpFunc func()
}

func (s *DoctorNotesTestSite) SetupSuite() {
	pgPool, _ := db.New(config.New())
	s.Repository = repo.NewDoctorNotes(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *DoctorNotesTestSite) TestUserCRUD() {

}

func (s *DoctorNotesTestSite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestDoctorNotesTestSuite(t *testing.T) {
	suite.Run(t, new(DoctorNotesTestSite))
}
