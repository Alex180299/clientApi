package api

import (
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()

	e.POST("/NutriNET/Cliente", newClient)
	e.GET("/NutriNET/Cliente", getAllClients)
	e.GET("/NutriNET/Cliente/:id", getClient)
	e.PUT("/NutriNET/Cliente/:id", updateClient)

	e.Logger.Fatal(e.Start(":1323"))
}
