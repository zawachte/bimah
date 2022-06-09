package services

import (
	"context"

	"github.com/zawachte/bimah/internal/models"
	"github.com/zawachte/bimah/internal/repositories"
)

// RuleService
type RuleService interface {
	GetRuleById(context.Context, int) (models.Rule, error)
	GetRules(context.Context) ([]models.Rule, error)
	CreateRule(context.Context, models.Rule) error
	DeleteRuleById(context.Context, int) error
}

type RuleServiceParams struct {
	DatabaseUrl string
}

// NewRuleService creates an order service.
func NewRuleService(ctx context.Context, params RuleServiceParams) (RuleService, error) {
	repo, err := repositories.NewRuleRepository(ctx, repositories.RuleRepositoryParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	return &ruleService{repo}, nil
}

type ruleService struct {
	ruleRepository repositories.RuleRepository
}

func (fs *ruleService) GetRules(ctx context.Context) ([]models.Rule, error) {
	return fs.ruleRepository.GetRules(ctx)
}

func (fs *ruleService) GetRuleById(ctx context.Context, id int) (models.Rule, error) {
	return fs.ruleRepository.GetRuleById(ctx, id)
}

func (fs *ruleService) CreateRule(ctx context.Context, rule models.Rule) error {
	return fs.ruleRepository.CreateRule(ctx, rule)
}

func (fs *ruleService) DeleteRuleById(ctx context.Context, id int) error {
	return fs.ruleRepository.DeleteRuleById(ctx, id)
}
