package controller

import (
	"github.com/boliev/coach/internal/request"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u User) Create(c *gin.Context) {
	var request request.UserCreation
	c.ShouldBind(&request)
	c.JSON(200, gin.H{"result": "User was created"})
}
