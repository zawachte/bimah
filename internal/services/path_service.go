package services

import (
	"context"

	"github.com/zawachte/bimah/internal/models"
	"github.com/zawachte/bimah/internal/repositories"
)

// PathService
type PathService interface {
	GetPathById(context.Context, int) (models.Path, error)
	GetPaths(context.Context) ([]models.Path, error)
	CreatePath(context.Context, models.Path) error
	DeletePathById(context.Context, int) error
}

type PathServiceParams struct {
	DatabaseUrl string
}

// NewPathService creates an path service.
func NewPathService(ctx context.Context, params PathServiceParams) (PathService, error) {
	repo, err := repositories.NewPathRepository(ctx, repositories.PathRepositoryParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	return &pathService{repo}, nil
}

type pathService struct {
	pathRepository repositories.PathRepository
}

func (fs *pathService) GetPaths(ctx context.Context) ([]models.Path, error) {
	return fs.pathRepository.GetPaths(ctx)
}

func (fs *pathService) GetPathById(ctx context.Context, id int) (models.Path, error) {
	return fs.pathRepository.GetPathById(ctx, id)
}

func (fs *pathService) CreatePath(ctx context.Context, path models.Path) error {
	return fs.pathRepository.CreatePath(ctx, path)
}

func (fs *pathService) DeletePathById(ctx context.Context, id int) error {
	return fs.pathRepository.DeletePathById(ctx, id)
}
