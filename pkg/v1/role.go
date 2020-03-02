package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type Role struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string       `bson:"id"`
	Name               string       `bson:"token"`
	Parent             string       `bson:"parent"`
	Permissions        []Permission `bson:"permissions"`
	Code               string       `bson:"code"`
}


func (role *Role) Save() error{
	err := config.RoleCollection.Save(role)
	if err != nil {
		return err
	}
	return nil
}

func (role *Role) FindAllPermissions() [] Permission{
	query := bson.M{"$and": []bson.M{
		{"id": role.Id},
	},
	}
	temp:=Role{}
	config.RoleCollection.Find(query).Query.One(&temp)
	return temp.Permissions
}
