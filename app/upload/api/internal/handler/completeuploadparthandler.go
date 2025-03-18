package handler

import (
	"net/http"

	"PanPan/app/upload/api/internal/logic"
	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CompleteUploadPartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteUploadPartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCompleteUploadPartLogic(r.Context(), svcCtx)
		err := l.CompleteUploadPart(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
