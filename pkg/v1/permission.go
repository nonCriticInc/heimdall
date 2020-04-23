package v1

import (
	"errors"
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo"
	"github.com/nonCriticInc/heimdall/config"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2/bson"
)
func CreatePermissions(context echo.Context) error {
	permissionDtoList, err := getCreatePermissionDtoListFromContext(context)
	if err != nil {
		return GenerateErrorResponse(context, "Payload Convertion Error!", err)
	}
	validationErr:=permissionDtoList.Validate()

	if validationErr!=nil{
		return GenerateErrorResponse(context, "Validation Error!", validationErr.Error())
	}
	var results []error

	for _,app:=range permissionDtoList.Permissions {
		temp :=app.GetPermission()
		if temp.FindById().Id != "" {
			results = append(results, errors.New("Permission by id "+temp.Id+" already exixts!"))
		} else {
			temp.Role=permissionDtoList.Role
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
	return GenerateSuccessResponse(context,permissionDtoList.Permissions,"Permissions Saved successfully!")
}

func FindChildPermissions(context echo.Context) error {
	id:=context.Param("id")
	if(id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	permission :=Permission{
		Parent: id,
	}
	return GenerateSuccessResponse(context, permission.FindAllChildPermissions(),"")
}

func FindAllParentPermissionsByRole(context echo.Context) error {
	role:=context.Param("role")
	if(role==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	permission :=Permission{
		Role: role,
	}
	return GenerateSuccessResponse(context, permission.FindAllParentPermissions(),"")
}


func FindAllPermissionsByRole(context echo.Context) error {
	role:=context.Param("role")
	if(role==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	permission :=Permission{
		Role: role,
	}
	return GenerateSuccessResponse(context, permission.FindAllPermissionsByRole(),"")
}

func FindPermissionById(context echo.Context) error {
	id:=context.Param("id")
	permission :=Permission{
		Id: id,
	}
	permission = permission.FindById()
	if(permission.Id==""){
		return GenerateSuccessResponse(context,nil,"")
	}
	return GenerateSuccessResponse(context, permission,"")
}

func getCreatePermissionDtoListFromContext(context echo.Context) (*CreatePermissionDtoList, error) {
	formData := new(PermissionPostRequestBody)
	if err := context.Bind(formData); err != nil {
		return nil, err
	}
	return &formData.Attributes, nil
}

type PermissionPostRequestBody struct {
	Id                 string     `json:"id"`
	Type               string     `json:"type"`
	Attributes         CreatePermissionDtoList   `json:"attributes"`
}

type CreatePermissionDtoList struct {
	Role  string `json:"role"`
	Permissions []CreatePermissionDto `json:"permissions"`
}


func (dtoList *CreatePermissionDtoList) Validate() error {
	role:=Role{
		Id: dtoList.Role,
	}
	if role.FindById().Id==""{
		return errors.New("No Role by id "+dtoList.Role+" Exists!")
	}

	for _, dto := range dtoList.Permissions {
		if (dto.Id == "") {
			dto.Id = uuid.NewV4().String()
		}
		dto.Role=role.Id
		if (dto.Name == "") {
			return errors.New("No Name has been provided!")
		}  else if (dto.Code == "") {
			return errors.New("No Code has been provided!")
		}


	}

	return nil
}

type CreatePermissionDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
	Code   string `json:"code"`
	Role  string `json:"role"`
}

func (dto CreatePermissionDto) GetPermission() Permission{
	
	permission:=Permission{
		Id:           dto.Id,
		Name:         dto.Name,
		Parent:       dto.Parent,
		Code:         dto.Code,
		Role:         dto.Role,
	}
return permission
}

//entities
type Permission struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 string `bson:"id" json:"id"`
	Name               string `bson:"token" json:"token"`
	Parent             string `bson:"parent" json:"parent"`
	Code               string `bson:"code" json:"code"`
	Role string `bson:"role"  json:"role"`
}


func (permission *Permission) Save() error{
	err := config.PermissionCollection.Save(permission)
	if err != nil {
		return err
	}
	return nil
}

func (permission *Permission) FindAllChildPermissions() [] Permission{
	query := bson.M{"$and": []bson.M{
		{"parent": permission.Id},
	},
	}
	temp:=[]Permission{}
	config.PermissionCollection.Find(query).Query.All(&temp)
	return temp
}

func (permission *Permission) FindAllParentPermissions() [] Permission{
	query := bson.M{"$and": []bson.M{
		{"parent": ""},
		{"role": permission.Role},
	},
	}
	temp:=[]Permission{}
	config.PermissionCollection.Find(query).Query.All(&temp)
	return temp
}

func(permission *Permission) FindAllPermissionsByRole()  [] Permission{
	query := bson.M{"$and": []bson.M{
		{"role": permission.Role},
	},
	}
	temp:=[]Permission{}
	config.PermissionCollection.Find(query).Query.All(&temp)
	return temp
}

func (permission *Permission) FindById()  Permission{
	query := bson.M{"$and": []bson.M{
		{"id": permission.Id},
	},
	}
	tempPermission:=Permission{}
	config.PermissionCollection.Find(query).Query.One(&tempPermission)
	return tempPermission
}
