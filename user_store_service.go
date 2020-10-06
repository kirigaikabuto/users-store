package users_store

import (
	movie_store "github.com/kirigaikabuto/movie-store"
)
type UserService interface {
	ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie,error)

}

type userService struct {
	amqpRequests AmqpRequests
}

func NewUserService(amqpReq AmqpRequests) UserService{
	return &userService{
		amqpReq,
	}
}

func (svc *userService) ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie,error) {
	movies, err := svc.ListMovies(cmd)
	if err != nil {
		return nil, err
	}
	return movies,nil
}
