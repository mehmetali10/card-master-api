package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5439
	user     = "postgres"
	password = "depixen-pass"
	dbname   = "postgres"
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

	//create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tb_casestudy(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		imguri TEXT NOT NULL,
		datecreated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		isstarred BOOLEAN NOT NULL DEFAULT false
	)`)
	CheckError(err)

	fmt.Println("Connected to db!")
	return db
}
