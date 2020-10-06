package users_store

type ListMoviesCommand struct {
	Count int64 `json:"count"`
}

func(cmd *ListMoviesCommand) Exec(service UserService) (interface{},error) {
	return
}