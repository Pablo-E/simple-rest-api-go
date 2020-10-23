package main

import (
	"log"
	"net/http"

	"./controllers"
	"./driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	_ "github.com/lib/pq"
)

func init() {
	gotenv.Load()
}

func main() {

	db := driver.GetConnection()

	controller := controllers.Controller{}

	router := mux.NewRouter()
	router.HandleFunc(`/clients`, controller.GetClients(db)).Methods(`GET`)
	router.HandleFunc(`/clients/{id}`, controller.GetClient(db)).Methods(`GET`)
	router.HandleFunc(`/clients`, controller.AddClient(db)).Methods(`POST`)
	router.HandleFunc(`/clients`, controller.UpdateClient(db)).Methods(`PUT`)
	router.HandleFunc(`/clients/{id}`, controller.RemoveClient(db)).Methods(`DELETE`)

	log.Fatal(http.ListenAndServe(`:8000`, router))

}
