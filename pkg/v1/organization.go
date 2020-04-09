package v1

import (
	"errors"
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo"
	"github.com/nonCriticInc/heimdall/config"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//apis
func CreateOrganizations(context echo.Context) error {
	organizationDtoList, err := getCreateOrganizationDtoListFromContext(context)
	if err != nil {
		return GenerateErrorResponse(context, "Payload Convertion Error!", err)
	}
	validationErr:=organizationDtoList.Validate()
	if validationErr!=nil{
			return GenerateErrorResponse(context, "Validation Error!", validationErr.Error())
	}
	var results []error
	for _,org:=range organizationDtoList.Organizations {
		temp := org.GetOrganization()
		if temp.FindById().Id != "" {
			results = append(results, errors.New("Organization by id "+temp.Id+" already exixts!"))
		} else {
			temp.Entity=organizationDtoList.Entity
		savingErr := temp.Save()
		if savingErr != nil {
			results = append(results, savingErr)
		}
	}
	}

	if len(results)>0{
		var errMessages [] string
		for _,result:=range results{
			errMessages=append(errMessages, result.Error())
		}
		return GenerateErrorResponse(context,"Data persisting Error!", errMessages)
	}
	return GenerateSuccessResponse(context,organizationDtoList.Organizations,"Organizations Saved successfully!")
}

func FindOrganizationById(context echo.Context) error {
	id:=context.Param("id")
	org:=Organization{
		Id: id,
	}
	org=org.FindById()
	if(org.Id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	return GenerateSuccessResponse(context,org,"")
}

func FindApplicationsByOrganization(context echo.Context) error {
	id:=context.Param("id")
	org:=Organization{
		Id: id,
	}
	if(org.Id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	return GenerateSuccessResponse(context,org.FindAllApplications(),"")
}


func getCreateOrganizationDtoListFromContext(context echo.Context) (*CreateOrganizationDtoList, error) {
	formData := new(OrganizationPostRequestBody)
	if err := context.Bind(formData); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &formData.Attributes, nil
}

type CreateOrganizationDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Adress string `json:"address"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Code   string `json:"code"`
	Entity string `bson:"entity"`
}

type OrganizationPostRequestBody struct {
	Id                 string     `json:"id"`
	Type               string     `json:"type"`
	Attributes         CreateOrganizationDtoList   `json:"attributes"`
}

type CreateOrganizationDtoList struct {
	Entity string `json:"entity"`
	Organizations []CreateOrganizationDto `json:"organizations"`
}


func (createOrganizationDto *CreateOrganizationDto) GetOrganization() (Organization) {
  organization:= Organization{
	  Id:           createOrganizationDto.Id,
	  Name:         createOrganizationDto.Name,
	  Adress:       createOrganizationDto.Adress,
	  Phone:        createOrganizationDto.Phone,
	  Email:        createOrganizationDto.Email,
	  Code:         createOrganizationDto.Code,
	  Entity:       createOrganizationDto.Entity,
  }
return organization
}

func (dtoList *CreateOrganizationDtoList) Validate() error {
	entity:=Entity{
		Id: dtoList.Entity,
	}
	if entity.FindById().Id==""{
		return errors.New("No Entity by id "+dtoList.Entity+" Exists!")
	}
	
	for _, dto := range dtoList.Organizations {
		if (dto.Id == "") {
			dto.Id = uuid.NewV4().String()
		}
		if (dto.Name == "") {
			return errors.New("No Name has been provided!")
		} else if (dto.Adress == "") {
			return errors.New("No Adress has been provided!")
		} else if (dto.Phone == "") {
			return errors.New("No Phone has been provided!")
		} else if (dto.Code == "") {
			return errors.New("No Code has been provided!")
		} else if (dto.Email == "") {
			return errors.New("No Email has been provided!")
		}
	}

	return nil
}

type Organization struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Name               string `bson:"name" json:"name"`
	Adress             string `bson:"address" json:"address"`
	Phone              string `bson:"phone" json:"phone"`
	Email              string `bson:"email" json:"email"`
	Code               string `bson:"code" json:"code"`
	Entity             string `bson:"entity" json:"entity"`
}

func (organization Organization) Save() error {
	err := config.OrganizationCollection.Save(&organization)
	if err != nil {
		return err
	}
	return nil
}

func (organization *Organization) FindById() Organization {
	query := bson.M{"$and": []bson.M{
		{"id": organization.Id},
	},
	}
	tempOrganization := Organization{}
	config.OrganizationCollection.Find(query).Query.One(&tempOrganization)
	return tempOrganization
}

func (organization *Organization) FindAllApplications() [] Application {
	query := bson.M{"$and": []bson.M{
		{"organization": organization.Id},
	},
	}
	apps := []Application{}
	config.ApplicationCollection.Find(query).Query.All(&apps)
	return apps
}
