package v1

import "github.com/go-bongo/bongo"

type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string         `bson:"id"`
	Name               string         `bson:"name"`
	Adress             string         `bson:"address"`
	Phone              string         `bson:"phone"`
	Email              string         `bson:"email"`
	Organizations      []Organization `bson:"organizations"`
}

type Organization struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string        `bson:"id"`
	Name               string        `bson:"name"`
	Adress             string        `bson:"address"`
	Phone              string        `bson:"phone"`
	Email              string        `bson:"email"`
	Applications       []Application `bson:"applications"`
}

type Application struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string     `bson:"id"`
	Name               string     `bson:"name"`
	Type               string     `bson:"type"`
	Certs              []Cert     `bson:"certs"`
	Clients            []Client   `bson:"clients"`
	Resources          []Resource `bson:"resources"`
	Roles              []Role     `bson:"roles"`
}

type Cert struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	PublicKey          string `bson:"publicKey"`
	PrivateKey         string `bson:"privateKey"`
}

type Client struct {
	bongo.DocumentBase   `bson:",inline"`
	Id                   string                `bson:"id"`
	Name                 string                `bson:"name"`
	UserName             string                `bson:"userName"`
	Password             string                `bson:"password"`
	IsActive             bool                  `bson:"isActive"`
	Clients_Oauth_Tokens []Clients_Oauth_Token `bson:"clients_Oauth_Tokens"`
}

type Clients_Oauth_Token struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Token              string `bson:"token"`
	Iat                int64  `bson:"iat"`
}

type Resource struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Name               string `bson:"token"`
}

type Role struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string       `bson:"id"`
	Name               string       `bson:"token"`
	Parent             string       `bson:"parent"`
	Permissions        []Permission `bson:"permissions"`
}

type Permission struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Name               string `bson:"token"`
	Parent             string `bson:"parent"`
}

type User struct {
	bongo.DocumentBase          `bson:",inline"`
	Id                          string `bson:"id"`
	Name                        string `bson:"token"`
	IsActive                    string `bson:"isActive"`
	Username                    string `bson:"username"`
	Password                    string `bson:"password"`
	Email                       string `bson:"email"`
	Phone                       string `bson:"phone"`
	IsEmailVerificationEnabled  bool   `bson:"isEmailVerificationEnabled"`
	IsMobileVerificationEnabled bool   `bson:"isMobileVerificationEnabled"`
	Entity                      struct {
		Id           string `bson:"id"`
		Organization [] struct  {
			Id          string `bson:"id"`
			Application [] struct {
				Id   string `bson:"id"`
				Role [] struct {
					Id string `bson:"id"`
				}
			}
		}
	}
}
