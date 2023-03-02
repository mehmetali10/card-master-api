package main

import (
	_ "github.com/lib/pq"

	"src/app/config"
	cardCommands "src/app/dbActions/cardDbActions/commands"
	cardQueries "src/app/dbActions/cardDbActions/queries"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	cardQueries.SelectCardById(1, db)
	cardCommands.DeleteCardById(99, db)

}
