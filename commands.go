package users_store

type CreateUserCommand struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (cmd *CreateUserCommand) Exec(service UserService) (interface{}, error) {
	return service.CreateUser(cmd)
}

type GetUserByUsername struct {
	Username string `json:"username"`
}

func (cmd *GetUserByUsername) Exec(service UserService) (interface{}, error) {
	return service.GetUserByUsername(cmd)
}
