package users


import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/abhishekbhar/bookstore-users-api/domain/users"
	"github.com/abhishekbhar/bookstore-users-api/services"
	"github.com/abhishekbhar/bookstore-users-api/utils/errors"
)


func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		return 0,err
	}
	return userId, nil;
}


func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.UsersService.CreateUser(user)

	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-Public") == "true"))

}
	

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.UsersService.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-Public") == "true"))
}


func Update(c *gin.Context) {
	var user users.User

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userId

	updatedUser, userUpdateErr := services.UsersService.UpdateUser(isPartial, user)
	if userUpdateErr != nil {
		c.JSON(userUpdateErr.Status, userUpdateErr)
		return
	}

	c.JSON(http.StatusOK, updatedUser.Marshall(c.GetHeader("x-Public") == "true"))
}

func Delete(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err = services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}


func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-Public") == "true"))
}