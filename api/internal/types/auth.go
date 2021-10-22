package types

type S3AuthorizationToken struct {
	AccessKey       string `json:"accessKey"`
	SecretKey       string `json:"secretKey"`
	Path            string `json:"path"`
	AuthorizationOK string `json:"authorizationOK"`
}
