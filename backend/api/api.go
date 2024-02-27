package api

import (
	"github.com/Acapetis/whatDoUMeme/model"
	"github.com/labstack/echo/v4"
)

type Api struct {
	// dbClient *client.Client
}

type IApi interface {
	GetAll(ctx echo.Context) *model.MemeData
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) GetAll(ctx echo.Context) *model.MemeData {
	return &model.MemeData{
		Name:        "test",
		Description: "test",
	}
}
