package services

import (
	"context"

	"github.com/zawachte/bimah/internal/models"
	"github.com/zawachte/bimah/internal/repositories"
)

// FleetService
type FleetService interface {
	GetFleetById(context.Context, int) (models.Fleet, error)
	GetFleets(context.Context) ([]models.Fleet, error)
	CreateFleet(context.Context, models.Fleet) error
	DeleteFleetById(context.Context, int) error
}

type FleetServiceParams struct {
	DatabaseUrl string
}

// NewFleetService creates an order service.
func NewFleetService(ctx context.Context, params FleetServiceParams) (FleetService, error) {
	repo, err := repositories.NewFleetRepository(ctx, repositories.FleetRepositoryParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	return &fleetService{repo}, nil
}

type fleetService struct {
	fleetRepository repositories.FleetRepository
}

func (fs *fleetService) GetFleets(ctx context.Context) ([]models.Fleet, error) {
	return nil, nil
}

func (fs *fleetService) GetFleetById(ctx context.Context, id int) (models.Fleet, error) {
	return models.Fleet{}, nil
}

func (fs *fleetService) CreateFleet(ctx context.Context, fleet models.Fleet) error {
	return nil
}

func (fs *fleetService) DeleteFleetById(ctx context.Context, id int) error {
	return nil
}
