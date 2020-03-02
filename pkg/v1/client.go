package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	bongo.DocumentBase   `bson:",inline"`
	Id                   string                `bson:"id"`
	Name                 string                `bson:"name"`
	UserName             string                `bson:"userName"`
	Password             string                `bson:"password"`
	IsActive             bool                  `bson:"isActive"`
	Clients_Oauth_Tokens []Clients_Oauth_Token `bson:"clients_Oauth_Tokens"`
	Code                 string                `bson:"code"`
}

func (client *Client) Save() error{
	err := config.ClientCollection.Save(client)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) FindById() Client{
	query := bson.M{"$and": []bson.M{
		{"id": client.Id},
	},
	}
	tempClient:=Client{}
	config.ClientsOauthTokenCollection.Find(query).Query.One((tempClient))
	return tempClient
}


func (client *Client) FindByUserName() Client{
	query := bson.M{"$and": []bson.M{
		{"userName": client.UserName},
	},
	}
	tempClient:=Client{}
	config.ClientsOauthTokenCollection.Find(query).Query.One((tempClient))
	return tempClient
}
