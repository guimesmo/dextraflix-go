package models

type Profile struct {
	ID int `db:"_id" json:"id"`
	Name string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Description string `db:description" json:"description"`
}