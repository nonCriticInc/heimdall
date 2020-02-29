package v1


import "github.com/go-bongo/bongo"

type Application struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string     `bson:"id"`
	Name               string     `bson:"name"`
	Type               string     `bson:"type"`
	Certs              []Cert     `bson:"certs"`
	Clients            []Client   `bson:"clients"`
	Resources          []Resource `bson:"resources"`
	Code string `bson:"code"`
	Roles              []Role     `bson:"roles"`
}
