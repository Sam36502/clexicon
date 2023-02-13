package model

import (
	"fmt"
	"strconv"

	"github.com/blevesearch/bleve/v2"
)

const INDEX_FILE = "wordindex.bleve"

var G_WordIndex bleve.Index

func InitSearch() error {

	var err error
	G_WordIndex, err = bleve.Open(INDEX_FILE)
	if err != nil {
		fmt.Println("No Index found; Rebuilding...")
		fmt.Println("This may take a while...")

		mapping := bleve.NewIndexMapping()
		G_WordIndex, err = bleve.New(INDEX_FILE, mapping)
		if err != nil {
			return err
		}

		allWords, err := G_WordBox.GetAll()
		if err != nil {
			fmt.Println("Failed to read words from data-")
			return err
		}

		for _, word := range allWords {
			err := G_WordIndex.Index(strconv.FormatUint(word.ID, 10), word)
			if err != nil {
				fmt.Printf("Failed to index word: %v\n%v\n", word, err)
			}
		}

	}

	if count, err := G_WordIndex.DocCount(); err == nil {
		fmt.Printf("Loaded %d index entries.", count)
	}

	return nil
}

func TermIndex() error {
	return G_WordIndex.Close()
}
