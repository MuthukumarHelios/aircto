package models

var (
	UserCollection ="users"
)

type User struct{

 	 Name string `json:"name" bson:"name"`
	 Email string `json:"email" bson:"email"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	Token string `json:"token" bson:"token"`
}
