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
	"github.com/jqllxew/YuriCore/main_service/service/login"
	"github.com/jqllxew/YuriCore/utils"
)

type LoginController interface {
	Handle(ctx context.Context) error
}

type loginControllerImpl struct {
	client   net.Conn
	username string
	password string
	seq      *uint8
}

func GetLoginController(client net.Conn, p *packet.PacketData, seq *uint8) LoginController {
	impl := loginControllerImpl{}

	impl.client = client
	impl.seq = seq
	impl.username = strings.TrimSpace(utils.ReadStringToNULL(p.Data, &p.CurOffset))
	impl.password = strings.TrimSpace(utils.ReadStringToNULL(p.Data, &p.CurOffset))
	p.CurOffset = len(p.Data) // offset 拉满，数据已经读完

	return &impl
}

func (l *loginControllerImpl) Handle(ctx context.Context) error {
	// 检查
	if len(l.password) == 0 || len(l.username) == 0 {
		out.OnSendMessage(packet.GetNextSeq(l.seq), l.client, constant.MessageDialogBox, constant.CSO_AuthReply_Wrong)
		return errors.New("null username or password")
	}

	verbose.DebugPrintf(2, "got login message %+v %+v", l.username, l.password)

	return login.GetLoginServiceImpl(l.username, l.password, l.client, l.seq).Handle(ctx)
}
