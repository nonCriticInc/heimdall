package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string         `bson:"id"`
	Name               string         `bson:"name"`
	Adress             string         `bson:"address"`
	Phone              string         `bson:"phone"`
	Code               string         `bson:"code"`
	Email              string         `bson:"email"`
	Organizations      []Organization `bson:"organizations"`
}

func (entity *Entity) Save() error{
	err := config.EntityCollection.Save(entity)
	if err != nil {
	return err
	}
	return nil
}

func (entity *Entity) FindAll() [] Entity{
	query := bson.M{}
	tempEntities := []Entity{}
	config.EntityCollection.Find(query).Query.All(&tempEntities)
	return tempEntities
}


func (entity *Entity) FindAllOrganizations() [] Organization{
	query := bson.M{"$and": []bson.M{
		{"id": entity.Id},
	},
	}
	temp:=Entity{}
	config.EntityCollection.Find(query).Query.One(&temp)
	return temp.Organizations
}



