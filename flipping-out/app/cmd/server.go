package cmd

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/joshbrgs/flipping-out/internal/api/v1"
	"github.com/joshbrgs/flipping-out/internal/app"
	gofeatureflag "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag/pkg"
	of "github.com/open-feature/go-sdk/openfeature"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func StartServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Creation of the GO Feature Flag provider to our relay proxy in kind
	provider, err := gofeatureflag.NewProvider(
		gofeatureflag.ProviderOptions{
			Endpoint: "http://go-feature-flag-relay-proxy:1031",
		})
	if err != nil {
		log.Fatalf("Failed to connect to Relay: %v", err)
	}

	// Setting the provider to the OpenFeature SDK
	err = of.SetProviderAndWait(provider)
	if err != nil {
		log.Fatalf("Failed set provider: %v", err)
	}
	ofClt := of.NewClient("my-openfeature-gin-client")

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:password@mongo-mongodb:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// bootstrap services
	container := app.NewContainer(mongoClient, ofClt)

	// start application
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1.RegisterRoutes(r, container)
	log.Println("Server started at :3001")

	r.Run(":3001")
}
