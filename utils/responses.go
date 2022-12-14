package utils

type UserResponse struct {
	Id       string `bson:"_id" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
