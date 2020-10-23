package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	clientrepository "../repository/client"
	"github.com/gorilla/mux"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Controller to be called from main
type Controller struct{}

var clients []models.Client

//GetClients controller
func (c Controller) GetClients(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var client models.Client

		clients = []models.Client{}
		clientRepo := clientrepository.ClientRepository{}
		clients = clientRepo.GetClients(db, client, clients)

		json.NewEncoder(w).Encode(clients)
	}
}

//GetClient controller
func (c Controller) GetClient(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var client models.Client
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		clientRepo := clientrepository.ClientRepository{}

		client = clientRepo.GetClient(db, client, id)

		json.NewEncoder(w).Encode(client)
	}
}

//AddClient controller
func (c Controller) AddClient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var client models.Client
		json.NewDecoder(r.Body).Decode(&client)

		clientRepo := clientrepository.ClientRepository{}
		id := clientRepo.AddClient(db, client)

		json.NewEncoder(w).Encode(id)
	}
}

//UpdateClient controller
func (c Controller) UpdateClient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var client models.Client

		json.NewDecoder(r.Body).Decode(&client)

		clientRepo := clientrepository.ClientRepository{}
		rowsUpdated := clientRepo.UpdateClient(db, client)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

//RemoveClient controller
func (c Controller) RemoveClient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		clientRepo := clientrepository.ClientRepository{}
		rowsDeleted := clientRepo.RemoveClient(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
