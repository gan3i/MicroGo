package users

import (
	"strings"

	"github.com/gan3i/microgo/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_date"`
}

func (u *User) Validate() *errors.RestErr {
	if strings.TrimSpace(u.Email) == "" {
		return errors.NewBadRequestError("Invalid email Id")
	}
	return nil
}

func (u *User) MapValues(other *User) *errors.RestErr {
	if strings.TrimSpace(other.FirstName) != "" {
		u.FirstName = other.FirstName
	}
	if strings.TrimSpace(other.LastName) != "" {
		u.LastName = other.LastName
	}
	if strings.TrimSpace(other.Email) != "" {
		u.Email = other.Email
	}
	return nil
}
