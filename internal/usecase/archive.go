package usecase

import (
	"booking_service/internal/entity/archive"
	"context"
)

// BookedArchiveUseCase -.
type BookedArchiveUseCase struct {
	Repo Archive
}

// NewBookedArchive -.
func NewBookedArchive(r Archive) *BookedArchiveUseCase {
	return &BookedArchiveUseCase{
		Repo: r,
	}
}

func (r *BookedArchiveUseCase) CreateArchive(ctx context.Context, req *archive.CreatedArchive) (*archive.Archive, error) {
	return r.Repo.CreateArchive(ctx, req)
}

func (r *BookedArchiveUseCase) GetArchive(ctx context.Context, req *archive.FieldValueReq) (*archive.Archive, error) {
	return r.Repo.GetArchive(ctx, req)
}

func (r *BookedArchiveUseCase) GetAllArchive(ctx context.Context, req *archive.GetAllArchives) (*archive.ArchivesType, error) {
	return r.Repo.GetAllArchive(ctx, req)
}

func (r *BookedArchiveUseCase) UpdateArchive(ctx context.Context, req *archive.UpdateArchive) (*archive.Archive, error) {
	return r.Repo.UpdateArchive(ctx, req)
}

func (r *BookedArchiveUseCase) DeleteArchive(ctx context.Context, req *archive.FieldValueReq) (*archive.StatusRes, error) {
	return r.Repo.DeleteArchive(ctx, req)
}
