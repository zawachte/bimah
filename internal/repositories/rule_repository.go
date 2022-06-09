package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zawachte/bimah/internal/models"
)

// RuleRepository
type RuleRepository interface {
	GetRuleById(context.Context, int) (models.Rule, error)
	GetRules(context.Context) ([]models.Rule, error)
	CreateRule(context.Context, models.Rule) error
	DeleteRuleById(context.Context, int) error
}

type RuleRepositoryParams struct {
	DatabaseUrl string
}

// NewRuleRepository creates a rule repository.
func NewRuleRepository(ctx context.Context, params RuleRepositoryParams) (RuleRepository, error) {
	if params.DatabaseUrl == "" {
		return &ruleRepositoryMemory{}, nil
	}

	dbPool, err := pgxpool.Connect(ctx, params.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return &ruleRepositoryPostGres{dbPool}, nil
}

type ruleRepositoryPostGres struct {
	dbPool *pgxpool.Pool
}

func (fr *ruleRepositoryPostGres) GetRules(ctx context.Context) ([]models.Rule, error) {
	return nil, nil
}

func (fr *ruleRepositoryPostGres) GetRuleById(ctx context.Context, id int) (models.Rule, error) {
	return models.Rule{}, nil
}

func (fr *ruleRepositoryPostGres) CreateRule(ctx context.Context, rule models.Rule) error {
	return nil
}

func (fr *ruleRepositoryPostGres) DeleteRuleById(ctx context.Context, id int) error {
	return nil
}

type ruleRepositoryMemory struct {
	simpleMap map[int]string
}

func (fr *ruleRepositoryMemory) GetRules(ctx context.Context) ([]models.Rule, error) {
	return nil, nil
}

func (fr *ruleRepositoryMemory) GetRuleById(ctx context.Context, id int) (models.Rule, error) {
	return models.Rule{}, nil
}

func (fr *ruleRepositoryMemory) CreateRule(ctx context.Context, rule models.Rule) error {
	return nil
}

func (fr *ruleRepositoryMemory) DeleteRuleById(ctx context.Context, id int) error {
	return nil
}
