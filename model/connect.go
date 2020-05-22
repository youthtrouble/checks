package model

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host = "localhost"
	port = 8000
	user = "postgres"
	password = "deswerf"
	dbname = "forthebirds"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

//InitDB establishes the database connection
func InitDB() *sql.DB {
	var err error
	// Connect to the postgres db
	//you might have to change the connection string to add your database credentials
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port , user ,password, dbname)
	db, err = sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("Connected to database")
	return db
}