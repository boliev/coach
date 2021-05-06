package controller

import (
	"fmt"
	"net/http"

	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/internal/request"
	"github.com/boliev/coach/internal/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserRepository domain.UserRepository
}

func (u User) Create(c *gin.Context) {
	var request request.UserCreation
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequest("bad request"))
		return
	}
	// repository := repository.NewUserMongoRepository()
	user := request.ToDomain()

	res, err := u.UserRepository.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequest(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("User %s was created. %s", request.Email, res)})
}

func (u User) List(c *gin.Context) {
	users, _ := u.UserRepository.FindAll()

	c.JSON(http.StatusOK, response.CreateUsersListFromDomain(users))
}

func (u User) One(c *gin.Context) {
	id := c.Param("id")
	user, _ := u.UserRepository.Find(id)

	c.JSON(http.StatusOK, response.CreateUserFromDomain(user))
}

func (u User) Delete(c *gin.Context) {
	id := c.Param("id")
	u.UserRepository.Delete(id)

	c.JSON(http.StatusOK, "deleted")
}
