package product

import (
	"net/http"

	"apiproduct/api/internal/logic/product"
	"apiproduct/api/internal/svc"
	"apiproduct/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CustomerAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCustomerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewCustomerAddLogic(r.Context(), svcCtx)
		resp, err := l.CustomerAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
