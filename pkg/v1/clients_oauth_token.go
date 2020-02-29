package v1
import "github.com/go-bongo/bongo"

type Clients_Oauth_Token struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Token              string `bson:"token"`
	Iat                int64  `bson:"iat"`
	Code string `bson:"code"`
}