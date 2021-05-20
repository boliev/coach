package coach

import (
	"github.com/boliev/coach/internal/controller"
	"github.com/boliev/coach/internal/middleware"
	"github.com/boliev/coach/internal/mongo"
	"github.com/boliev/coach/internal/user"
	"github.com/boliev/coach/pkg/config"
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (app App) Start() {
	config := getConfig()
	mongoClient := mongo.NewClient(config.GetString("mongo_uri"))
	userRepository := mongo.NewUserMongoRepository(
		mongoClient,
		config.GetString("main_database"),
		config.GetString("users_collection"),
	)
	jwtCreator := user.NewJwtCreator(config.GetString("jwt_secret"), config.GetInt("jwt_days"))
	userService := user.CreateUserService(userRepository, jwtCreator)
	authHandler := middleware.NewAuthHandler(jwtCreator)

	r := gin.New()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			userController := &controller.User{
				UserRepository: userRepository,
				UserService:    userService,
			}

			users.POST("/", userController.Create)
			users.POST("/auth", userController.Auth)
			users.GET("/", authHandler.Handle, userController.List)
			users.GET("/:id", userController.One)
			users.DELETE("/:id", userController.Delete)
		}
	}
	r.Run()
}

func getConfig() *config.Config {
	config, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}
