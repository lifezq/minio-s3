package middleware

import (
	"errors"
	"net/http"
	"strings"

	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type Authorization struct {
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func (a *Authorization) AuthorizationHandle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/object") {

			s3Authorization := strings.Trim(r.Header.Get(types.S3_AUTHORIZATION), " ")
			if s3Authorization == "" {
				w.WriteHeader(401)
				httpx.Error(w, errors.New("Forbidden"))
				return
			}

			token, err := types.ParseToken(s3Authorization)
			if err != nil {
				w.WriteHeader(401)
				httpx.Error(w, errors.New("Forbidden"))
				return
			}

			err = token.Valid()
			if err != nil {
				w.WriteHeader(401)
				httpx.Error(w, errors.New("Forbidden"))
				return
			}
		}

		next(w, r)
	}
}
