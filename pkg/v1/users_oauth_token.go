package v1

import "github.com/go-bongo/bongo"

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
