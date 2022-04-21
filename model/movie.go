package model

type Movie struct {
	Id       string   `json:"id" bson:"_id,omitempty"`
	ISBN     string   `json:"isbn" bson:"isbn"`
	Title    string   `json:"title" bson:"title"`
	Director Director `json:"director" bson:"director"`
}

type Director struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}
