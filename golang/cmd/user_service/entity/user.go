package entity

type User struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	FullName string `json:"full_name" bson:"full_name"`
	Password string `json:"password" bson:"password"`
}
