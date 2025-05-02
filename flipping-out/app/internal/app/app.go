package app

import (
	"github.com/joshbrgs/flipping-out/internal/repositories"
	"github.com/joshbrgs/flipping-out/internal/services"
	of "github.com/open-feature/go-sdk/openfeature"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	MongoClient   *mongo.Client
	FeatureClient *of.Client

	FeatureRepo    repositories.FeatureFlagRepository
	FeatureService services.FeatureService
	WelcomeService services.WelcomeService
}

func NewContainer(mongoClient *mongo.Client, featureClient *of.Client) *Container {
	repo := repositories.NewFeatureFlagRepository(mongoClient, "appConfig", "featureFlags")
	service := services.NewFeatureService(repo)
	welcomeService := services.NewWelcomeService()

	return &Container{
		MongoClient:    mongoClient,
		FeatureClient:  featureClient,
		FeatureRepo:    repo,
		FeatureService: service,
		WelcomeService: welcomeService,
	}
}
