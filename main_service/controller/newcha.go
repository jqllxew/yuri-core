package controller

import (
	"context"
	"errors"
	"github.com/jqllxew/YuriCore/verbose"
	"net"
	"strings"

	"github.com/jqllxew/YuriCore/main_service/constant"
	"github.com/jqllxew/YuriCore/main_service/model/out"
	"github.com/jqllxew/YuriCore/main_service/model/packet"
	"github.com/jqllxew/YuriCore/main_service/service/character"
	"github.com/jqllxew/YuriCore/utils"
)

type NewCharacterController interface {
	Handle(ctx context.Context) error
}

type newCharacterControllerImpl struct {
	client   net.Conn
	nickname string
	seq      *uint8
}

func GetNewCharacter(client net.Conn, p *packet.PacketData, seq *uint8) NewCharacterController {
	impl := newCharacterControllerImpl{}

	impl.client = client
	impl.nickname = strings.TrimSpace(utils.ReadStringToNULL(p.Data, &p.CurOffset))
	impl.seq = seq
	p.CurOffset = len(p.Data) // offset 拉满，数据已经读完

	return &impl
}

func (n *newCharacterControllerImpl) Handle(ctx context.Context) error {
	// 检查
	if len(n.nickname) == 0 {
		out.OnSendMessage(packet.GetNextSeq(n.seq), n.client, constant.MessageDialogBox, constant.CSO_NewCharacter_Wrong)
		return errors.New("null nickname")
	}

	verbose.DebugPrintf(2, "got nickname message %+v %+v", n.nickname)

	return character.GetNewServiceImpl(n.nickname, n.client).Handle(ctx)
}
