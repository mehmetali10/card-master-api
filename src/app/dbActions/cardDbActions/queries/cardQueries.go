package queries

import (
	"database/sql"

	card "src/app/models/cards"

	_ "github.com/lib/pq"
)

type cardModel card.CardModel

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func SelectCards(db *sql.DB) []card.CardModel {
	tuples, err := db.Query(`SELECT * FROM "tb_casestudy"`)
	checkError(err)

	var cardCollections []card.CardModel

	defer tuples.Close()
	for tuples.Next() {
		var card card.CardModel

		err = tuples.Scan(&card.Id, &card.Title, &card.Description, &card.ImgUri, &card.DateCreated, &card.IsStarred)
		checkError(err)
		cardCollections = append(cardCollections, card)
	}
	return cardCollections
}

func SelectCardById(id int, db *sql.DB) card.CardModel {
	query := `SELECT * FROM "tb_casestudy" WHERE "id"=$1`
	tuple, err := db.Query(query, id)
	checkError(err)

	defer tuple.Close()
	var card card.CardModel
	for tuple.Next() {
		err = tuple.Scan(&card.Id, &card.Title, &card.Description, &card.ImgUri, &card.DateCreated, &card.IsStarred)
		checkError(err)
	}
	return card
}
