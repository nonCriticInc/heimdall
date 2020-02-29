package v1
import "github.com/go-bongo/bongo"


type Cert struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	PublicKey          string `bson:"publicKey"`
	PrivateKey         string `bson:"privateKey"`
	Code string `bson:"code"`
}

