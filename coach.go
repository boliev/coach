package coach

import (
	"github.com/boliev/coach/internal/controller"
	"github.com/boliev/coach/internal/mongo"
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (app App) Start() {
	mongoClient := mongo.NewClient("mongodb://coach:123456@localhost:27017")
	userRepository := mongo.NewUserMongoRepository(mongoClient, "coach", "user")

	r := gin.New()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			userController := &controller.User{
				UserRepository: userRepository,
			}

			users.POST("/", userController.Create)
			users.GET("/", userController.List)
			users.GET("/:id", userController.One)
			users.DELETE("/:id", userController.Delete)
		}
	}
	r.Run()
}
