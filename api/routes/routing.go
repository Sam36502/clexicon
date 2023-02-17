package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {

	// Static files
	e.Static("/", "static")

	// Lang Endpoints
	e.PUT("/lang", PutLang)
	e.GET("/lang", GetLang)

	// Tag Endpoints
	e.PUT("/tag", PutTag)
	e.GET("/tag", GetTag)

	// Word Endpoints
	e.PUT("/lang/:lid/word", PutWord)
	e.GET("/lang/:lid/word", GetWord)
	e.GET("/lang/:lid/search", SearchWord)

}
