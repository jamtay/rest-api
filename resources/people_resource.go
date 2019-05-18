package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	. "rest-api/config"
	peopleDAO "rest-api/dao"
	"rest-api/helpers"
	"rest-api/models"
)

var peopleDao = peopleDAO.PeopleDAO{}
var peopleConfig = Config{}

func init() {
	peopleConfig.ReadPeopleConfig()
	peopleDao.Server = peopleConfig.Server
	peopleDao.Database = peopleConfig.Database
	peopleDao.Connect()
}


//https://stackoverflow.com/questions/13255907/in-go-http-handlers-why-is-the-responsewriter-a-value-but-the-request-a-pointer
func GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := peopleDao.FindAllPeople()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, people)
	//json.NewEncoder(w).Encode(peopleDAO.FindAllPeople())
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person, err := peopleDao.FindById(params["id"])
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, person)
	//json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	//Example person: Person{ID: "1", Firstname: "James", Lastname: "Taylor", Address: &Address{City: "London", State: "England"}}
	defer r.Body.Close()
	var person models.Person


	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	person.ID = bson.NewObjectId()
	if err := peopleDao.Insert(person); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusCreated, person)
	//json.NewEncoder(w).Encode(dao.FindAllPeople())
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := peopleDao.Delete(person); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}