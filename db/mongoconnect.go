package db
import (
	"os"
	mgo "gopkg.in/mgo.v2"
)

//GetSession :  connect to mongodb and gives sesson
func ConnectMongo() (*mgo.Database, *mgo.Session) {

	
	uri := os.Getenv("MONGODB_URL")
	// cv := config.Config()

	
	mgoSession, err := mgo.Dial(uri)

	
	if err != nil {
		panic(err)
	}
	databaseName := os.Getenv("DBNAME")

	db := mgoSession.DB(databaseName)

	return db, mgoSession
}