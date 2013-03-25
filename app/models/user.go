package models

import (
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"regexp"
)

type User struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	Firstname      string        `Firstname`
	Lastname       string        `Lastname`
	Email          string        `Email`
	HashedPassword []byte        `HashedPassword`
	Password       string
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

var nameRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Firstname,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{64},
		revel.Match{nameRegex},
	)

	v.Check(user.Lastname,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{64},
		revel.Match{nameRegex},
	)

	v.Check(user.Email,
		revel.Required{},
	)

	v.Email(user.Email)

	ValidatePassword(v, user.Password).
		Key("user.Password")

}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MinSize{8},
	)
}