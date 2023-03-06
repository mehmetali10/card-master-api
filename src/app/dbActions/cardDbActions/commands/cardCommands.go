package commands

import (
	"database/sql"
	"fmt"

	card "src/app/models/cards"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertCard(card card.CardModel, db *sql.DB) bool {
	command := `insert into 
					"tb_casestudy"("title","description","imguri","datecreated","isstarred")
					values($1,$2,$3,$4,$5)`

	_, err := db.Exec(command, card.Title, card.Description, card.ImgUri, card.DateCreated, false)
	checkError(err)
	return true
}

func DeleteCardById(id int, db *sql.DB) bool {
	deleteCommand := `DELETE FROM "tb_casestudy" where id=$1`

	_, err := db.Exec(deleteCommand, id)
	checkError(err)
	return true
}

func UpdateCardById(card card.CardModel, db *sql.DB) bool {
	updateCommand := `UPDATE "tb_casestudy" SET "title"=$1, "description"=$2, "imguri"=$3, "isstarred"=$4 WHERE "id"=$5`
	fmt.Print(card)
	_, err := db.Exec(updateCommand, card.Title, card.Description, card.ImgUri, card.IsStarred, card.Id)
	checkError(err)

	return true
}
