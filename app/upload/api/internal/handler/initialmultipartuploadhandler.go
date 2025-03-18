package handler

import (
	"net/http"

	"PanPan/app/upload/api/internal/logic"
	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func initialMultipartUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitialMultipartUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInitialMultipartUploadLogic(r.Context(), svcCtx)
		resp, err := l.InitialMultipartUpload(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
