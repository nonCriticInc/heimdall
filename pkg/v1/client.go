package v1
import "github.com/go-bongo/bongo"

type Client struct {
	bongo.DocumentBase   `bson:",inline"`
	Id                   string                `bson:"id"`
	Name                 string                `bson:"name"`
	UserName             string                `bson:"userName"`
	Password             string                `bson:"password"`
	IsActive             bool                  `bson:"isActive"`
	Clients_Oauth_Tokens []Clients_Oauth_Token `bson:"clients_Oauth_Tokens"`
	Code string `bson:"code"`
}
