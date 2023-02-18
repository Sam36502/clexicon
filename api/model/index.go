package model

import (
	"fmt"
	"strconv"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/v2/analysis/lang/en"
	"github.com/blevesearch/bleve/v2/mapping"
)

const INDEX_FILE = "wordindex.bleve"

var G_WordIndex bleve.Index

func InitSearch() error {

	var err error
	G_WordIndex, err = bleve.Open(INDEX_FILE)
	if err != nil {
		fmt.Println("No Index found; Building...")
		fmt.Println("This may take a while...")

		mapping := buildIndexMapping()
		G_WordIndex, err = bleve.New(INDEX_FILE, mapping)
		if err != nil {
			return err
		}

		allWords, err := G_WordBox.GetAll()
		if err != nil {
			fmt.Println("Failed to read words from data")
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

func buildIndexMapping() mapping.IndexMapping {

	// a generic reusable mapping for english text
	englishTextFieldMapping := bleve.NewTextFieldMapping()
	englishTextFieldMapping.Analyzer = en.AnalyzerName

	// a generic reusable mapping for keyword text
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword.Name

	// Tag Information Mapping
	tagMapping := bleve.NewDocumentMapping()
	tagMapping.AddFieldMappingsAt("tag", keywordFieldMapping)
	tagMapping.AddFieldMappingsAt("name", keywordFieldMapping)
	tagMapping.AddFieldMappingsAt("desc", englishTextFieldMapping)

	// Lang Information Mapping
	langMapping := bleve.NewDocumentMapping()
	langMapping.AddFieldMappingsAt("id", bleve.NewNumericFieldMapping())
	langMapping.AddFieldMappingsAt("code", keywordFieldMapping)
	langMapping.AddFieldMappingsAt("name", keywordFieldMapping)
	langMapping.AddFieldMappingsAt("desc", englishTextFieldMapping)
	langMapping.AddFieldMappingsAt("ancestors", keywordFieldMapping)

	// Word Information Mapping
	wordMapping := bleve.NewDocumentMapping()
	wordMapping.AddFieldMappingsAt("orthography", keywordFieldMapping)
	wordMapping.AddFieldMappingsAt("romanisation", keywordFieldMapping)
	wordMapping.AddFieldMappingsAt("ipa", keywordFieldMapping)
	wordMapping.AddFieldMappingsAt("meanings", englishTextFieldMapping)
	wordMapping.AddFieldMappingsAt("etymology", englishTextFieldMapping)
	wordMapping.AddFieldMappingsAt("notes", englishTextFieldMapping)
	wordMapping.AddSubDocumentMapping("tags", tagMapping)
	wordMapping.AddSubDocumentMapping("language", langMapping)

	// Create index map
	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("words", wordMapping)
	return indexMapping
}
