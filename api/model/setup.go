package model

import "github.com/objectbox/objectbox-go/objectbox"

// Object Boxes
var g_ObjBox *objectbox.ObjectBox
var G_LangBox *LangBox
var G_TagBox *TagBox
var G_WordBox *WordBox

// Queries
var G_LangCodeQuery *LangQuery
var G_TagQuery *TagQuery
var G_WordSearchQuery *WordQuery

func InitModel() error {
	g_ObjBox, err := objectbox.NewBuilder().Model(ObjectBoxModel()).Build()
	if err != nil {
		return err
	}

	// Set up lang box and queries
	G_LangBox = BoxForLang(g_ObjBox)
	G_LangCodeQuery = G_LangBox.Query(Lang_.Code.Equals("", true /* Case-sensitive */))

	// Set up tag box and queries
	G_TagBox = BoxForTag(g_ObjBox)
	G_TagQuery = G_TagBox.Query(Tag_.Tag.Equals("", false /* Case-insensitive */))

	// Set up word box and queries
	G_WordBox = BoxForWord(g_ObjBox)
	G_WordSearchQuery = G_WordBox.Query(
		Word_.Orthography.Contains("", false),
		Word_.Romanisation.Contains("", false),
		Word_.Pronunciation.Contains("", false),
	)

	return nil
}

func TermModel() {
	g_ObjBox.Close()
}
