
package models

var (
	IssueCollection = "issue"
)


type Issue struct {

	Title string `json:"title" bson:"title"`
	Desc string `json:"desc" bson:"desc"`
     

}