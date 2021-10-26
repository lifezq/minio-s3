package handler

import (
	"net/http"

	"gitlab.energy-envision.com/storage/internal/logic"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ObjectListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ObjectListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.S3Authorization = r.Header.Get(types.S3_AUTHORIZATION)

		l := logic.NewObjectListLogic(r.Context(), ctx)
		resp, err := l.ObjectList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
