package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"PanPan/common/errorx"
	"PanPan/utils"
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

type contextKey string
const (
	userIDKey = contextKey("user_id")
)

// JWTMiddleware jwt middleware
func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// Passthrough to next handler if need
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errorx.NewCodeError(30001, errorx.ErrHeadNil))
			w.Write(err)
			return
		}
		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errorx.NewCodeError(30002, errorx.ErrHeadFormat))
			w.Write(err)
			return
		}
		parseToken, isExpire, err := utils.ParseToken(parts[1], parts[2])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			err, _ := json.Marshal(errorx.NewCodeError(30003, errorx.ErrTokenProve))
			w.Write(err)
			return
		}
		if isExpire {
			// 如果token过期，重新生成token
			parts[1], parts[2] = utils.GetToken(parseToken.ID, parseToken.State)
			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s %s", parts[1], parts[2]))
		}
		// 将用户id放入context
		r = r.WithContext(context.WithValue(r.Context(), userIDKey, parseToken.ID))
		next(w, r)
	}
}
