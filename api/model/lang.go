package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Lang struct {
	ID            uint64   `json:"id" objectbox:"id"`
	Code          string   `json:"code" objectbox:"unique:index"`
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	AncestorCodes []string `json:"ancestors"`
}
