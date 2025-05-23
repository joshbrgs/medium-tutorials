package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/joshbrgs/flipping-out/internal/app"
)

func RegisterRoutes(r *gin.Engine, c *app.Container) {
	api := r.Group("/v1")

	registerFlagRoutes(api, c)
	registerUserRoutes(api, c)
	registerWebsockets(api, c)
}

func registerFlagRoutes(r *gin.RouterGroup, c *app.Container) {
	api := r.Group("/flags")

	flagController := NewFlagController(c.FeatureService, c.FeatureClient)

	api.GET("", flagController.getFlagsHandler)
	api.GET("/:id", flagController.getFlagHandler)
	api.POST("", flagController.createFlagHandler)
	api.PATCH("/:id", flagController.updateFlagHandler)
	api.DELETE("/:id", flagController.deleteFlagHandler)
}

func registerUserRoutes(r *gin.RouterGroup, c *app.Container) {
	api := r.Group("/welcome")

	exampleController := NewWelcomeController(c.WelcomeService, c.FeatureClient)

	api.GET("", exampleController.getWelcomeHandler)
}

func registerWebsockets(r *gin.RouterGroup, c *app.Container) {
	api := r.Group("/ws")

	websocketController := NewWebsocketController(c.WebsocketHub)

	api.GET("", websocketController.HandleWebsocket)
}
