package coach

import (
	"github.com/boliev/coach/internal/controller"
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (app App) Start() {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			controller := &controller.User{}
			users.POST("/", controller.Create)
			users.GET("/", controller.List)
		}
	}
	r.Run()
}
