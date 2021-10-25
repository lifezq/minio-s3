package handler

import (
	"net/http"

	"gitlab.energy-envision.com/storage/internal/logic"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

		req.S3Authorization = r.Header.Get(types.S3_AUTHORIZATION)

		l := logic.NewRemoveLogic(r.Context(), ctx)
		err := l.RemoveObject(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
