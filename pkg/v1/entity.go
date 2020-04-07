package v1

import (
	"errors"
	"github.com/go-bongo/bongo"
    "github.com/twinj/uuid"
	"github.com/labstack/echo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//apis
func CreateEntity(context echo.Context) error {
	createEntityDto, err := getCreateEntityDtoFromContext(context)
	if err != nil {
		return GenerateErrorResponse(context,"Payload Convertion Error!",err.Error())
	}
	if createEntityDto.Validate() != nil {
		return GenerateErrorResponse(context,"Validation Error!",err.Error())
	}
	entity:=createEntityDto.GetEntity()
	temp:=entity.FindById()
	if temp.Id==""{
		return GenerateErrorResponse(context,"Entity Already exists!",err.Error())
	}
	persistingErr := entity.Save()
	if (persistingErr != nil) {
		return GenerateErrorResponse(context,"Data persisting Error!",err.Error())
	}
	return GenerateSuccessResponse(context,entity,"Entity Saved successfully!")
}

func FindById(context echo.Context) error {
	id:=context.Param("id")
	entity:=Entity{
		Id: id,
	}
	entity=entity.FindById()
	if(entity.Id==""){
		return GenerateSuccessResponse(context,nil,"Entity Saved successfully!")
	}
	return GenerateSuccessResponse(context,entity,"Entity Saved successfully!")
}

//dtos
type CreateEntityDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Adress string `json:"address"`
	Phone  string `json:"phone"`
	Code   string `json:"code"`
	Email  string `json:"email"`
}

func (createEntityDto *CreateEntityDto) GetEntity() (Entity) {

	entity := Entity{
		Id:            createEntityDto.Id,
		Name:          createEntityDto.Name,
		Adress:        createEntityDto.Adress,
		Phone:         createEntityDto.Phone,
		Code:          createEntityDto.Code,
		Email:         createEntityDto.Email,
	}
	return entity
}
func (createEntityDto *CreateEntityDto) Validate() error {
	if (createEntityDto.Id == "") {
		createEntityDto.Id=uuid.NewV4().String()
	}
	if (createEntityDto.Name == "") {
		return errors.New("No Name has been provided!")
	} else if (createEntityDto.Adress == "") {
		return errors.New("No Adress has been provided!")
	} else if (createEntityDto.Phone == "") {
		return errors.New("No Phone has been provided!")
	} else if (createEntityDto.Code == "") {
		return errors.New("No Code has been provided!")
	} else if (createEntityDto.Email == "") {
		return errors.New("No Email has been provided!")
	}
	return nil
}

func getCreateEntityDtoFromContext(context echo.Context) (*CreateEntityDto, error) {
	formData := new(PostRequestBody)
	if err := context.Bind(formData); err != nil {
		return nil, err
	}
	log.Print(formData.Attributes)

	temp:=formData.Attributes
	createEntityDto := temp.(CreateEntityDto)
	createEntityDto.Id = formData.Id.(string)


	return &createEntityDto, nil
}
//entities
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string         `bson:"id"`
	Name               string         `bson:"name"`
	Adress             string         `bson:"address"`
	Phone              string         `bson:"phone"`
	Code               string         `bson:"code"`
	Email              string         `bson:"email"`
}



func (entity *Entity) Save() error {
	err := config.EntityCollection.Save(entity)
	if err != nil {
		return err
	}
	return nil
}


func (entity *Entity) FindById() Entity {
	query := bson.M{"$and": []bson.M{
		{"id": entity.Id},
	},
	}
	tempEntity := Entity{}
	config.EntityCollection.Find(query).Query.One(&tempEntity)
	return tempEntity
}

func (entity *Entity) FindAll() [] Entity {
	query := bson.M{}
	tempEntities := []Entity{}
	config.EntityCollection.Find(query).Query.All(&tempEntities)
	return tempEntities
}

func (entity *Entity) FindAllOrganizations() [] Organization {
	query := bson.M{"$and": []bson.M{
		{"entity": entity.Id},
	},
	}
	temp := []Organization{}
	config.OrganizationCollection.Find(query).Query.All(&temp)
	return temp
}
