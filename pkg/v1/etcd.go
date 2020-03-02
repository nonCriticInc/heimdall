package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type ETCD struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	ApplicationId      string `bson:"applicationId"`
	UserId             string `bson:"userId"`
	IsActive           bool   `bson:"isActive"`
}

func (etcd *ETCD) Save() error{
	err := config.EntityCollection.Save(etcd)
	if err != nil {
		return err
	}
	return nil
}

func (etcd *ETCD) FindByUserIdAndApplicationId() Application{
	query := bson.M{"$and": []bson.M{
		{"userId": etcd.UserId},
		{"applicationId": etcd.ApplicationId},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return tempApp
}

