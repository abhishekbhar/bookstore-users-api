package services


import (
	"github.com/abhishekbhar/bookstore-users-api/domain/users"
	"github.com/abhishekbhar/bookstore-users-api/utils/errors"
	"github.com/abhishekbhar/bookstore-users-api/utils/crypto_utils"
)


var (
	UsersService UsersServiceInterface = &usersService{}
)


type usersService struct {
}

type UsersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user,nil
}


func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr){
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	user := &users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}
	
	return user,nil
}


func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if !isPartial {
		currentUser.FirstName = user.FirstName
		currentUser.LastName  = user.LastName
		currentUser.Email     = user.Email
	} else {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName  = user.LastName
		}
		if user.Email != "" {
			currentUser.Email     = user.Email
		}
	}

	if err = currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
	
}



func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	currentUser, err := UsersService.GetUser(userId)
	if err != nil {
		return err
	}

	if err = currentUser.Delete(); err != nil {
		return err
	}

	return nil	
}



func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)	
}


func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil

}