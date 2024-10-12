package models

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Name     string `json:"name" bson:"name"`
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
	Pancard  string `json:"pancard" bson:"pancard"`
	Phone    string `json:"phone" bson:"phone"`
}
