package users_store

import (
	"fmt"
	"testing"
)

var (
	userStoreTest   UserStore
	userServiceTest UserService
	err             error
	username        string = "yerassyl"
	password        string = "passanya"
	email           string = "tleugazy98@gmail.com"
	fullName        string = "Tleugazy Yerassyl"
)

func TestNewMongoStore(t *testing.T) {
	mongoConfig := MongoConfig{
		Host:     "localhost",
		Port:     "27017",
		Database: "recommendation_system",
	}
	userStoreTest, err = NewMongoStore(mongoConfig)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("test %s completed \n", t.Name())
}

func TestNewUserService(t *testing.T) {
	userServiceTest = NewUserService(userStoreTest)
	fmt.Printf("test %s completed \n", t.Name())
}

func TestUserService_CreateUser(t *testing.T) {
	cmd := &CreateUserCommand{
		Username: username,
		Email:    email,
		Password: password,
		FullName: fullName,
	}
	_, err := userServiceTest.CreateUser(cmd)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("test %s completed \n", t.Name())
}

func TestUserService_GetUserByUsername(t *testing.T) {
	cmd := &GetUserByUsername{
		Username: "12322",
	}
	user, err := userServiceTest.GetUserByUsername(cmd)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("test %s completed \n", t.Name())
	fmt.Printf("user data -> %s", user)

}
