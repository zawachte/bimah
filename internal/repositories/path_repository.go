package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zawachte/bimah/internal/models"
)

// PathRepository
type PathRepository interface {
	GetPathById(context.Context, int) (models.Path, error)
	GetPaths(context.Context) ([]models.Path, error)
	CreatePath(context.Context, models.Path) error
	DeletePathById(context.Context, int) error
}

type PathRepositoryParams struct {
	DatabaseUrl string
}

// NewPathRepository creates a path repository.
func NewPathRepository(ctx context.Context, params PathRepositoryParams) (PathRepository, error) {
	if params.DatabaseUrl == "" {
		return &pathRepositoryMemory{}, nil
	}

	dbPool, err := pgxpool.Connect(ctx, params.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return &pathRepositoryPostGres{dbPool}, nil
}

type pathRepositoryPostGres struct {
	dbPool *pgxpool.Pool
}

func (fr *pathRepositoryPostGres) GetPaths(ctx context.Context) ([]models.Path, error) {
	return nil, nil
}

func (fr *pathRepositoryPostGres) GetPathById(ctx context.Context, id int) (models.Path, error) {
	return models.Path{}, nil
}

func (fr *pathRepositoryPostGres) CreatePath(ctx context.Context, path models.Path) error {
	return nil
}

func (fr *pathRepositoryPostGres) DeletePathById(ctx context.Context, id int) error {
	return nil
}

type pathRepositoryMemory struct {
	simpleMap map[int]string
}

func (fr *pathRepositoryMemory) GetPaths(ctx context.Context) ([]models.Path, error) {
	return nil, nil
}

func (fr *pathRepositoryMemory) GetPathById(ctx context.Context, id int) (models.Path, error) {
	return models.Path{}, nil
}

func (fr *pathRepositoryMemory) CreatePath(ctx context.Context, path models.Path) error {
	return nil
}

func (fr *pathRepositoryMemory) DeletePathById(ctx context.Context, id int) error {
	return nil
}
