package start

import (
	"fmt"

	"github.com/jqllxew/YuriCore/conf"
	"github.com/jqllxew/YuriCore/main_service/client"
	"github.com/jqllxew/YuriCore/main_service/infrastructure/user_service"
)

func initUserService() {
	client.InitUserClient(user_service.NewUserService(fmt.Sprintf("%s:%d", conf.Config.UserAdress, conf.Config.UserPort)))
}
