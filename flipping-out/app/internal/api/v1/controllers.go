package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshbrgs/flipping-out/internal/models"
	"github.com/joshbrgs/flipping-out/internal/services"
	of "github.com/open-feature/go-sdk/openfeature"
)

type FlagController struct {
	service    services.FeatureService
	flagClient *of.Client
}

func NewFlagController(service services.FeatureService, flagClient *of.Client) *FlagController {
	return &FlagController{service: service, flagClient: flagClient}
}

func (fc *FlagController) getFlagsHandler(c *gin.Context) {
	flags, err := fc.service.GetAllFlags(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch flags", "err_detailed": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flags)
}

func (fc *FlagController) getFlagHandler(c *gin.Context) {
	id := c.Param("id")
	flag, err := fc.service.GetFlagByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "flag not found", "err_detailed": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flag)
}

func (fc *FlagController) createFlagHandler(c *gin.Context) {
	var flag models.FeatureFlag
	if err := c.ShouldBindJSON(&flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if err := fc.service.CreateFlag(c.Request.Context(), flag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create flag", "err_detailed": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, flag)
}

func (fc *FlagController) updateFlagHandler(c *gin.Context) {
	id := c.Param("id")
	var update models.FeatureFlag
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if err := fc.service.UpdateFlag(c.Request.Context(), id, update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update flag", "err_detailed": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "flag updated"})
}

func (fc *FlagController) deleteFlagHandler(c *gin.Context) {
	id := c.Param("id")
	if err := fc.service.DeleteFlag(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete flag", "err_detailed": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "flag deleted"})
}
