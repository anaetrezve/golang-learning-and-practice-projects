package main

import (
	"net/http"

	"github.com/anaetrezve/golang-mongodb-rest-api/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/users/:ID", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.DELETE("/users/:ID", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return s
}
