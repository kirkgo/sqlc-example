package main

import (
	"context"
	"database/sql"
	"log"

	postgres "github.com/kirkgo/sqlc-example/postgres"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "user=postgres host=/tmp dbname=sqlc-example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.New(conn)

	user, err := db.CreateUser(context.Background(), postgres.CreateUserParams{
		Firstname: "John",
		Lastname:  "Doe",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}
