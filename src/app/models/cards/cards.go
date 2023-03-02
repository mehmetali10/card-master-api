package cards

import (
	"time"
)

type CardModel struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImgUri      string    `json:"imgUri"`
	DateCreated time.Time `json:"dateCreated"`
	IsStarred   *bool     `json:"isStarred"`
}

type JsonResponse struct {
	Type    string      `json:"type"`
	Data    []CardModel `json:"data"`
	Message string      `json:"message"`
}
