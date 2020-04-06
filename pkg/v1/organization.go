package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

//type OrganizationDto struct {
//	Id                 string        `json:"id"`
//	Name               string        `json:"name"`
//	Adress             string        `json:"address"`
//	Phone              string        `json:"phone"`
//	Email              string        `json:"email"`
//	Code               string        `json:"code"`
//	Applications       []ApplicationDto `json:"applications"`
//}

//func (organizationDto *OrganizationDto) GetOrganization() (Organization) {
//   organization:= Organization{
//	   Id:           organizationDto.Id,
//	   Name:         organizationDto.Name,
//	   Adress:       organizationDto.Adress,
//	   Phone:        organizationDto.Phone,
//	   Email:        organizationDto.Email,
//	   Code:         organizationDto.Code,
//	   Applications: nil,
//   }
// return organization
//}


type Organization struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string        `bson:"id"`
	Name               string        `bson:"name"`
	Adress             string        `bson:"address"`
	Phone              string        `bson:"phone"`
	Email              string        `bson:"email"`
	Code               string        `bson:"code"`
	Applications       []string `bson:"applications"`
	Entity string `bson:"entity"`
}


func (organization *Organization) Save() error{
	err := config.OrganizationCollection.Save(organization)
	if err != nil {
		return err
	}
	return nil
}

func (organization *Organization) FindAllApplications() [] Application{
	query := bson.M{"$and": []bson.M{
		{"organization": organization.Id},
	},
	}
	apps:=[]Application{}
	config.ApplicationCollection.Find(query).Query.All(&apps)
	return apps
}
