package controller

import (
	"context"

	"github.com/jqllxew/YuriCore/main_service/model/server"
	"github.com/jqllxew/YuriCore/room_service/service"
)

type NewRoomController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type newRoomControllerImpl struct {
	room server.Room
}

func NewNewRoomController(room server.Room) NewRoomController {
	return &newRoomControllerImpl{
		room: room,
	}
}

func (r *newRoomControllerImpl) Handle(ctx context.Context) (*server.Room, error) {

	return service.NewNewRoomService(r.room).Handle(ctx)
}
