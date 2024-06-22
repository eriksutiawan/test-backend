package delivery

import (
	"net/http"
	"test-backend/internal/app/prediction/usecase"

	"github.com/gin-gonic/gin"
)

// Predict godoc
// @Summary Predict label from an image
// @Description Predicts label from uploaded image
// @Tags prediction
// @Accept multipart/form-data
// @Produce json
// @Param userID header string true "User ID"
// @Param image formData file true "Image file to predict"
// @Success 200 {object} usecase.PredictResponse
// @Failure 400 {object} map[string]string "error": "Invalid image"
// @Failure 500 {object} map[string]string "error": "Internal Server Error"
// @Router /predict [post]
func Predict(c *gin.Context, prediction usecase.Predictor) {
	userID, _ := c.Get("userID")
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID is not a string"})
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
		return
	}

	resp, err := prediction.Predict(c, usecase.PredictDto{
		File:   file,
		UserId: userIDStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prediction": resp})
}
