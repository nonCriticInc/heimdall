package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type Application struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string     `bson:"id"`
	Name               string     `bson:"name"`
	Type               string     `bson:"type"`
	Certs              []Cert     `bson:"certs"`
	Clients            []Client   `bson:"clients"`
	Resources          []Resource `bson:"resources"`
	Code               string     `bson:"code"`
	Roles              []Role     `bson:"roles"`
}


func (application *Application) Save() error{
	err := config.ApplicationCollection.Save(application)
	if err != nil {
		return err
	}
	return nil
}

func (application *Application) FindAll() [] Application{
	query := bson.M{}
	tempApplications:= []Application{}
	config.ApplicationCollection.Find(query).Query.All(&tempApplications)
	return tempApplications
}

func (application *Application) FindById() Application{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return tempApp
}

func (application *Application) FindResourcesById() []Resource{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return tempApp.Resources
}

func (application *Application) FindRolesById() []Role{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return tempApp.Roles
}


func (application *Application) FindClientsById() []Client{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return tempApp.Clients
}
