package users_store

type ListMoviesCommand struct {
	Count int64 `json:"count,omitempty"`
}

func(cmd *ListMoviesCommand) Exec(service UserService) (interface{},error) {
	return service.ListMovies(cmd)
}

type GetMovieByNameCommand struct {
	Name string `json:"name"`
}

func(cmd *GetMovieByNameCommand) Exec(service UserService) (interface{},error) {
	return service.GetMovieByName(cmd)
}