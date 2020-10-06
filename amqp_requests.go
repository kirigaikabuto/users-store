package users_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	movie_store "github.com/kirigaikabuto/movie-store"
)
type AmqpRequests struct {
	clt amqp.Client
}

func NewAmqpRequests(clt amqp.Client) *AmqpRequests{
	return &AmqpRequests{
		clt: clt,
	}
}

func (r *AmqpRequests) GetListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie,error){
	response, err := r.call("movie.list",cmd)
	if err != nil{
		return nil, err
	}
	var movies []movie_store.Movie
	err = json.Unmarshal(response.Body, &movies)
	if err != nil {
		return nil, err
	}
	return movies,nil
}

func (r *AmqpRequests) GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie,error){
	response, err := r.call("movie.getByName",cmd)
	if err != nil{
		return nil, err
	}
	var movie *movie_store.Movie
	err = json.Unmarshal(response.Body, &movie)
	if err != nil {
		return nil, err
	}
	return movie,nil
}

func (r *AmqpRequests) call(path string, data interface{}) (*amqp.Message,error){
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	respone,err := r.clt.Call(path, amqp.Message{
		Body: jsonData,
	})
	return respone,nil
}