package controller

import (
	"context"

	"github.com/jqllxew/YuriCore/main_service/model/server"
	"github.com/jqllxew/YuriCore/room_service/service"
)

type ServerListController interface {
	Handle(ctx context.Context) ([]server.Server, error)
}

type serverListControllerImpl struct {
}

func NewServerListController() ServerListController {
	return &serverListControllerImpl{}
}

func (u *serverListControllerImpl) Handle(ctx context.Context) ([]server.Server, error) {

	return service.NewServerListService().Handle(ctx)
}
