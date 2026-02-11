package models

type Card struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id"`
	Topic    string `json:"topic"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type CreateCardRequest struct {
	Topic    string `json:"topic"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	UserID   uint   `json:"UID"`
}

type UpdateCardRequest struct {
	Topic    *string `json:"topic"`
	Question *string `json:"question"`
	Answer   *string `json:"answer"`
	CardID   uint    `json:"CardID"`
}
