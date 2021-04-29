package users

import (
	"fmt"

	"github.com/gan3i/microgo/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (u *User) Save() *errors.RestErr {

	if userDB[u.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exist", u.Id))
	}
	userDB[u.Id] = u
	return nil
}
func (u *User) Get() *errors.RestErr {
	result := userDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", u.Id))
	}
	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.CreatedDate = result.CreatedDate
	return nil
}

// func Get(userId int64)(*User, *errors.RestErr){
// 	return  nil, nil
// }

// func Save(user User) *errors.RestErr {
// 	return nil
// }
