package start

import (
	"github.com/jqllxew/YuriCore/user_service/client"
	"github.com/jqllxew/YuriCore/user_service/infrastructure/userTable"
)

func initUserTable() {
	client.InitUserTableClient(userTable.NewUserTableImpl())
}
