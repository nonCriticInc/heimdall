package v1


type PostRequestBody struct {
	Id                 interface{}     `json:"id"`
	Type               string     `json:"type"`
	Attributes         interface{}     `json:"attributes"`
}