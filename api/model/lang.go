package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Lang struct {
	ID   uint64 `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
