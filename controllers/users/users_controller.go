package users


import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/abhishekbhar/bookstore-users-api/domain/users"
	"github.com/abhishekbhar/bookstore-users-api/services"
)


func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		return
	}

	result, saveError := services.CreateUser(user)

	if saveError != nil {
		//TODO: Handle User creation Error
		return
	}

	c.JSON(http.StatusCreated, result)
	// c.String(http.StatusNotImplemented, "implement me!")

}
	

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}


func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}