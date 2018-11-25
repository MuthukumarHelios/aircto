package dal 

import (
	"fmt"
	"aircto/db"
	"aircto/models"	
)

type Dal struct{
}

func(d *Dal) InsertUser(userData models.User) error{

	db, _ := db.ConnectMongo()
 
	 col := db.C("users")
	  
	 fmt.Println("------")
	 return col.Insert(userData)


	// return nil
}
