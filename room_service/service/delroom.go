package service

import (
	"context"

	"github.com/jqllxew/YuriCore/room_service/client"
)

type DelRoomService interface {
	Handle(ctx context.Context) error
}

type delRoomServiceImpl struct {
	roomID uint16
}

func NewDelRoomService(roomID uint16) DelRoomService {
	return &delRoomServiceImpl{
		roomID: roomID,
	}
}

func (g *delRoomServiceImpl) Handle(ctx context.Context) error {

	return client.GetRoomTableClient().DeleteRoom(ctx, g.roomID)
}
