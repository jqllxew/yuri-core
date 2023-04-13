package out

import (
	"github.com/jqllxew/YuriCore/main_service/constant"
	"github.com/jqllxew/YuriCore/utils"
)

func BuildJoinHost(host uint32) []byte {
	buf := make([]byte, 8)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTJoinHost, &offset)
	utils.WriteUint32(&buf, host, &offset)
	return buf[:offset]
}
