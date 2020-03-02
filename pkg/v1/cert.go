package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
)

type Cert struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id"`
	PublicKey          string `bson:"publicKey"`
	PrivateKey         string `bson:"privateKey"`
	Code               string `bson:"code"`
}

func (cert *Cert) Save() error{
	err := config.CertCollection.Save(cert)
	if err != nil {
		return err
	}
	return nil
}
