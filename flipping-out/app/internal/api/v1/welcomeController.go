package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshbrgs/flipping-out/internal/services"
	of "github.com/open-feature/go-sdk/openfeature"
)

type WelcomeController struct {
	service    services.WelcomeService
	flagClient *of.Client
}

func NewWelcomeController(service services.WelcomeService, flagClient *of.Client) *WelcomeController {
	return &WelcomeController{service: service, flagClient: flagClient}
}

// Example of using a featureflag decoupled from the buisness logic
func (fc *WelcomeController) getWelcomeHandler(c *gin.Context) {
	welcomeMessage, _ := fc.flagClient.BooleanValue(c, "welcome-message", false, of.EvaluationContext{})
	user := c.GetHeader("x-example-header")

	if welcomeMessage {
		msg := fc.service.HelloWorld()
		c.JSON(http.StatusOK, msg)
	} else {
		msg := fc.service.HelloWorldAgain(user)
		c.JSON(http.StatusOK, msg)
	}
}
