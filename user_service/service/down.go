package service

import (
	"context"
	"errors"
	"github.com/jqllxew/YuriCore/verbose"

	"github.com/jqllxew/YuriCore/user_service/client"
)

type DownService interface {
	Handle(ctx context.Context) error
}

type downServiceImpl struct {
	userID uint32
}

func NewDownService(userID uint32) DownService {
	return &downServiceImpl{
		userID: userID,
	}
}

func (d *downServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserTableClient().GetUserByID(ctx, d.userID)
	if u == nil || u.UserID == 0 {
		return errors.New("down: wrong user id")
	}
	if err := client.GetDBClient().UpdateUser(ctx, u); err != nil {
		verbose.DebugPrintf(2, "down: update user failed! user=%+v\n", *u)
	}
	return client.GetUserTableClient().DeleteUserByID(ctx, d.userID)
}
