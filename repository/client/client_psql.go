package clientrepository

import (
	"database/sql"
	"log"

	"../../models"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//ClientRepository to be called
type ClientRepository struct{}

var clients []models.Client

//GetClients controller
func (c ClientRepository) GetClients(db *sql.DB, client models.Client, clients []models.Client) []models.Client {

	rows, err := db.Query(`SELECT * FROM public.clients`)

	logFatal(err)

	defer rows.Close()

	for rows.Next() {

		var client models.Client

		err = rows.Scan(&client.ID, &client.Name, &client.Lastname)

		logFatal(err)

		clients = append(clients, client)
	}

	return clients
}

//GetClient controller
func (c ClientRepository) GetClient(db *sql.DB, client models.Client, id int) models.Client {

	rows := db.QueryRow(`SELECT * FROM public.clients WHERE id=$1`, id)

	err := rows.Scan(&client.ID, &client.Name, &client.Lastname)

	logFatal(err)

	return client
}

//AddClient controller
func (c ClientRepository) AddClient(db *sql.DB, client models.Client) int {

	err := db.QueryRow(`INSERT INTO public.clients (id, name, lastname) values($1, $2, $3) RETURNING id`,
		client.ID, client.Name, client.Lastname).Scan(&client.ID)

	logFatal(err)

	return client.ID
}

//UpdateClient controller
func (c ClientRepository) UpdateClient(db *sql.DB, client models.Client) int64 {

	result, err := db.Exec(`UPDATE  public.clients SET name=$1, lastname=$2 WHERE ID=$3`,
		&client.Name, &client.Lastname, &client.ID)
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

//RemoveClient controller
func (c ClientRepository) RemoveClient(db *sql.DB, id int) int64 {

	result, err := db.Exec(`DELETE FROM public.clients WHERE id=$1`, id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
