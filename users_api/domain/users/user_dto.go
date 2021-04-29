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

func (u User) Validate() *errors.RestErr {
	if strings.TrimSpace(u.Email) == "" {
		return errors.NewBadRequestError("Invalid email Id")
	}
	return nil
}
