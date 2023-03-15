package main

import (
	"database/sql"
	"github.com/EdsonGustavoTofolo/desafio-client-server-api-golang/internal/infra/controllers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	db = initDatabase()
	defer db.Close()
	startServer()
}

func startServer() {
	initHttpHandleFuncs()

	log.Println("Running server on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func initHttpHandleFuncs() {
	controller := controllers.CotacaoController{Db: db}
	http.HandleFunc("/cotacao", controller.GetCotacaoHandler)
}

func initDatabase() (db *sql.DB) {
	log.Println("Connecting to database...")

	db, err := sql.Open("sqlite3", "./cotacao.db")

	if err != nil {
		panic(err)
	}

	log.Println("Database connected successfully.")

	return
}
