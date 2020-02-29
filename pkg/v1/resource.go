package v1
import "github.com/go-bongo/bongo"

type Resource struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	Name               string `bson:"token"`
	Code string `bson:"code"`
}