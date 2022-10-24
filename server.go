package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibnu-a/Golang-RESTful-API/config"
	"github.com/ibnu-a/Golang-RESTful-API/controller"
	"github.com/ibnu-a/Golang-RESTful-API/repository"
	"github.com/ibnu-a/Golang-RESTful-API/service"
	"gorm.io/gorm"
)

//global variable
var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWtService        = service.NewJwtService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
