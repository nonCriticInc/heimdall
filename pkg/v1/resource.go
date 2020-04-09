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
func CreateReources(context echo.Context) error {
	resourceDtoList, err := getCreateResourceDtoListFromContext(context)
	if err != nil {
		return GenerateErrorResponse(context, "Payload Convertion Error!", err)
	}

	validationErr:=resourceDtoList.Validate()
	if validationErr!=nil{
		return GenerateErrorResponse(context, "Validation Error!", validationErr.Error())
	}
	var results []error

	for _,resource:=range resourceDtoList.Resources {
		temp :=resource.GetResource()
		if temp.FindById().Id != "" {
			results = append(results, errors.New("Resource by id "+temp.Id+" already exixts!"))
		} else {
			temp.Application=resourceDtoList.Application
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
	return GenerateSuccessResponse(context,resourceDtoList.Resources,"Resources Saved successfully!")

}

func FindResourceById(context echo.Context) error {
	id:=context.Param("id")
	resource:=Resource{
		Id: id,
	}
	resource=resource.FindById()
	if(resource.Id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	return GenerateSuccessResponse(context,resource,"")
}


func getCreateResourceDtoListFromContext(context echo.Context) (*CreateResourceDtoList, error) {
	formData := new(ResourcePostRequestBody)
	if err := context.Bind(formData); err != nil {
		return nil, err
	}
	return &formData.Attributes, nil
}


type ResourcePostRequestBody struct {
	Id                 string     `json:"id"`
	Type               string     `json:"type"`
	Attributes         CreateResourceDtoList   `json:"attributes"`
}


type CreateResourceDtoList struct {
	Application string `json:"application"`
	Resources []CreateResourceDto `json:"resources"`
}
type CreateResourceDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Application string `json:"application"`
}

func (createResourceDto *CreateResourceDto) GetResource() (Resource) {
	app:= Resource{
		Id:          createResourceDto.Id,
		Name:        createResourceDto.Name,
		Code:        createResourceDto.Code,
		Application: createResourceDto.Application,
	}
	return app
}

func (dtoList *CreateResourceDtoList) Validate() error {
	app:=Application{
		Id: dtoList.Application,
	}
	if app.FindById().Id==""{
		return errors.New("No Entity by id "+dtoList.Application+" Exists!")
	}
	for _, dto := range dtoList.Resources {
		if (dto.Id == "") {
			dto.Id = uuid.NewV4().String()
		}
		if (dto.Name == "") {
			return errors.New("No Name has been provided!")
		} else if (dto.Code == "") {
			return errors.New("No Code has been provided!")
		}
	}
	return nil
}


type Resource struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Name               string `bson:"token" json:"token"`
	Code               string `bson:"code" json:"code"`
	Application    string `bson:"application" json:"application"`
}

func (resource *Resource) FindById() Resource{
	query := bson.M{"$and": []bson.M{
		{"id": resource.Id},
	},
	}
	tempResource:=Resource{}
	config.ResourceCollection.Find(query).Query.One(&tempResource)
	return tempResource
}

func (resource *Resource) Save() error{
	err := config.ResourceCollection.Save(resource)
	if err != nil {
		return err
	}
	return nil
}