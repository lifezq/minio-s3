package client

import "time"

type CommonObjectInfo struct {
	ETag         string                 `json:"etag"`
	Key          string                 `json:"key"`
	LastModified time.Time              `json:"lastModified"`
	Size         int64                  `json:"size"`
	Metadata     map[string]interface{} `json:"metadata"`
}
