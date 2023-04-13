package start

import "github.com/jqllxew/YuriCore/conf"

func Init(exePath string) {
	initConfig(exePath)
	if conf.Config.EnableDataBase == 1 {
		InitDBService()
	}
	initUserTable()
}
