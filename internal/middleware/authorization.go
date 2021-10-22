package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/lifezq/minio-s3/internal/config"
	"github.com/lifezq/minio-s3/internal/types"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

type Authorization struct {
	cache cache.Cache
}

func NewAuthorization(c config.Config) *Authorization {
	return &Authorization{
		cache: cache.New(c.CacheConf, syncx.NewSingleFlight(), cache.NewStat(types.CACHE_REDIS_STATE),
			sql.ErrNoRows, []cache.Option{}...),
	}
}

func (a *Authorization) AuthorizationHandle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/object") {

			if strings.Trim(r.Header.Get(types.S3_AUTHORIZATION), " ") == "" {
				w.WriteHeader(401)
				httpx.Error(w, errors.New("Forbidden"))
				return
			}

			var token types.S3AuthorizationToken
			err := a.cache.Get(types.CacheS3AuthorizationKey(r.Header.Get(types.S3_AUTHORIZATION)), &token)
			if err != nil || token.AuthorizationOK != types.S3_AUTHORIZATION_OK {
				w.WriteHeader(401)
				httpx.Error(w, errors.New("Forbidden"))
				return
			}
		}

		next(w, r)
	}
}
