package controllers
import (
	"github.com/gorilla/mux"
	"github.com/saif101/go-bookstore/pkg/controllers"
)
var RegisterBookstoreControllers= func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreatBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.handleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}