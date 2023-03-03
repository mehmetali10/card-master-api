package cardhandlers

import (
	"encoding/json"
	"net/http"
	database "src/app/config"
	"strconv"

	cardCommands "src/app/dbActions/cardDbActions/commands"
	cardQueries "src/app/dbActions/cardDbActions/queries"
	card "src/app/models/cards"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()
	cardCollections := cardQueries.SelectCards(db)

	var response = card.JsonResponse{Type: "success", Data: cardCollections, Message: "Success getting"}
	json.NewEncoder(w).Encode(response)
}

func GetCardById(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	var stringId = r.Header.Get("id")
	id, err := strconv.Atoi(stringId)
	database.CheckError(err)
	returnedCard := cardQueries.SelectCardById(id, db)
	var response = card.JsonResponse{Type: "success", Data: []card.CardModel{returnedCard}, Message: "success getting"}
	json.NewEncoder(w).Encode(response)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	var newCard card.CardModel
	err := json.NewDecoder(r.Body).Decode(&newCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := cardCommands.InsertCard(newCard, db)
	var response card.JsonResponse

	if result {
		response = card.JsonResponse{Type: "success insert", Data: nil, Message: "201"}
	} else {
		response = card.JsonResponse{Type: "failed insert", Data: nil, Message: "401"}
	}

	json.NewEncoder(w).Encode(response)
}
