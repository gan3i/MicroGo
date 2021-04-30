package user_service

import (
	"github.com/gan3i/microgo/domain/users"
	"github.com/gan3i/microgo/utils/errors"
)

//error should at he end
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		Id: userId,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{
		Id: userId,
	}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {

	oUser, err := GetUser(user.Id)

	if err != nil {
		return nil, err
	}

	if mapErr := oUser.MapValues(&user); mapErr != nil {
		return nil, mapErr
	}

	if updateErr := oUser.Update(); updateErr != nil {
		return nil, updateErr
	}

	return oUser, nil

}
