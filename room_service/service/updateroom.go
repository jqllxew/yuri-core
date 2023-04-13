package service

import (
	"context"
	"github.com/jqllxew/YuriCore/main_service/model/server"
	"github.com/jqllxew/YuriCore/room_service/client"
	"github.com/jqllxew/YuriCore/verbose"
)

type UpdateRoomService interface {
	UpdateRoom(ctx context.Context) (*server.Room, error)
	UpdateRoomSafe(ctx context.Context) (*server.Room, error)
}

type updateRoomServiceImpl struct {
	room *server.Room
}

func NewUpdateRoomService(room *server.Room) UpdateRoomService {
	return &updateRoomServiceImpl{
		room: room,
	}
}

func (u *updateRoomServiceImpl) UpdateRoomSafe(ctx context.Context) (*server.Room, error) {
	if err := client.GetRoomTableClient().UpdateRoomSafe(ctx, u.room); err != nil {
		return nil, err
	}
	verbose.DebugPrintf(2, "update room successfully")
	return client.GetRoomTableClient().GetRoomInfo(ctx, u.room.RoomId)
}

func (u *updateRoomServiceImpl) UpdateRoom(ctx context.Context) (*server.Room, error) {
	if err := client.GetRoomTableClient().UpdateRoom(ctx, u.room); err != nil {
		return nil, err
	}
	verbose.DebugPrintf(2, "update room successfully")
	return client.GetRoomTableClient().GetRoomInfo(ctx, u.room.RoomId)
}
