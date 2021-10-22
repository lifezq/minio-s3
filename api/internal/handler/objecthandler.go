package handler

import (
	"net/http"

	"github.com/lifezq/minio-s3/api/internal/logic"
	"github.com/lifezq/minio-s3/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func ObjectHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewObjectLogic(r.Context(), ctx)
		resp, err := l.Object()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
