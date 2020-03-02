package v1

import "github.com/go-bongo/bongo"

type Organization struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string        `bson:"id"`
	Name               string        `bson:"name"`
	Adress             string        `bson:"address"`
	Phone              string        `bson:"phone"`
	Email              string        `bson:"email"`
	Code               string        `bson:"code"`
	Applications       []Application `bson:"applications"`
}
