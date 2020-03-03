package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

type Users_Oauth_Token struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	UserId             string `bson:"userId"`
	Token              string `bson:"token"`
	TokenIat           int64  `bson:"tokenIat"`
	RefreshToken       string `bson:"refreshToken"`
	RefreshTokenIat    int64  `bson:"refreshTokenIat"`
	Code               string `bson:"code"`
}


func (users_Oauth_Token *Users_Oauth_Token) Save() error{
	err := config.UsersOauthTokenCollection.Save(users_Oauth_Token)
	if err != nil {
		return err
	}
	return nil
}

func (users_Oauth_Token *Users_Oauth_Token) FindByUserId() Users_Oauth_Token{
	query := bson.M{"$and": []bson.M{
		{"userId": users_Oauth_Token.UserId},
	},
	}
	temp:=Users_Oauth_Token{}
	config.UsersOauthTokenCollection.Find(query).Query.One(&temp)
	return temp
}
