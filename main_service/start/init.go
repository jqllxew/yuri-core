package start

import (
	"github.com/jqllxew/YuriCore/conf"
	"github.com/jqllxew/YuriCore/utils"
)

func Init(exePath string) {
	initConfig(exePath)
	utils.InitConverter(conf.Config.CodePage)
	initUserService()
	initUserCache()
	initRoomService()
}
