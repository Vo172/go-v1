package login

import (
	"apiproduct/api/internal/svc"
	"apiproduct/api/internal/types"
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(w http.ResponseWriter, req *types.LoginReq) (resp *types.LoginReply, err error) {
	c := l.svcCtx.Keycloak

	jwt, err := c.Gocloak.Login(
		c.ClientId,
		c.ClientSecret,
		c.Realm,
		req.Username,
		req.Password)

	fmt.Println("Token Keycloak : " + jwt.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	return &types.LoginReply{
		Id:           0,
		Name:         req.Username,
		Gender:       "",
		AccessToken:  jwt.AccessToken,
		AccessExpire: int64(jwt.ExpiresIn),
		RefreshAfter: int64(jwt.RefreshExpiresIn),
	}, nil
}
