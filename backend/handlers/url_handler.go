package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	urlService services.URLService
}

func NewURLHandler(urlService services.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

// ProcessURL godoc
// @Summary Process URL cleanup or redirection
// @Description Refine, canonicalize, or redirect a URL based on the designated operation type
// @Tags url-processor
// @Accept json
// @Produce json
// @Param request body models.URLRequest true "URL processing request"
// @Success 200 {object} models.URLResponse
// @Failure 400 {object} map[string]string
// @Router /url-process [post]
func (h *URLHandler) ProcessURL(c *gin.Context) {
	var req models.URLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	processedURL, err := h.urlService.ProcessURL(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.URLResponse{
		ProcessedURL: processedURL,
	})
}
