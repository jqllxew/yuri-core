package controller

import (
	"github.com/jqllxew/YuriCore/verbose"
	"net"

	"github.com/jqllxew/YuriCore/main_service/model/packet"
)

type VersionController interface {
	Handle() error
}

type versionControllerImpl struct {
	info   []byte
	client net.Conn
}

func GetVersionController(p *packet.PacketData, client net.Conn) VersionController {
	version := versionControllerImpl{}
	version.info = p.Data[p.CurOffset:]
	version.client = client
	p.CurOffset = len(p.Data) // offset 拉满，数据已经读完

	return &version
}

func (v *versionControllerImpl) Handle() error {
	verbose.DebugPrintf(2, "Recived version info %+v from %+v", v.info, v.client.RemoteAddr().String())
	return nil
}
