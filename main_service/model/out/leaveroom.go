package out

import (
	"github.com/jqllxew/YuriCore/main_service/constant"
	"github.com/jqllxew/YuriCore/utils"
)

func BuildLeaveRoom(id uint32) []byte {
	buf := make([]byte, 8)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTPlayerLeave, &offset)
	utils.WriteUint32(&buf, id, &offset)
	return buf[:offset]
}
