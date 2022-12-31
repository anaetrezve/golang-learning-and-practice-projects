package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anaetrezve/golang-mongodb-rest-api/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("ID")

	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(ID)

	u := models.User{}

	if err := uc.Session.DB("mongo-golang").C("Users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = bson.NewObjectId()

	uc.Session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("ID")

	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(ID)

	if err := uc.Session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid, "\n")
}
