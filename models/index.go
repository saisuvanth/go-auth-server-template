package models

type User struct {
	Id       string `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
