package models

import "github.com/kamva/mgm/v3"

// Todo Model (MongoDB Collection)
//type Todo struct {
	//mgm.DefaultModel `bson:",inline"`
	//Title            string `json:"title" bson:"title"`
	//Completed        bool   `json:"completed" bson:"completed"`
//}
type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Title string `json:"title" bson:"title"`
	Completed bool `json:"Completed" bson:"Completed"`
}