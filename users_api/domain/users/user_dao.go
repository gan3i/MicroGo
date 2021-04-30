package users

import (
	"fmt"
	"strings"

	"github.com/gan3i/microgo/datasource/mysql/users_db"
	"github.com/gan3i/microgo/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	querySelectUser = "SELECT * FROM users WHERE id = ?"
	queryDeleteUser = "DELETE FROM users WHERE id = ?"
	queryUpdateUser = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
)

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	//alternate but slow way, accordig to benchmarks
	//result, err := users_db.Client.Exec(queryInsertUser,u.FirstName, u.LastName, u.Email, u.CreatedDate)

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.CreatedDate)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			return errors.NewInternalServerError(fmt.Sprintf("Email Id %s is in use by another account", u.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to  get user ID: %s", err.Error()))
	}

	u.Id = userId
	return nil
}

func (u *User) Get() *errors.RestErr {

	stmt, prepareErr := users_db.Client.Prepare(querySelectUser)
	if prepareErr != nil {
		return errors.NewInternalServerError("Error while fetching user : " + prepareErr.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)

	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreatedDate); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError(fmt.Sprintf("user %d does not exist", u.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user  %d : %s", u.Id, err.Error()))
	}
	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, prepareErr := users_db.Client.Prepare(queryDeleteUser)
	if prepareErr != nil {
		return errors.NewInternalServerError("Error while deleting the user")
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(u.Id)

	if deleteErr != nil {
		return errors.NewInternalServerError(deleteErr.Error())
	}
	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, prepareErr := users_db.Client.Prepare(queryUpdateUser)
	if prepareErr != nil {
		return errors.NewInternalServerError("error while updating the user")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)

	if updateErr != nil {
		return errors.NewInternalServerError("error exce : " + updateErr.Error())
	}
	return nil
}

// func Get(userId int64)(*User, *errors.RestErr){
// 	return  nil, nil
// }

// func Save(user User) *errors.RestErr {
// 	return nil
// }
