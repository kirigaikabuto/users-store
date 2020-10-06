package users_store

import (
	"bytes"
	"encoding/json"
	"fmt"
	movie_store "github.com/kirigaikabuto/movie-store"
	"io/ioutil"
	"net/http"
)

type UserService interface {
	ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error)
	GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie,error)
	GetMovieById(cmd *GetMovieByIdCommand) (*MovieRecommendResponse, error)
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

func (svc *userService) GetMovieById(cmd *GetMovieByIdCommand) (*MovieRecommendResponse, error){
	movie, err := svc.amqpRequests.GetMovieById(cmd)
	if err != nil{
		return nil, err
	}
	requestBody ,err := json.Marshal(map[string]string{
		"title":movie.Name,
	})
	if err != nil {
		return nil, err
	}
	resp , err := http.Post("http://127.0.0.1:5000/api/v1/recommend/","application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	output := &MovieRecommendResponse{}
	fmt.Println(string(body))
	err = json.Unmarshal(body,&output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

