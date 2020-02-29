package v1

import "github.com/go-bongo/bongo"

type ETCD struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	ApplicationId      string `bson:"applicationId"`
	UserId             string `bson:"userId"`
	IsActive           bool   `bson:"isActive"`
}
