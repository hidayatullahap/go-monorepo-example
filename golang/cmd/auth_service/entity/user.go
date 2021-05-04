package entity

type User struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	FullName string `json:"full_name" bson:"full_name"`
	Password string `json:"password" bson:"password"`
}

type Token struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID   string `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`
	Token    string `json:"token" bson:"token"`
}
