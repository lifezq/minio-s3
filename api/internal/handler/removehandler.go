package handler

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"

	"github.com/lifezq/minio-s3/api/internal/logic"
	"github.com/lifezq/minio-s3/api/internal/svc"
	"github.com/lifezq/minio-s3/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func RemoveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("请求发生错误：%v", err)
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveLogic(r.Context(), ctx)
		err := l.RemoveObject(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
