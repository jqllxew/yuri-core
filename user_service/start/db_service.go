package start

import (
	"github.com/jqllxew/YuriCore/conf"
	"github.com/jqllxew/YuriCore/user_service/client"
	"github.com/jqllxew/YuriCore/user_service/infrastructure/db_service"
)

func InitDBService() {
	client.InitDBClient(db_service.NewDBImpl(
		conf.Config.DBUserName,
		conf.Config.DBpassword,
		conf.Config.DBaddress,
		conf.Config.DBport,
	))
}
