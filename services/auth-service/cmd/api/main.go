package main

import (
	"database/sql"
	"fangjjcs/auth/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const PORT = "80"

var dbConnCounts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting authentication service")

	// TODO connect DB
	conn := connectDB()
	if conn == nil {
		log.Panic("Can not connect to Postgres!")
	}

	// Set up app
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	// Define http server
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	// Start the http server
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Verify DB connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB() *sql.DB {
	dsn := os.Getenv("DSN")

	// Loop trying to connect until success
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres is not ready yet")
			dbConnCounts++
		} else {
			log.Println("Connect to Postgres!")
			return conn
		}

		if dbConnCounts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Try again in two seconds ...")
		time.Sleep(2 * time.Second)
		continue

	}
}
