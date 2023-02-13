package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Word struct {
	ID            uint64   `json:"id" objectbox:"id"`
	Orthography   string   `json:"orthography"`
	Romanisation  string   `json:"romanisation" objectbox:"index"`
	Pronunciation string   `json:"ipa"`
	Meanings      []string `json:"meanings"`
	Tags          []*Tag   `json:"tags"`
	Etymology     string   `json:"etymology"`
	RootWord      uint64   `json:"root"`
	Notes         string   `json:"notes"`
	Language      *Lang    `json:"language" objectbox:"link"`
}
