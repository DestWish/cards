package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Cards    []Card `json:"cards,omitempty" gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserPassword struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
}
