//go:build minio
// +build minio

package client

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/minio/minio-go/v7/pkg/notification"
	"github.com/minio/minio-go/v7/pkg/replication"
	"github.com/minio/minio-go/v7/pkg/sse"
	"github.com/minio/minio-go/v7/pkg/tags"
)

type CopyDestOptions = minio.CopyDestOptions
type CopySrcOptions = minio.CopySrcOptions
type UploadInfo = minio.UploadInfo
type GetObjectOptions = minio.GetObjectOptions
type PutObjectOptions = minio.PutObjectOptions
type RetentionMode = minio.RetentionMode
type ValidityUnit = minio.ValidityUnit
type BucketVersioningConfiguration = minio.BucketVersioningConfiguration
type Object = minio.Object
type ObjectInfo = minio.ObjectInfo
type GetObjectLegalHoldOptions = minio.GetObjectLegalHoldOptions
type LegalHoldStatus = minio.LegalHoldStatus
type GetObjectTaggingOptions = minio.GetObjectTaggingOptions
type BucketInfo = minio.BucketInfo
type ObjectMultipartInfo = minio.ObjectMultipartInfo
type ListObjectsOptions = minio.ListObjectsOptions
type MakeBucketOptions = minio.MakeBucketOptions
type PostPolicy = minio.PostPolicy
type PutObjectLegalHoldOptions = minio.PutObjectLegalHoldOptions
type PutObjectRetentionOptions = minio.PutObjectRetentionOptions
type PutObjectTaggingOptions = minio.PutObjectTaggingOptions
type SnowballOptions = minio.SnowballOptions
type SnowballObject = minio.SnowballObject
type BucketOptions = minio.BucketOptions
type RemoveObjectOptions = minio.RemoveObjectOptions
type RemoveObjectTaggingOptions = minio.RemoveObjectTaggingOptions
type RemoveObjectsOptions = minio.RemoveObjectsOptions
type RemoveObjectError = minio.RemoveObjectError
type RestoreRequest = minio.RestoreRequest
type SelectObjectOptions = minio.SelectObjectOptions
type SelectResults = minio.SelectResults
type StatObjectOptions = minio.StatObjectOptions
type LifecycleConfiguration = lifecycle.Configuration
type NotificationConfiguration = notification.Configuration
type NotificationInfo = notification.Info
type SseConfiguration = sse.Configuration
type TagsTags = tags.Tags
type ReplicationConfig = replication.Config
type ReplicationMetrics = replication.Metrics
type ReplicationResyncTargetsInfo = replication.ResyncTargetsInfo
