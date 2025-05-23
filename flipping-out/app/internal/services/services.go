package services

import (
	"context"
	"fmt"
	"log"

	"github.com/joshbrgs/flipping-out/internal/models"
	"github.com/joshbrgs/flipping-out/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type FeatureService interface {
	IsFeatureEnabled(ctx context.Context, name string) (bool, error)
	GetAllFlags(ctx context.Context) ([]models.FeatureFlag, error)
	GetFlagByID(ctx context.Context, id string) (*models.FeatureFlag, error)
	CreateFlag(ctx context.Context, flag models.FeatureFlag) error
	UpdateFlag(ctx context.Context, id string, update models.FeatureFlag) error
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
	idd, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert ID. Err: %w", err)
	}
	return s.repo.GetByID(ctx, idd)
}

func (s *featureService) CreateFlag(ctx context.Context, flag models.FeatureFlag) error {
	return s.repo.Create(ctx, flag)
}

func (s *featureService) UpdateFlag(ctx context.Context, id string, update models.FeatureFlag) error {
	idd, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("Failed to convert ID. Err: %w", err)
	}
	model, err := s.GetFlagByID(ctx, id)
	if err != nil {
		return fmt.Errorf("Failed to get flag by ID. Err: %w", err)
	}

	log.Println(update)

	updateData := models.FeatureFlag{
		Flag:        model.Flag,
		Variations:  model.Variations,
		DefaultRule: update.DefaultRule,
	}

	log.Println(updateData)

	return s.repo.Update(ctx, idd, updateData)
}

func (s *featureService) DeleteFlag(ctx context.Context, id string) error {
	idd, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("Failed to convert ID. Err: %w", err)
	}
	return s.repo.Delete(ctx, idd)
}
