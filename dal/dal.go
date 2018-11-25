package dal 

import (
	"fmt"
	"aircto/db"
	"aircto/models"	
	// "gopkg.in/mgo.v2/bson"

)

type Dal struct{
}

func(d *Dal) InsertUser(userData models.User) error{

	db, _ := db.ConnectMongo()
 
	 col := db.C("users")
	  
	 fmt.Println("------")

	 return col.Insert(userData)
}
func(d *Dal)LoginUser(userData models.User) (u models.User){

	return 
}
