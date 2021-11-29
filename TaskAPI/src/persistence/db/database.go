package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "taskapp"
	password = "taskapp"
	dbname   = "taskapp"
)

type Database struct {
	connection       *sql.DB
	connectionString string
}

func CreateDb() *Database {
	database := &Database{}
	database.connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return database
}

func (d *Database) Connect() {
	db, err := sql.Open("postgres", d.connectionString)
	if err != nil {
		panic(err)
	}
	d.connection = db
	fmt.Println("Database connected")
}

func (d *Database) GetDbConnection() *sql.DB {
	return d.connection
}

func (d *Database) Disconnect() {
	d.connection.Close()
	fmt.Println("Database disconnected")
}
