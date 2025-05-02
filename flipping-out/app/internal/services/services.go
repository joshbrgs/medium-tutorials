package services

import (
	"context"

	"github.com/joshbrgs/flipping-out/internal/models"
	"github.com/joshbrgs/flipping-out/internal/repositories"
)

type FeatureService interface {
	IsFeatureEnabled(ctx context.Context, name string) (bool, error)
	GetAllFlags(ctx context.Context) ([]models.FeatureFlag, error)
	GetFlagByID(ctx context.Context, id string) (*models.FeatureFlag, error)
	CreateFlag(ctx context.Context, flag models.FeatureFlag) error
	UpdateFlag(ctx context.Context, id string, update map[string]interface{}) error
	DeleteFlag(ctx context.Context, id string) error
}

type featureService struct {
	repo repositories.FeatureFlagRepository
}

func NewFeatureService(repo repositories.FeatureFlagRepository) FeatureService {
	return &featureService{repo: repo}
}

func (s *featureService) IsFeatureEnabled(ctx context.Context, name string) (bool, error) {
	flags, err := s.repo.GetAll(ctx)
	if err != nil {
		return false, err
	}

	for _, flag := range flags {
		if flag.Flag == name {
			return flag.Variations.DefaultVar, nil
		}
	}

	return false, nil
}

func (s *featureService) GetAllFlags(ctx context.Context) ([]models.FeatureFlag, error) {
	return s.repo.GetAll(ctx)
}

func (s *featureService) GetFlagByID(ctx context.Context, id string) (*models.FeatureFlag, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *featureService) CreateFlag(ctx context.Context, flag models.FeatureFlag) error {
	return s.repo.Create(ctx, flag)
}

func (s *featureService) UpdateFlag(ctx context.Context, id string, update map[string]interface{}) error {
	return s.repo.Update(ctx, id, update)
}

func (s *featureService) DeleteFlag(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
