package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"minio-s3/internal/logic"
	"minio-s3/internal/svc"
	"minio-s3/internal/types"
)

func ShortenHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), ctx)
		resp, err := l.Shorten(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
