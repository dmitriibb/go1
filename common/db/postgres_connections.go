package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "qwerty1"
	dbname   = "go1"
)

var ()

func TestConnectPostgres() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func UseConnection(f func(db *sql.DB) any) any {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("db connection open")
	defer func() {
		connection.Close()
		fmt.Println("db connection close")
	}()
	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	return f(connection)
}
