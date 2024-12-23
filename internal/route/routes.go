package route

import (
	"tspo_final/internal/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	userRepository := repository.NewUserRepository(db)
	SetupUsersRoutes(userRepository, route)

	orderRepository := repository.NewOrderRepository(db)
	SetupOrdersRoutes(orderRepository, route)

	SetupAuthRoutes(userRepository, route)

	route.GET("/swagger/*any", FixRequestUri)

	return route
}

func FixRequestUri(c *gin.Context) {
	if c.Request.RequestURI == "" {
		c.Request.RequestURI = c.Request.URL.Path
	}

	ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
}
