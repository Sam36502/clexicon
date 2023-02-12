package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {

	e.PUT("/lang", PutLang)
	e.GET("/lang", GetLangs)

}
