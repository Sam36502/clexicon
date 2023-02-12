package model

import "github.com/objectbox/objectbox-go/objectbox"

var g_ObjBox *objectbox.ObjectBox
var G_LangBox *LangBox

func InitModel() error {
	g_ObjBox, err := objectbox.NewBuilder().Model(ObjectBoxModel()).Build()
	if err != nil {
		return err
	}

	G_LangBox = BoxForLang(g_ObjBox)
	return nil
}

func TermModel() {
	g_ObjBox.Close()
}
