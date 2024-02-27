package main

import (
	"net/http"

	"github.com/Acapetis/whatDoUMeme/api"
	"github.com/Acapetis/whatDoUMeme/handler"
	"github.com/Acapetis/whatDoUMeme/service"
	"github.com/labstack/echo/v4"
)

func init() {
	// TODO
}

func main() {
	handler := handler.NewHandler(service.NewService(api.NewApi()))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/getall", handler.GetAll)
	e.Logger.Fatal(e.Start(":1323"))
}
