package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type KeyCloakMiddleware struct {
	keycloak *Keycloak
}

func NewMiddleware(keycloak *Keycloak) *KeyCloakMiddleware {
	return &KeyCloakMiddleware{keycloak: keycloak}
}

func (auth *KeyCloakMiddleware) extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func (auth *KeyCloakMiddleware) VerifyToken(next http.HandlerFunc) http.HandlerFunc {

	f := func(w http.ResponseWriter, r *http.Request) {

		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		token = auth.extractBearerToken(token)

		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := auth.keycloak.Gocloak.RetrospectToken(token, auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwt, _, err := auth.keycloak.Gocloak.DecodeAccessToken(token, auth.keycloak.Realm)

		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}
		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))

		// check if the token isn't expired and valid
		if !result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}

	return http.HandlerFunc(f)
}

func convertToStringSlice(data []interface{}) []string {
	result := make([]string, len(data))
	for i, v := range data {
		if str, ok := v.(string); ok {
			result[i] = str
		}
	}
	return result
}
