package users

import (
	"fmt"
	"strings"
	"github.com/abhishekbhar/bookstore-users-api/utils/errors"
	"github.com/abhishekbhar/bookstore-users-api/utils/date_utils"
	"github.com/abhishekbhar/bookstore-users-api/datasources/mysql/users_db"
)


const (
	errNoRow	          = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=? , last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

var (
	usersDB = make(map[int64]*User)
)


func(user *User) Get()  *errors.RestErr{
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	} 

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()


	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()),
		)
	}


	return nil
}

func (user *User) Save() *errors.RestErr {

	
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email,user.Status, user.Password, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error trying to save user: %s", err.Error()))
	}


	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error trying to save user: %s", err.Error()))
	}


	user.Id = userId
	return nil
}


func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
    defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("error when updating user %d: %s", user.Id, err.Error()))
	}

	return nil
}


func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
    defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("error when updating user %d: %s", user.Id, err.Error()))
	}

	return nil
}


func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No user macthing status %s", status))
	}
	return results, nil

}  