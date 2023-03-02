package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1q2w3e"
	dbname   = "CardDB"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectDB() *sql.DB {
	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	//check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected to db!")
	return db
}
