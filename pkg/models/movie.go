package models

import (
	"github.com/allenkays/special-garbanzo/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name       string `gorm: ""json:"name"`
	Director   string `json: "author"`
	Production string `json: "production"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}
func (b *Movie) CreateMovie() *Movie {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}

func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}
