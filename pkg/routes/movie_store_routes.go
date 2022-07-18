package routes

import (
	"github.com/allenkays/special-garbanzo/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controllers.DeleteMovie).Methods("DELETE")
}
