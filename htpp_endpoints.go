package users_store

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type HttpEndpointsFactory interface {
	ListMoviesEndpoint() func(w http.ResponseWriter, r *http.Request)
	GetMovieByNameEndpoint(nameParam string) func(w http.ResponseWriter, r *http.Request)
}

type httpEndpointsFactory struct {
	userService UserService
}

type customError struct {
	Message string `json:"message"`
}

func NewHttpEndpoints(userService UserService) HttpEndpointsFactory {
	return &httpEndpointsFactory{userService: userService}
}

func (httpFac *httpEndpointsFactory) ListMoviesEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		listMovieReq := &ListMoviesCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(listMovieReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		count, err := strconv.ParseInt(r.URL.Query().Get("count"),10,64)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		listMovieReq.Count = count
		data, err := listMovieReq.Exec(httpFac.userService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusCreated, data)
	}
}

func (httpFac *httpEndpointsFactory) GetMovieByNameEndpoint(nameParam string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		movieReq := &GetMovieByNameCommand{}
		vars := mux.Vars(r)
		name, ok := vars[nameParam]
		if !ok {
			respondJSON(w, http.StatusInternalServerError, &customError{"no token param"})
			return
		}
		movieReq.Name = name
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(movieReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := movieReq.Exec(httpFac.userService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusCreated, data)
	}
}


func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}