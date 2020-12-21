package users_store

import (
	cmn_lib "github.com/kirigaikabuto/common-lib"
	"time"
)

type User struct {
	Id           string               `json:"id"`
	Username     string               `json:"username"`
	Password     string               `json:"password"`
	FullName     string               `json:"full_name"`
	PhoneNumber  string               `json:"phone_number"`
	Email        string               `json:"email"`
	RegisterType cmn_lib.RegisterType `json:"register_type"`
	RegisterDate time.Time            `json:"register_date"`
}

type UserUpdate struct {
	Id          int64   `json:"id"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	FullName    *string `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}
