package service

import (
	"context"

	"github.com/jqllxew/YuriCore/main_service/model/user"
	"github.com/jqllxew/YuriCore/user_service/client"
)

type UpdateBuymenuService interface {
	Handle(ctx context.Context) (*user.UserInfo, error)
}

type updateBuymenuServiceImpl struct {
	UserID    uint32
	BuymenuID uint16
	Slot      uint8
	ItemID    uint16
}

func NewUpdateBuymenuServiceImpl(UserID uint32, BuymenuID uint16, Slot uint8, ItemID uint16) UpdateBuymenuService {
	return &updateBuymenuServiceImpl{
		UserID:    UserID,
		BuymenuID: BuymenuID,
		Slot:      Slot,
		ItemID:    ItemID,
	}
}

func (u *updateBuymenuServiceImpl) Handle(ctx context.Context) (*user.UserInfo, error) {

	return client.GetUserTableClient().UpdateBuymenu(ctx, u.UserID, u.BuymenuID, u.Slot, u.ItemID)
}
