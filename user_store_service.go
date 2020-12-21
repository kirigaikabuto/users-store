package users_store

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	cmn_lib "github.com/kirigaikabuto/common-lib"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	CreateUser(cmd *CreateUserCommand) (*User, error)
}

type userService struct {
	userStore UserStore
}

func NewUserService(userStore UserStore) UserService {
	return &userService{
		userStore,
	}
}

func (svc *userService) CreateUser(cmd *CreateUserCommand) (*User, error) {
	if cmd.Email == "" && cmd.PhoneNumber == "" {
		return nil, errors.New("email or phone_number should be included")
	} else if cmd.FullName == "" {
		return nil, errors.New("full name should be included")
	} else if cmd.Password == "" {
		return nil, errors.New("password should be included")
	} else if cmd.Username == "" {
		return nil, errors.New("username should be included")
	}
	user := &User{}
	if cmd.Email != "" && cmd.PhoneNumber != "" {
		user.Email = cmd.Email
		user.PhoneNumber = cmd.PhoneNumber
	} else if cmd.Email != "" {
		user.Email = cmd.Email
	} else if cmd.PhoneNumber != "" {
		user.PhoneNumber = cmd.PhoneNumber
	}
	if cmd.Email != "" {
		user.RegisterType = cmn_lib.Email
	} else if cmd.PhoneNumber != "" {
		user.RegisterType = cmn_lib.Phone
	}
	user.FullName = cmd.FullName
	user.Username = cmd.Username
	hash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), 5)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.RegisterDate = time.Now()
	uuId := uuid.New()
	user.Id = uuId.String()
	newUser, err := svc.userStore.Create(user)
	fmt.Println(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
