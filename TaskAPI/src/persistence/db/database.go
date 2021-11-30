package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	user     = "taskapp"
	password = "taskapp"
	dbname   = "taskapp"
)

var host string = os.Getenv("DB_HOST")

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
	for {
		err := d.connection.Ping()
		if err == nil {
			break
		}
	}
	fmt.Println("Database ready to accept communication!")
}

func (d *Database) GetDbConnection() *sql.DB {
	return d.connection
}

func (d *Database) Disconnect() {
	d.connection.Close()
	fmt.Println("Database disconnected")
}
