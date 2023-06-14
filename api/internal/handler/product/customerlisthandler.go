package product

import (
	"net/http"

	"apiproduct/api/internal/logic/product"
	"apiproduct/api/internal/svc"
	"apiproduct/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CustomerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewCustomerListLogic(r.Context(), svcCtx)
		resp, err := l.CustomerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
