package types

import (
	"fmt"
	"strings"
)

const (
	METADATA_PREFIX = "X-Amz-Meta-"

	OBJECT_TYPE_FOLDER  = "folder"
	OBJECT_TYPE_FILE    = "file"
	OBJECT_TYPE_UNKNOWN = "unknown"
)

func MetadataStripX(meta string) string {
	meta = strings.ToLower(meta)
	if strings.HasPrefix(meta, strings.ToLower(METADATA_PREFIX)) {
		return meta[len(METADATA_PREFIX):]
	}

	return meta
}

func MetadataAddX(meta string) string {
	return fmt.Sprintf("%s%s", METADATA_PREFIX, meta)
}

func PathAndObjectIDFromObjectName(objectName string) (string, string, string) {
	idx := strings.LastIndex(objectName, "/")
	if idx == -1 {
		return OBJECT_TYPE_UNKNOWN, objectName, objectName
	}

	if strings.Trim(objectName[idx+1:], " ") == "" {
		return OBJECT_TYPE_FOLDER, objectName[:idx+1], objectName[idx+1:]
	}

	return OBJECT_TYPE_FILE, objectName[:idx+1], objectName[idx+1:]
}
