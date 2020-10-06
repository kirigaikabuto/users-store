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

type GetMovieByIdCommand struct {
	Id int64 `json:"id"`
}

func(cmd *GetMovieByIdCommand) Exec(service UserService) (interface{},error) {
	return service.GetMovieById(cmd)
}

type MovieRecommend struct {
	Name string `json:"name"`
	Score float64 `json:"score"`
}

type MovieRecommendResponse struct {
	 Movies []*MovieRecommend `json:"result"`
}

