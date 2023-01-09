package main

import (
	"bookstore/pkg/db"
	"bookstore/pkg/server"
	"log"
)

func main() {
	conn, err := db.New("localhost", "5432", "user", "password", "bookstore")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	srv := server.New(conn)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

