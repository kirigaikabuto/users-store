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
	GetUserByUsername(cmd *GetUserByUsername) (*User, error)
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
	existingUser, err := svc.userStore.GetByUsername(user.Username)
	if err != nil {
		if err.Error() != "no user by this username" {
			return nil, err
		}
	}
	if existingUser != nil {
		return nil, errors.New("user with that username already exist")
	}
	if user.Email != "" {
		existingUser, err = svc.userStore.GetByEmail(user.Email)
		if err != nil {
			if err.Error() != "no user by this email" {
				return nil, err
			}
		}
		if existingUser != nil {
			return nil, errors.New("user with that email already exist")
		}
	}
	if user.PhoneNumber != "" {
		existingUser, err = svc.userStore.GetByPhone(user.PhoneNumber)
		if err != nil {
			if err.Error() != "no user by this phone" {
				return nil, err
			}
		}
		if existingUser != nil {
			return nil, errors.New("user with that phone already exist")
		}
	}

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

func (svc *userService) GetUserByUsername(cmd *GetUserByUsername) (*User, error) {
	if cmd.Username == "" {
		return nil, errors.New("please enter the username")
	}
	user, err := svc.userStore.GetByUsername(cmd.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
