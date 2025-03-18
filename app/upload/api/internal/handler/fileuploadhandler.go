package handler

import (
	"net/http"

	"PanPan/app/upload/api/internal/logic"
	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func fileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		err := l.FileUpload(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
