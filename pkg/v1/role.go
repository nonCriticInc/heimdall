package v1

import "github.com/go-bongo/bongo"

type Role struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string       `bson:"id"`
	Name               string       `bson:"token"`
	Parent             string       `bson:"parent"`
	Permissions        []Permission `bson:"permissions"`
	Code               string       `bson:"code"`
}
