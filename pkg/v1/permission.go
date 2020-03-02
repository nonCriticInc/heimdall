package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
)

type Permission struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Name               string `bson:"token"`
	Parent             string `bson:"parent"`
	Code               string `bson:"code"`
}


func (permission *Permission) Save() error{
	err := config.PermissionCollection.Save(permission)
	if err != nil {
		return err
	}
	return nil
}