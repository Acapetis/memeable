package service

import (
	"github.com/Acapetis/whatDoUMeme/api"
	"github.com/Acapetis/whatDoUMeme/model"
	"github.com/labstack/echo/v4"
)

type Service struct {
	api api.IApi
}

type IService interface {
	GetAll(ctx echo.Context) *model.MemeData
}

func NewService(a api.IApi) *Service {
	return &Service{
		api: a,
	}
}

func (s *Service) GetAll(ctx echo.Context) *model.MemeData {
	result := s.api.GetAll(ctx)
	return result
}
