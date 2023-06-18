package common

import (
	"net/http"

	"crmcore-customer-go/api/internal/logic/common"
	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommonDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailCommonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := common.NewCommonDetailLogic(r.Context(), svcCtx)
		resp, err := l.CommonDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
