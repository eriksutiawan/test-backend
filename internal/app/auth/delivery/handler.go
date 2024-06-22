package delivery

import (
	"net/http"
	"test-backend/internal/app/auth/usecase"
	"test-backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, password and email
// @Tags auth
// @Accept json
// @Produce json
// @Param user body usecase.AuthRegisterDto true "User registration details"
// @Success 200 {object} map[string]string "message": "User created successfully"
// @Failure 400 {object} map[string]string "error": "Bad Request"
// @Failure 500 {object} map[string]string "error": "Internal Server Error"
// @Router /register [post]
func Register(c *gin.Context, auth usecase.IAuth) {
	var user usecase.AuthRegisterDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := auth.Register(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login godoc
// @Summary Login a user
// @Description Login a user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body usecase.LoginDto true "User login details"
// @Success 200 {object} usecase.UserResponse "token": "access_token", "expired_token": "token expiration in Unix timestamp"
// @Failure 400 {object} map[string]string "error": "Bad Request"
// @Failure 401 {object} map[string]string "error": "Unauthorized"
// @Router /login [post]
func Login(c *gin.Context, auth usecase.IAuth) {
	var user usecase.LoginDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := auth.Login(c, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password or username"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp})
}

// GoogleLogin godoc
// @Summary Initiates Google OAuth login
// @Description Redirects user to Google OAuth login page
// @Tags auth
// @Produce json
// @Success 307 "Temporary Redirect"
// @Router /auth/google [get]
func GoogleLogin(c *gin.Context) {
	url := utils.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback godoc
// @Summary Callback endpoint for Google OAuth
// @Description Handles Google OAuth callback and exchanges code for access token
// @Tags auth
// @Produce json
// @Param code query string true "Authorization code received from Google"
// @Success 200 {object} map[string]string "message": "User authenticated successfully"
// @Failure 500 {object} map[string]string "error": "Failed to exchange token" or "Failed to get user info"
// @Router /auth/google/callback [get]
func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := utils.GoogleOAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()
	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
}
