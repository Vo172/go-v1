package middleware

import "github.com/Nerzal/gocloak"

type Keycloak struct {
	Gocloak      gocloak.GoCloak // keycloak client
	ClientId     string          // clientId specified in Keycloak
	ClientSecret string          // client secret specified in Keycloak
	Realm        string          // realm specified in Keycloak
}

func NewKeycloak(basePath string, clientId string, clientSecret string, realm string) *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient(basePath),
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Realm:        realm,
	}
}
