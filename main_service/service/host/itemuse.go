package host

import (
	"context"
	"errors"
	"github.com/jqllxew/YuriCore/verbose"
	"net"

	"github.com/jqllxew/YuriCore/main_service/client"
	"github.com/jqllxew/YuriCore/main_service/constant"
	"github.com/jqllxew/YuriCore/main_service/model/out"
	"github.com/jqllxew/YuriCore/main_service/model/packet"
	"github.com/jqllxew/YuriCore/utils"
)

type ItemUsingService interface {
	Handle(ctx context.Context) error
}

type itemUsingServiceImpl struct {
	UserID uint32
	ItemID uint16
	Unk00  uint8
	Unk01  uint8
	client net.Conn
}

func NewItemUsingService(p *packet.PacketData, client net.Conn) ItemUsingService {
	return &itemUsingServiceImpl{
		UserID: utils.ReadUint32(p.Data, &p.CurOffset),
		ItemID: utils.ReadUint16(p.Data, &p.CurOffset),
		Unk00:  utils.ReadUint8(p.Data, &p.CurOffset),
		Unk01:  utils.ReadUint8(p.Data, &p.CurOffset),
		client: client,
	}
}

func (i *itemUsingServiceImpl) Handle(ctx context.Context) error {
	host := client.GetUserCacheClient().GetUserByConnection(ctx, i.client)
	if host == nil || host.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	dest_u := client.GetUserCacheClient().GetUserByID(ctx, i.UserID)
	if dest_u == nil || dest_u.CurrentRoomId == 0 {
		return errors.New("can't find dest_user or dest_user not in room")
	}

	// 发送结果
	rst := utils.BytesCombine(
		packet.BuildHeader(
			host.GetNextSeq(),
			constant.PacketTypeHost,
		),
		out.BuildHostItemUsing(dest_u.UserID, i.ItemID, 1),
	)
	packet.SendPacket(rst, host.CurrentConnection)
	verbose.DebugPrintf(2, "User %+v used item=%+v in match", dest_u.UserName, i.ItemID)
	return nil
}
