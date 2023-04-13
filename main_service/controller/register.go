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
	"github.com/jqllxew/YuriCore/main_service/service/register"
	"github.com/jqllxew/YuriCore/utils"
)

type RegisterController interface {
	Handle(ctx context.Context) error
}

type registerControllerImpl struct {
	client   net.Conn
	username string
	password string
	seq      *uint8
}

func GetRegisterController(client net.Conn, p *packet.PacketData, seq *uint8) RegisterController {
	impl := registerControllerImpl{}

	impl.client = client
	impl.seq = seq
	impl.username = strings.TrimSpace(utils.ReadStringToNULL(p.Data, &p.CurOffset))
	impl.password = strings.TrimSpace(utils.ReadStringToNULL(p.Data, &p.CurOffset))
	p.CurOffset = len(p.Data) // offset 拉满，数据已经读完

	return &impl
}

func (r *registerControllerImpl) Handle(ctx context.Context) error {
	// 检查
	if len(r.password) == 0 || len(r.username) == 0 {
		out.OnSendMessage(packet.GetNextSeq(r.seq), r.client, constant.MessageDialogBox, constant.CSO_AuthReply_Wrong)
		return errors.New("null username or password")
	}

	verbose.DebugPrintf(2, "got register message %+v %+v", r.username, r.password)

	return register.GetRegisterServiceImpl(r.username, r.password, r.client, r.seq).Handle(ctx)
}
