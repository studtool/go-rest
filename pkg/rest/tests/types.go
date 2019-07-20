package tests

//go:generate easyjson

//easyjson:json
type TestingJSONType struct {
	SomeField string `json:"some_field"`
}
