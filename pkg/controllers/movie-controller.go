package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/allenkays/special-garbanzo/pkg/models"
	"github.com/allenkays/special-garbanzo/pkg/utils"
	"github.com/gorilla/mux"
)

var NewMovie models.Movie

func GetMovie(w http.ResponseWriter, r *http.Request) {
	newMovies := models.GetAllMovies()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, _ := models.GetMovieById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	b := CreateMovie.CreateMovie()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error trying to parse")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movie{}
	utils.ParseBody(r, updateMovie)

	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	movieDetails, db := models.GetMovieById(ID)

	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateMovie.Director != "" {
		movieDetails.Director = updateMovie.Director
	}
	if updateMovie.Production != "" {
		movieDetails.Production = updateMovie.Production
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
