package main

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"src/app"
	"src/app/config"
	ca "src/app/dbActions"
	card "src/app/models/cards"
)

func main() {
	fmt.Println("Hello, Modules!")
	app.PrintHello()
	db := config.ConnectDB()
	defer db.Close()
	c1 := card.CardModel{Id: 1, Title: "ap-title", Description: "api-desc", ImgUri: "api-img", DateCreated: time.Now(), IsStarred: false}
	ca.InsertCard(card.CardModel(c1), db)

}
