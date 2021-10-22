package types

import "fmt"

const (
	BUCKET_PREFIX = "buk"
	GROUP_PREFIX  = "group"
	POLICY_PREFIX = "policy"
	HOME_PREFIX   = "home"

	S3_AUTHORIZATION    = "S3Authorization"
	S3_AUTHORIZATION_OK = "ok"

	CACHE_REDIS_STATE             = "crs"
	CACHE_S3_AUTHORIZATION_PREFIX = "s3.cache.auth"

	ENGINE_MINIO = "MiniO"
)

func BucketName(tenantID, nameSpace string) string {
	return fmt.Sprintf("%s.%s.%s", BUCKET_PREFIX, tenantID, nameSpace)
}

func GroupName(tenantID, nameSpace string) string {
	return fmt.Sprintf("%s.%s.%s", GROUP_PREFIX, tenantID, nameSpace)
}

func PolicyName(tenantID, nameSpace, userID string) string {
	return fmt.Sprintf("%s.%s.%s.%s", POLICY_PREFIX, tenantID, nameSpace, userID)
}

func Home(s string) string {
	return fmt.Sprintf("%s.%s", HOME_PREFIX, s)
}

func BucketUserObjectPath(accessKey string, path, object string) string {
	return fmt.Sprintf("%s.%s/%s/%s", HOME_PREFIX, accessKey, path, object)
}

func CacheS3AuthorizationKey(key string) string {
	return fmt.Sprintf("%s.%s", CACHE_S3_AUTHORIZATION_PREFIX, key)
}
