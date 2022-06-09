package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zawachte/bimah/internal/models"
)

// TlsconfigRepository
type TlsconfigRepository interface {
	GetTlsconfigById(context.Context, int) (models.Tlsconfig, error)
	GetTlsconfigs(context.Context) ([]models.Tlsconfig, error)
	CreateTlsconfig(context.Context, models.Tlsconfig) error
	DeleteTlsconfigById(context.Context, int) error
}

type TlsconfigRepositoryParams struct {
	DatabaseUrl string
}

// NewTlsconfigRepository creates a tlsconfig repository.
func NewTlsconfigRepository(ctx context.Context, params TlsconfigRepositoryParams) (TlsconfigRepository, error) {
	if params.DatabaseUrl == "" {
		return &tlsconfigRepositoryMemory{}, nil
	}

	dbPool, err := pgxpool.Connect(ctx, params.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return &tlsconfigRepositoryPostGres{dbPool}, nil
}

type tlsconfigRepositoryPostGres struct {
	dbPool *pgxpool.Pool
}

func (fr *tlsconfigRepositoryPostGres) GetTlsconfigs(ctx context.Context) ([]models.Tlsconfig, error) {
	return nil, nil
}

func (fr *tlsconfigRepositoryPostGres) GetTlsconfigById(ctx context.Context, id int) (models.Tlsconfig, error) {
	return models.Tlsconfig{}, nil
}

func (fr *tlsconfigRepositoryPostGres) CreateTlsconfig(ctx context.Context, tlsconfig models.Tlsconfig) error {
	return nil
}

func (fr *tlsconfigRepositoryPostGres) DeleteTlsconfigById(ctx context.Context, id int) error {
	return nil
}

type tlsconfigRepositoryMemory struct {
	simpleMap map[int]string
}

func (fr *tlsconfigRepositoryMemory) GetTlsconfigs(ctx context.Context) ([]models.Tlsconfig, error) {
	return nil, nil
}

func (fr *tlsconfigRepositoryMemory) GetTlsconfigById(ctx context.Context, id int) (models.Tlsconfig, error) {
	return models.Tlsconfig{}, nil
}

func (fr *tlsconfigRepositoryMemory) CreateTlsconfig(ctx context.Context, tlsconfg models.Tlsconfig) error {
	return nil
}

func (fr *tlsconfigRepositoryMemory) DeleteTlsconfigById(ctx context.Context, id int) error {
	return nil
}
