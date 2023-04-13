package controller

import (
	"context"
	"fmt"
	"github.com/jqllxew/YuriCore/verbose"
	"net"
	"strings"

	"github.com/jqllxew/YuriCore/main_service/constant"
	"github.com/jqllxew/YuriCore/main_service/model/packet"
	"github.com/jqllxew/YuriCore/main_service/service/chat"
	"github.com/jqllxew/YuriCore/utils"
)

type ChatController interface {
	Handle(ctx context.Context) error
}

type chatControllerImpl struct {
	client      net.Conn
	messageType uint8
	message     string
	seq         *uint8
}

func GetChatController(client net.Conn, p *packet.PacketData) ChatController {
	impl := chatControllerImpl{}

	impl.client = client
	impl.messageType = utils.ReadUint8(p.Data, &p.CurOffset)
	impl.message = strings.TrimSpace(string(p.Data[p.CurOffset:])) // 读取剩下的所有消息，当作用户输入
	p.CurOffset = len(p.Data)                                      // offset 拉满，数据已经读完

	return &impl
}

func (c *chatControllerImpl) Handle(ctx context.Context) error {
	// 检查
	if len(c.message) == 0 {
		return fmt.Errorf("recived chat packet %+v with null message", c.messageType)
	}

	verbose.DebugPrintf(2, "got chat message %+v", c.message)
	if strings.HasPrefix(c.message, "/register") {
		msgs := strings.Split(c.message, " ")
		var seq uint8
		register := &registerControllerImpl{}
		register.username = msgs[1]
		register.password = msgs[2]
		register.seq = &seq
		register.client = c.client
		return register.Handle(ctx)
	} else if strings.HasPrefix(c.message, "/login") {
		msgs := strings.Split(c.message, " ")
		var seq uint8
		login := &loginControllerImpl{}
		login.username = msgs[1]
		login.password = msgs[2]
		login.seq = &seq
		login.client = c.client
		return login.Handle(ctx)
	}

	switch c.messageType {
	case constant.MessageTypeChannel:
		return chat.GetChatInfra(c.message, c.client).ChannelHandler(ctx)
	case constant.MessageTypeRoom:
		return chat.GetChatInfra(c.message, c.client).RoomHandler(ctx)
	default:
		return fmt.Errorf("Unknown chat packet %+v", c.messageType)
	}
}
