package users_store

import (
	movie_store "github.com/kirigaikabuto/movie-store"
)

type UserService interface {
	ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error)
	GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie,error)
	GetMovieById(cmd *GetMovieByIdCommand) (*movie_store.Movie, error)
}

type userService struct {
	amqpRequests AmqpRequests
}

func NewUserService(amqpReq AmqpRequests) UserService {
	return &userService{
		amqpReq,
	}
}

func (svc *userService) ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error) {
	movies, err := svc.amqpRequests.GetListMovies(cmd)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (svc *userService) GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie,error) {
	movie, err := svc.amqpRequests.GetMovieByName(cmd)
	if err != nil{
		return nil, err
	}
	return movie, nil
}

func (svc *userService) GetMovieById(cmd *GetMovieByIdCommand) (*movie_store.Movie, error){
	movie, err := svc.amqpRequests.GetMovieById(cmd)
	if err != nil{
		return nil, err
	}
	return movie, nil
}

