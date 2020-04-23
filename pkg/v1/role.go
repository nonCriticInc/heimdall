package v1

import (
	"errors"
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo"
	"github.com/nonCriticInc/heimdall/config"
	"gopkg.in/mgo.v2/bson"
)

func CreateRoles(context echo.Context) error {
	roleDtoList, err := getCreateRoleDtoListFromContext(context)
	if err != nil {
		return GenerateErrorResponse(context, "Payload Convertion Error!", err)
	}
	validationErr:=roleDtoList.Validate()

	if validationErr!=nil{
		return GenerateErrorResponse(context, "Validation Error!", validationErr.Error())
	}
	var results []error

	for _,app:=range roleDtoList.Roles {
		temp :=app.GetRole()
		if temp.FindById().Id != "" {
			results = append(results, errors.New("Role by id "+temp.Id+" already exixts!"))
		} else {
			temp.Application=roleDtoList.Application
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
	return GenerateSuccessResponse(context,roleDtoList.Roles,"Roles Saved successfully!")
}

func getCreateRoleDtoListFromContext(context echo.Context) (*CreateRoleDtoList, error) {
	formData := new(RolePostRequestBody)
	if err := context.Bind(formData); err != nil {
		return nil, err
	}
	return &formData.Attributes, nil
}


type RolePostRequestBody struct {
	Id                 string     `json:"id"`
	Type               string     `json:"type"`
	Attributes         CreateRoleDtoList   `json:"attributes"`
}

type CreateRoleDtoList struct {
	Application  string `json:"application"`
	Roles []CreateRoleDto `json:"roles"`
}

func (dtoList *CreateRoleDtoList) Validate() error {
	app:=Application{
		Id: dtoList.Application,
	}
	if app.FindById().Id==""{
		return errors.New("No App by id "+dtoList.Application+" Exists!")
	}

	return nil
}

type CreateRoleDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
	Code   string `json:"code"`
	Application  string `json:"application"`
}

func (dto CreateRoleDto) GetRole() Role{

	permission:=Role{
		Id:           dto.Id,
		Name:         dto.Name,
		Parent:       dto.Parent,
		Permissions:  nil,
		Code:         dto.Code,
		Application:  dto.Application,
	}
	return permission
}


// entites
type Role struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Name               string       `bson:"name" json:"name"`
	Parent             string       `bson:"parent" json:"parent"`
	Permissions        []string `bson:"permissions" json:"permissions"`
	Code               string       `bson:"code" json:"code"`
	Application        string   `bson:"application" json:"application"`
}


func (role *Role) Save() error{
	err := config.RoleCollection.Save(role)
	if err != nil {
		return err
	}
	return nil
}

func (role *Role) FindAllChildRoles() [] Role{
	query := bson.M{"$and": []bson.M{
		{"parent": role.Id},
	},
	}
	temp:=[]Role{}
	config.RoleCollection.Find(query).Query.All(&temp)
	return temp
}


func (role *Role) FindAllParentRoles() [] Role{
	query := bson.M{"$and": []bson.M{
		{"parent": ""},
		{"application": role.Application},
	},
	}
	temp:=[]Role{}
	config.RoleCollection.Find(query).Query.All(&temp)
	return temp
}

func (role *Role) FindAllParentPermissions() [] Permission{
	query := bson.M{"$and": []bson.M{
		{"parent": ""},
		{"role": role.Id},
	},
	}
	temp:=[]Permission{}
	config.PermissionCollection.Find(query).Query.All(&temp)
	return temp
}

func (role *Role) FindAllPermissions() [] Permission{
	query := bson.M{"$and": []bson.M{
		{"role": role.Id},
	},
	}
	temp:=[] Permission{}
	config.PermissionCollection.Find(query).Query.All(&temp)
	return temp
}

func (role *Role) FindById() Role {
	query := bson.M{"$and": []bson.M{
		{"id": role.Id},
	},
	}
	tempRole:=Role{}
	config.RoleCollection.Find(query).Query.One(&tempRole)
	return tempRole
}
