package service

import (
	"context"

	"github.com/jqllxew/YuriCore/main_service/model/user"
	"github.com/jqllxew/YuriCore/user_service/client"
)

type FriendsService interface {
	Handle(ctx context.Context) ([]user.UserInfo, error)
}

type friendsServiceImpl struct {
	userID uint32
}

func NewFriendsService(userID uint32) FriendsService {
	return &friendsServiceImpl{
		userID: userID,
	}
}

func (f *friendsServiceImpl) Handle(ctx context.Context) ([]user.UserInfo, error) {

	return client.GetUserTableClient().GetUserFriends(ctx, f.userID)
}
