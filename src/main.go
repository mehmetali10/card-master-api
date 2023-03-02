package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"src/app"
	"src/app/config"
	ca "src/app/dbActions"
)

func main() {
	fmt.Println("Hello, Modules!")
	app.PrintHello()
	db := config.ConnectDB()
	defer db.Close()
	//c1 := card.CardModel{Id: 1, Title: "ap-title", Description: "api-desc", ImgUri: "api-img", DateCreated: time.Now(), IsStarred: false}
	//ca.InsertCard(card.CardModel(c1), db)
	c := ca.SelectCardById(1, db)
	c.Title = "updated-title"
	ca.UpdateCardById(c, db)
	fmt.Println(c)
}
