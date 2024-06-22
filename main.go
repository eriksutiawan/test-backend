package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	authdelivery "test-backend/internal/app/auth/delivery"
	authusecase "test-backend/internal/app/auth/usecase"
	predictiondelivery "test-backend/internal/app/prediction/delivery"
	predictionrepo "test-backend/internal/app/prediction/repository"
	predictionusecase "test-backend/internal/app/prediction/usecase"
	userdelivery "test-backend/internal/app/user/delivery"
	userrepo "test-backend/internal/app/user/repository"
	userusecase "test-backend/internal/app/user/usecase"
	"test-backend/internal/pkg/middleware"
	"test-backend/internal/pkg/utils"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a documentation for test backend.

// @host      localhost:8080
// @BasePath  /api
func main() {
	// Load environment variables
	utils.LoadEnv()

	// Create a Gin router
	r := gin.Default()

	// Connect to MongoDB
	db := utils.ConnectToMongo()

	// Initialize repositories
	userRepository := userrepo.NewUserRepository(db)
	predictionRepository := predictionrepo.NewPredictionRepository(db)

	// Initialize use cases/services
	userService := userusecase.NewUser(userRepository)
	authService := authusecase.NewAuth(userService)
	predictionService := predictionusecase.NewPrediction(predictionRepository)

	// Auth routes group
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", func(c *gin.Context) {
			authdelivery.Register(c, authService)
		})
		authGroup.POST("/login", func(c *gin.Context) {
			authdelivery.Login(c, authService)
		})
		authGroup.POST("/google/login", authdelivery.GoogleLogin)
		authGroup.GET("/google/callback", authdelivery.GoogleCallback)
	}

	// User routes group
	userGroup := r.Group("/api/user")
	{
		userGroup.GET("/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
			userdelivery.GetUser(c, userService)
		})
	}

	// Prediction routes group
	predictionGroup := r.Group("/api/predict")
	{
		predictionGroup.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			predictiondelivery.Predict(c, predictionService)
		})
	}

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run the server on port 8080
	r.Run(":8080")
}
