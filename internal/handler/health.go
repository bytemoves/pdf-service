package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)


type HealthHandler interface {
	Liveness (c *gin.Context)
	Readiness ( c*gin.Context)
}


type healthHandler struct{
	log *slog.Logger
} 


func NewHealth (log *slog.Logger) healthHandler {
	return healthHandler{
		log: log,
	}
}

type HealthStatus string

const OKStatus HealthStatus = "ok"

type HealthResponse struct {
	Status HealthStatus `json:"status"`
}


var okHealthResponse = HealthResponse{Status: OKStatus}


func (h healthHandler) Liveness(c *gin.Context) {
	h.log.Info("Handled liveness check request")
	c.JSON(http.StatusOK, okHealthResponse)
}

// Readiness
//
//	@Tags			health
//	@Produce		json
//
//	@Success		200 {object} handler.HealthResponse
//	@Router			/health/readiness [get]
func (h healthHandler) Readiness(c *gin.Context) {
	h.log.Info("Handled readiness check request")
	c.JSON(http.StatusOK, okHealthResponse)
}