package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Oracle struct {
		DriverName string
		DataSource string
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Keycloak struct {
		ClientId     string // clientId specified in Keycloak
		ClientSecret string // client secret specified in Keycloak
		Realm        string // realm specified in Keycloak
		BasePath     string // realm specified in Keycloak
	}

	CacheRedis cache.CacheConf
}
