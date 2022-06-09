package repositories

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/zawachte/bimah/internal/models"
)

// FleetRepository
type FleetRepository interface {
	GetFleetById(context.Context, int) (models.Fleet, error)
	GetFleets(context.Context) ([]models.Fleet, error)
	CreateFleet(context.Context, models.Fleet) error
	DeleteFleetById(context.Context, int) error
}

type FleetRepositoryParams struct {
	DatabaseUrl string
}

// NewFleetRepository creates a fleet repository.
func NewFleetRepository(ctx context.Context, params FleetRepositoryParams) (FleetRepository, error) {
	if params.DatabaseUrl == "" {
		return &fleetRepositoryMemory{}, nil
	}

	dbPool, err := pgxpool.Connect(ctx, params.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return &fleetRepositoryPostGres{dbPool}, nil
}

type fleetSql struct {
	CreationDate time.Time        `json:"creation_on,omitempty"`
	Endpoints    pgtype.TextArray `json:"endpoints,omitempty"`
	Id           int              `json:"fleet_id,omitempty"`
	Tags         pgtype.TextArray `json:"tags,omitempty"`
}

func fleetSqlToFleet(input fleetSql) models.Fleet {

	tags := []string{}
	for _, ele := range input.Tags.Elements {
		tags = append(tags, ele.String)
	}

	endpoints := []string{}
	for _, ele := range input.Endpoints.Elements {
		endpoints = append(endpoints, ele.String)
	}

	return models.Fleet{
		CreationDate: &input.CreationDate,
		Endpoints:    &endpoints,
		Tags:         &tags,
		Id:           &input.Id,
	}
}

func fleetToFleetSql(input models.Fleet) (fleetSql, error) {

	var tagsPgArray pgtype.TextArray
	err := tagsPgArray.Set(input.Tags)
	if err != nil {
		return fleetSql{}, nil
	}

	var endpointsPgArray pgtype.TextArray
	err = endpointsPgArray.Set(input.Endpoints)
	if err != nil {
		return fleetSql{}, nil
	}

	return fleetSql{
		CreationDate: *input.CreationDate,
		Endpoints:    endpointsPgArray,
		Tags:         tagsPgArray,
		Id:           *input.Id,
	}, nil
}

type fleetRepositoryPostGres struct {
	dbPool *pgxpool.Pool
}

func (fr *fleetRepositoryPostGres) GetFleets(ctx context.Context) ([]models.Fleet, error) {
	rows, err := fr.dbPool.Query(ctx, "select * from fleets")
	if err != nil {
		return nil, err
	}

	fleets := []models.Fleet{}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		fleetInternal := fleetSql{
			CreationDate: values[0].(time.Time),
			Endpoints:    values[1].(pgtype.TextArray),
			Tags:         values[2].(pgtype.TextArray),
			Id:           values[3].(int),
		}

		fleets = append(fleets, fleetSqlToFleet(fleetInternal))
	}

	return fleets, nil
}

func (fr *fleetRepositoryPostGres) GetFleetById(ctx context.Context, id int) (models.Fleet, error) {
	fleetInternal := fleetSql{}
	row := fr.dbPool.QueryRow(ctx, "select * from fleets where fleet_id=$1", id)
	err := row.Scan(
		&fleetInternal.CreationDate,
		&fleetInternal.Endpoints,
		&fleetInternal.Tags,
		&fleetInternal.Id)
	if err != nil {
		return models.Fleet{}, nil
	}

	return fleetSqlToFleet(fleetInternal), nil
}

func (fr *fleetRepositoryPostGres) CreateFleet(ctx context.Context, fleet models.Fleet) error {
	var lastInsertID int

	fleetInternal, err := fleetToFleetSql(fleet)
	if err != nil {
		return err
	}

	err = fr.dbPool.QueryRow(ctx, "INSERT INTO fleets(created_on, endpoints, tags) VALUES($1, $2, $3) returning fleet_id;",
		time.Now(),
		fleetInternal.Endpoints,
		fleetInternal.Tags).Scan(&lastInsertID)
	if err != nil {
		return err
	}

	return nil
}

func (fr *fleetRepositoryPostGres) DeleteFleetById(ctx context.Context, id int) error {
	var deletedID int

	err := fr.dbPool.QueryRow(ctx, "delete from fleets where fleet_id=$1 returning fleet_id;", id).Scan(&deletedID)
	if err != nil {
		return err
	}

	return nil
}

func (fr *fleetRepositoryPostGres) Close() {
	fr.dbPool.Close()
}

type fleetRepositoryMemory struct {
	mu        sync.Mutex
	simpleMap map[int]models.Fleet
}

func (fr *fleetRepositoryMemory) GetFleets(ctx context.Context) ([]models.Fleet, error) {
	fleets := []models.Fleet{}
	for _, v := range fr.simpleMap {
		fleets = append(fleets, v)
	}

	return fleets, nil
}

func (fr *fleetRepositoryMemory) GetFleetById(ctx context.Context, id int) (models.Fleet, error) {
	fleet, ok := fr.simpleMap[id]
	if !ok {
		return models.Fleet{}, errors.New("no fleet found")
	}

	return fleet, nil
}

func (fr *fleetRepositoryMemory) CreateFleet(ctx context.Context, fleet models.Fleet) error {

	fr.mu.Lock()
	defer fr.mu.Unlock()
	currentLength := len(fr.simpleMap)
	fr.simpleMap[currentLength+1] = fleet
	return nil
}

func (fr *fleetRepositoryMemory) DeleteFleetById(ctx context.Context, id int) error {
	fr.mu.Lock()
	defer fr.mu.Unlock()
	delete(fr.simpleMap, id)
	return nil
}
