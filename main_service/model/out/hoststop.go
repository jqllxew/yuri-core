package out

import (
	"github.com/jqllxew/YuriCore/main_service/constant"
	"github.com/jqllxew/YuriCore/utils"
)

func BuildHostStop() []byte {
	buf := make([]byte, 1)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTHostStop, &offset)

	return buf[:offset]
}
