package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

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
	Code                        string `bson:"code"`
	Entity                      struct {
		Id           string `bson:"id"`
		Organization [] struct {
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

func (user *User) Save() error{
	err := config.UserCollection.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) FindByUsername() User{
	query := bson.M{"$and": []bson.M{
		{"username": user.Username},
	},
	}
	temp:=User{}
	config.UserCollection.Find(query).Query.One(&temp)
	return temp
}
