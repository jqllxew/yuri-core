package service

import (
	"context"

	"github.com/jqllxew/YuriCore/main_service/model/server"
	"github.com/jqllxew/YuriCore/room_service/client"
)

type ServerListService interface {
	Handle(ctx context.Context) ([]server.Server, error)
}

type serverListServiceImpl struct {
}

func NewServerListService() ServerListService {
	return &serverListServiceImpl{}
}

func (s *serverListServiceImpl) Handle(ctx context.Context) ([]server.Server, error) {
	return client.GetRoomTableClient().GetServerList(ctx)
}
