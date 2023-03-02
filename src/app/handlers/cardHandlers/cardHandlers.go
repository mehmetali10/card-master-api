package cardhandlers

import (
	"encoding/json"
	"net/http"
	database "src/app/config"

	//cardCommands "src/app/dbActions/cardDbActions/commands"
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
