package services

import (
	"context"

	"github.com/zawachte/bimah/internal/models"
	"github.com/zawachte/bimah/internal/repositories"
)

// TlsconfigService
type TlsconfigService interface {
	GetTlsconfigById(context.Context, int) (models.Tlsconfig, error)
	GetTlsconfigs(context.Context) ([]models.Tlsconfig, error)
	CreateTlsconfig(context.Context, models.Tlsconfig) error
	DeleteTlsconfigById(context.Context, int) error
}

type TlsconfigServiceParams struct {
	DatabaseUrl string
}

// NewTlsconfigService creates an tlsconfig service.
func NewTlsconfigService(ctx context.Context, params TlsconfigServiceParams) (TlsconfigService, error) {
	repo, err := repositories.NewTlsconfigRepository(ctx, repositories.TlsconfigRepositoryParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	return &tlsconfigService{repo}, nil
}

type tlsconfigService struct {
	tlsconfigRepository repositories.TlsconfigRepository
}

func (fs *tlsconfigService) GetTlsconfigs(ctx context.Context) ([]models.Tlsconfig, error) {
	return fs.tlsconfigRepository.GetTlsconfigs(ctx)
}

func (fs *tlsconfigService) GetTlsconfigById(ctx context.Context, id int) (models.Tlsconfig, error) {
	return fs.tlsconfigRepository.GetTlsconfigById(ctx, id)
}

func (fs *tlsconfigService) CreateTlsconfig(ctx context.Context, tlsconfig models.Tlsconfig) error {
	return fs.tlsconfigRepository.CreateTlsconfig(ctx, tlsconfig)
}

func (fs *tlsconfigService) DeleteTlsconfigById(ctx context.Context, id int) error {
	return fs.tlsconfigRepository.DeleteTlsconfigById(ctx, id)
}
