package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	. "rest-api/config"
	. "rest-api/dao"
	"rest-api/helpers"
	"rest-api/models"
)

var dao = MoviesDAO{}
var moviesConfig = Config{}

func init() {
	moviesConfig.ReadMoviesConfig()
	dao.Server = moviesConfig.Server
	dao.Database = moviesConfig.Database
	dao.Connect()
}

func AllMoviesEndPoint(w http.ResponseWriter, jr *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, movies)
}

func FindMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, movie)
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusCreated, movie)
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := dao.Delete(movie); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}