package v1
import "github.com/go-bongo/bongo"

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
	Code string `bson:"code"`
	Entity       struct {
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
