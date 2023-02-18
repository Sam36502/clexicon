package main

import (
	"clexicon-api/model"
	"clexicon-api/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := Initialisation()

	// Start API
	e.Logger.Fatal(e.Start(":34573"))

	Termination()
}

func Initialisation() *echo.Echo {
	var err error
	err = model.InitModel()
	if err != nil {
		fmt.Println("Failed to initialise data-model:", err)
		os.Exit(1)
	}

	err = model.InitSearch()
	if err != nil {
		fmt.Println("Failed to initialise word-index", err)
		os.Exit(1)
	}

	e := echo.New()
	routes.InitRoutes(e)

	return e
}

func Termination() {

	err := model.TermIndex()
	if err != nil {
		fmt.Println("Failed to close index:", err)
	}

	model.TermModel()

}
