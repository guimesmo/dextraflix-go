package models

type User struct {
	ID int `db:"_id" json:"id"`
	Name string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Password string `db:"password" json:"password"validate:"required"`
}