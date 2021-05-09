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

	result, saveError := services.CreateUser(user)

	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)

}
	

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}

	c.JSON(http.StatusOK, user)
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

	updatedUser, userUpdateErr := services.UpdateUser(isPartial, user)
	if userUpdateErr != nil {
		c.JSON(userUpdateErr.Status, userUpdateErr)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func Delete(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err = services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}


func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}