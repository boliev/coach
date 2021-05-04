package controller

import (
	"fmt"
	"net/http"

	"github.com/boliev/coach/internal/request"
	"github.com/boliev/coach/internal/response"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u User) Create(c *gin.Context) {
	var request request.UserCreation
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequest("bad request"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("User %s was created", request.Email)})
}
