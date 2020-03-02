package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
)

type Resource struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Name               string `bson:"token"`
	Code               string `bson:"code"`
}



func (resource *Resource) Save() error{
	err := config.ResourceCollection.Save(resource)
	if err != nil {
		return err
	}
	return nil
}