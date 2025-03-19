package handler

import (
	"net/http"

	"PanPan/app/download/api/internal/logic"
	"PanPan/app/download/api/internal/svc"
	"PanPan/app/download/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadMinioHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadMinioReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadMinioLogic(r.Context(), svcCtx)
		err := l.DownloadMinio(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
