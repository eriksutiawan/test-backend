package delivery

import (
	"net/http"
	"test-backend/internal/app/user/usecase"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get user details by ID
// @Description Retrieves user details based on user ID
// @Tags user
// @Accept json
// @Produce json
// @Param userID header string true "User ID"
// @Success 200 {object} usecase.UserResponse
// @Failure 500 {object} map[string]string "error": "Internal Server Error"
// @Router /user [get]
func GetUser(c *gin.Context, user usecase.IUserGetter) {
	userID, _ := c.Get("userID")
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID is not a string"})
		return
	}

	resp, err := user.GetUserByID(c, userIDStr)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": resp})
}
