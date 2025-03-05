package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/thanh2k4/Chat-app/application"
)

func main() {
	db := "postgres://postgres:0@localhost:5432/chat_app?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), db)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	defer conn.Close(context.Background())
	fmt.Println("Connected to database")
	app := application.New()
	err = app.Start(context.TODO())
	if err != nil {
		panic(err)
	}

}
