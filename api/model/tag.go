package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Tag struct {
	ID          uint64 `json:"id" objectbox:"id"`
	Tag         string `json:"tag" objectbox:"unique:index"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}
