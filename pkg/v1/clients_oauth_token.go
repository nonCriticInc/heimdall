package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
)

type Clients_Oauth_Token struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Token              string `bson:"token"`
	Iat                int64  `bson:"iat"`
	Code               string `bson:"code"`
}


func (clients_Oauth_Token *Clients_Oauth_Token) Save() error{
	err := config.ClientsOauthTokenCollection.Save(clients_Oauth_Token)
	if err != nil {
		return err
	}
	return nil
}
