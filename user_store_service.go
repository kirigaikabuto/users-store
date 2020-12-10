package users_store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	movie_store "github.com/kirigaikabuto/movie-store"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"time"
)

type UserService interface {
	ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error)
	GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie, error)
	GetMovieById(cmd *GetMovieByIdCommand) ([]movie_store.Movie, error)
	CreateUser(cmd *CreateUserCommand) (*User, error)
}

type userService struct {
	amqpRequests AmqpRequests
	userStore    UserStore
}

func NewUserService(amqpReq AmqpRequests, userStore UserStore) UserService {
	return &userService{
		amqpReq,
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
	user.FullName = cmd.FullName
	user.Username = cmd.Username
	hash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), 5)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.RegisterDate = time.Now()
	return user, nil
}

func (svc *userService) ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error) {
	movies, err := svc.amqpRequests.GetListMovies(cmd)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (svc *userService) GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie, error) {
	movie, err := svc.amqpRequests.GetMovieByName(cmd)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (svc *userService) GetMovieById(cmd *GetMovieByIdCommand) ([]movie_store.Movie, error) {
	movie, err := svc.amqpRequests.GetMovieById(cmd)
	if err != nil {
		return nil, err
	}
	requestBody, err := json.Marshal(map[string]string{
		"title": movie.Name,
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("http://127.0.0.1:5000/api/v1/recommend/", "application/json", bytes.NewBuffer(requestBody))
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
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	var movies []movie_store.Movie
	for _, v := range output.Movies {
		newMovie, err := svc.amqpRequests.GetMovieByName(&GetMovieByNameCommand{
			v.Name,
		})
		if err != nil {
			return nil, err
		}
		newMovie.Score = v.Score
		movies = append(movies, *newMovie)
	}
	return movies, nil
}
