package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Oracle struct {
		DriverName string
		DataSource string
	}

	Keycloak struct {
		ClientId     string // clientId specified in Keycloak
		ClientSecret string // client secret specified in Keycloak
		Realm        string // realm specified in Keycloak
		BasePath     string // realm specified in Keycloak
	}

	CacheRedis cache.CacheConf

	CommonRpc zrpc.RpcClientConf
}
