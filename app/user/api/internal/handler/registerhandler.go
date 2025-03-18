package handler

import (
	"net/http"

	"PanPan/app/user/api/internal/logic"
	"PanPan/app/user/api/internal/svc"
	"PanPan/app/user/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
