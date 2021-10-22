package handler

import (
	"net/http"

	"github.com/lifezq/minio-s3/internal/logic"
	"github.com/lifezq/minio-s3/internal/svc"
	"github.com/lifezq/minio-s3/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("请求发生错误：%v", err)
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
