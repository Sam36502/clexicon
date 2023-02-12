package main

import (
	"clexicon-api/model"
	"clexicon-api/routes"
	"clexicon-api/util"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := Initialisation()

	// Start API
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", util.GlobalConfig.ApiPort)))

	Termination()
}

func Initialisation() *echo.Echo {

	// Initialisation
	var err error
	err = util.InitConfig()
	if err != nil {
		fmt.Println("Failed to read config:", err)
		os.Exit(1)
	}

	err = model.InitModel()
	if err != nil {
		fmt.Println("Failed to initialise data-model:", err)
		os.Exit(1)
	}

	e := echo.New()
	routes.InitRoutes(e)

	return e
}

func Termination() {

	model.TermModel()

}
