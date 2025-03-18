package handler

import (
	"net/http"

	"PanPan/app/user/api/internal/logic"
	"PanPan/app/user/api/internal/svc"
	"PanPan/common/response"

)

func GithubCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGithubCallbackLogic(r.Context(), svcCtx)
		resp, err := l.GithubCallback()
		response.HttpResponse(r, w, resp, err)
	}
}
