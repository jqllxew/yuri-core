package start

import (
	"github.com/jqllxew/YuriCore/main_service/client"
	"github.com/jqllxew/YuriCore/main_service/infrastructure/user_cache"
)

func initUserCache() {
	client.InitUserCacheClient(user_cache.NewUserCacheImpl())
}
