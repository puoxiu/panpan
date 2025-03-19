package handler

import (
	"net/http"

	"PanPan/app/download/api/internal/logic"
	"PanPan/app/download/api/internal/svc"
	"PanPan/app/download/api/internal/types"
	"PanPan/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadCOSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadCOSReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadCOSLogic(r.Context(), svcCtx)
		err := l.DownloadCOS(&req, w, r)
		response.HttpResponse(r, w, nil, err)
	}
}
