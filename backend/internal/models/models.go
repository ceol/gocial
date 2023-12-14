package models

var Models []interface{}

func init() {
	Models = append(Models, &User{})
	Models = append(Models, &Post{})
}
