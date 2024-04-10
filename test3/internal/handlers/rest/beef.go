package rest

import (
	"7solutionstest3/internal/models"
	"7solutionstest3/internal/utils/httputil"
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BeefService interface {
	BeefSummary(ctx context.Context) (*models.BeefResponse, error)
}

type BeefSummaryHandler struct {
	beefService BeefService
}

func NewBeefSummaryHandler(beefService BeefService) *BeefSummaryHandler {
	return &BeefSummaryHandler{beefService: beefService}
}

func (h *BeefSummaryHandler) BeefSummary(c *gin.Context) {
	request := c.Request
	ctx := request.Context()

	data, err := h.beefService.BeefSummary(ctx)
	if err != nil {
		slog.Error(err.Error(), err)
		err := httputil.NewHttpError(err.Error(), "", http.StatusInternalServerError)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}
