package controller

import (
	"fmt"
	"net/http"

	"github.com/boliev/coach/internal/repository"
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
	repository := repository.NewUserMongoRepository()
	user := request.ToDomain()

	res, err := repository.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequest(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("User %s was created. %s", request.Email, res)})
}

func (u User) List(c *gin.Context) {
	repository := repository.NewUserMongoRepository()

	users, _ := repository.FindAll()

	c.JSON(http.StatusOK, response.UsersList{
		Data:  users,
		Count: len(users),
	})
}

func (u User) One(c *gin.Context) {
	repository := repository.NewUserMongoRepository()
	id := c.Param("id")
	user, _ := repository.Find(id)

	c.JSON(http.StatusOK, user)
}

func (u User) Delete(c *gin.Context) {
	repository := repository.NewUserMongoRepository()
	id := c.Param("id")
	repository.Delete(id)

	c.JSON(http.StatusOK, "deleted")
}
