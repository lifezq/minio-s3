package handler

import (
	"net/http"

	"gitlab.energy-envision.com/storage/internal/logic"
	"gitlab.energy-envision.com/storage/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ObjectHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewObjectLogic(r.Context(), ctx)
		resp, err := l.Object(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
