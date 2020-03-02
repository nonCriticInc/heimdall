package v1

import "github.com/go-bongo/bongo"

type Permission struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Name               string `bson:"token"`
	Parent             string `bson:"parent"`
	Code               string `bson:"code"`
}
