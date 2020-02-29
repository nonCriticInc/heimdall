package v1

import "github.com/go-bongo/bongo"

type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string         `bson:"id"`
	Name               string         `bson:"name"`
	Adress             string         `bson:"address"`
	Phone              string         `bson:"phone"`
	Code string `bson:"code"`
	Email              string         `bson:"email"`
	Organizations      []Organization `bson:"organizations"`
}
