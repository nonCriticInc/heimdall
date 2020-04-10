package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type Role struct {
	bongo.DocumentBase `bson:",inline"`
	Id                  string `bson:"id" json:"id"`
	Name               string       `bson:"name" json:"name"`
	Parent             string       `bson:"parent" json:"parent"`
	Permissions        []string `bson:"permissions" json:"permissions"`
	Code               string       `bson:"code" json:"code"`
	Application        string   `bson:"application" json:"application"`
}


func (role *Role) Save() error{
	err := config.RoleCollection.Save(role)
	if err != nil {
		return err
	}
	return nil
}

func (role *Role) FindAllChildRoles() [] Role{
	query := bson.M{"$and": []bson.M{
		{"parent": role.Id},
	},
	}
	temp:=[]Role{}
	config.RoleCollection.Find(query).Query.All(&temp)
	return temp
}


func (role *Role) FindAllParentRoles() [] Role{
	query := bson.M{"$and": []bson.M{
		{"parent": ""},
	},
	}
	temp:=[]Role{}
	config.RoleCollection.Find(query).Query.All(&temp)
	return temp
}

func (role *Role) FindAllPermissions() [] Permission{
	query := bson.M{"$and": []bson.M{
		{"id": role.Id},
	},
	}
	temp:=Role{}
	config.PermissionCollection.Find(query).Query.One(&temp)
	return nil
}
