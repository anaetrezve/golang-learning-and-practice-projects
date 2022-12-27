package routes

import (
	"github.com/anaetrezve/book-management-system/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
