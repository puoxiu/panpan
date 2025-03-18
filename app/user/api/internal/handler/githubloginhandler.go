package handler

import (
	"net/http"

	"PanPan/app/user/api/internal/logic"
	"PanPan/app/user/api/internal/svc"
	"PanPan/common/response"

)

func GithubLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGithubLoginLogic(r.Context(), svcCtx)
		err := l.GithubLogin()
		response.HttpResponse(r, w, nil, err)
	}
}
