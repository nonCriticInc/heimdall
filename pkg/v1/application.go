package v1

import (
	"errors"
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo"
	"github.com/nonCriticInc/heimdall/config"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2/bson"
)

//apis
func CreateApplications(context echo.Context) error {
	applicationDtoList, err := getCreateApplicationDtoListFromContext(context)
	if err != nil {
		return GenerateErrorResponse(context, "Payload Convertion Error!", err)
	}

	validationErr:=applicationDtoList.Validate()
	if validationErr!=nil{
		return GenerateErrorResponse(context, "Validation Error!", validationErr.Error())
	}
	var results []error

	for _,app:=range applicationDtoList.Applications {
		temp :=app.GetApplication()
		if temp.FindById().Id != "" {
			results = append(results, errors.New("Application by id "+temp.Id+" already exixts!"))
		} else {
			temp.Organization=applicationDtoList.Organization
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
	return GenerateSuccessResponse(context,applicationDtoList.Applications,"Applications Saved successfully!")

}
func FindApplicationById(context echo.Context) error {
	id:=context.Param("id")
	app:=Application{
		Id: id,
	}
	app=app.FindById()
	if(app.Id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	return GenerateSuccessResponse(context,app,"")
}


func FindResourcesByApplication(context echo.Context) error {
	id:=context.Param("id")
	app:=Application{
		Id: id,
	}
	if(app.Id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	return GenerateSuccessResponse(context,app.FindResourcesByApplicationId(),"")
}

func getCreateApplicationDtoListFromContext(context echo.Context) (*CreateApplicationDtoList, error) {
	formData := new(ApplicationPostRequestBody)
	if err := context.Bind(formData); err != nil {
		return nil, err
	}
	return &formData.Attributes, nil
}

type ApplicationPostRequestBody struct {
	Id                 string     `json:"id"`
	Type               string     `json:"type"`
	Attributes         CreateApplicationDtoList   `json:"attributes"`
}


type CreateApplicationDtoList struct {
	Organization string `json:"organization"`
	Applications []CreateApplicationDto `json:"applications"`
}


type CreateApplicationDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type string `json:"type"`
	Code   string `json:"code"`
	Email  string `json:"email"`
	Organization string `json:"organization"`
}

func (createApplicationDto *CreateApplicationDto) GetApplication() (Application) {
	app:= Application{
		Id:           createApplicationDto.Id,
		Name:         createApplicationDto.Name,
		Type:         createApplicationDto.Type,
		Code:         createApplicationDto.Code,
		Organization: createApplicationDto.Organization,
	}
	return app
}

func (dtoList *CreateApplicationDtoList) Validate() error {
	org:=Organization{
		Id: dtoList.Organization,
	}
	if org.FindById().Id==""{
		return errors.New("No Application by id "+dtoList.Organization+" Exists!")
	}

	for _, dto := range dtoList.Applications {
		if (dto.Id == "") {
			dto.Id = uuid.NewV4().String()
		}
		if (dto.Name == "") {
			return errors.New("No Name has been provided!")
		} else if (dto.Type == "") {
			return errors.New("No Type has been provided!")
		} else if (dto.Code == "") {
			return errors.New("No Code has been provided!")
		} else if (dto.Email == "") {
			return errors.New("No Email has been provided!")
		}
	}

	return nil
}



type Application struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Name               string `bson:"name" json:"name"`
	Type             string `bson:"type" json:"type"`
	Code               string `bson:"code" json:"code"`
	Organization             string `bson:"organization" json:"organization"`
}


func (application Application) Save() error{
	err := config.ApplicationCollection.Save(&application)
	if err != nil {
		return err
	}
	return nil
}

func (application *Application) FindAll() [] Application{
	query := bson.M{}
	tempApplications:= []Application{}
	config.ApplicationCollection.Find(query).Query.All(&tempApplications)
	return tempApplications
}

func (application *Application) FindById() Application{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One(&tempApp)
	return tempApp
}

func (application *Application) FindResourcesByApplicationId() []Resource{
	query := bson.M{"$and": []bson.M{
		{"application": application.Id},
	},
	}
	tempResource:=[] Resource{}
	config.ResourceCollection.Find(query).Query.All((&tempResource))
	return tempResource
}

func (application *Application) FindRolesById() []Role{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return nil
}


func (application *Application) FindClientsById() []Client{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return nil
}



func (application *Application) FindCertByApplicationId() []Cert{
	query := bson.M{"$and": []bson.M{
		{"id": application.Id},
	},
	}
	tempApp:=Application{}
	config.ApplicationCollection.Find(query).Query.One((tempApp))
	return nil
}
