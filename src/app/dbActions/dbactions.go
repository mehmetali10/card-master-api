package dbactions

import (
	"database/sql"
	"fmt"

	card "src/app/models/cards"

	_ "github.com/lib/pq"
)

type cardModel card.CardModel

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func SelectCards(db *sql.DB) []cardModel {
	tuples, err := db.Query(`SELECT * FROM "Cards"`)
	CheckError(err)

	var cardCollections []cardModel

	defer tuples.Close()
	for tuples.Next() {
		var card cardModel

		err = tuples.Scan(&card.Id, &card.Title, &card.Description, &card.ImgUri, &card.DateCreated, &card.IsStarred)
		CheckError(err)
		cardCollections = append(cardCollections, card)
	}
	return cardCollections
}

func SelectCardById(id int, db *sql.DB) card.CardModel {
	query := `SELECT * FROM "Cards" WHERE "id"=$1`
	tuple, err := db.Query(query, id)
	CheckError(err)

	defer tuple.Close()
	var card card.CardModel
	for tuple.Next() {
		err = tuple.Scan(&card.Id, &card.Title, &card.Description, &card.ImgUri, &card.DateCreated, &card.IsStarred)
		CheckError(err)
	}
	return card
}

func InsertCard(card card.CardModel, db *sql.DB) bool {
	command := `insert into 
					"Cards"("title","description","imgUri","dateCreated")
					values($1,$2,$3,$4)`

	_, err := db.Exec(command, card.Title, card.Description, card.ImgUri, card.DateCreated)
	CheckError(err)
	return true
}

func DeleteCardById(id int, db *sql.DB) bool {
	deleteCommand := `DELETE FROM "Cards" where id=$1`

	_, err := db.Exec(deleteCommand, id)
	CheckError(err)
	return true
}

// UpdateCardById(id int, card Card, db *sql.DB)
func UpdateCardById(card card.CardModel, db *sql.DB) {
	updateCommand := `UPDATE "Cards" SET "title"=$1, "description"=$2, "imgUri"=$3, "isStarred"=$4 WHERE "id"=$5`

	_, err := db.Exec(updateCommand, card.Title, card.Description, card.ImgUri, card.IsStarred, card.Id)
	CheckError(err)
	fmt.Println("update success")
}

//StarCardById(id int, db *sql.DB)
//UnstarCardById(id int, db *sql.DB)
