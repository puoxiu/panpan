package handler

import (
	"net/http"

	"PanPan/app/user/api/internal/logic"
	"PanPan/app/user/api/internal/svc"
	"PanPan/app/user/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendcodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterByPhoneRep
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSendcodeLogic(r.Context(), svcCtx)
		resp, err := l.Sendcode(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
