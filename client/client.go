package client

import (
	"context"
	"io"
	"net/url"
	"time"
)

type Client interface {
	BucketExists(ctx context.Context, bucketName string) (bool, error)
	ComposeObject(ctx context.Context, dst CopyDestOptions, srcs ...CopySrcOptions) (UploadInfo, error)
	CopyObject(ctx context.Context, dst CopyDestOptions, src CopySrcOptions) (UploadInfo, error)
	EnableVersioning(ctx context.Context, bucketName string) error
	EndpointURL() *url.URL
	FGetObject(ctx context.Context, bucketName, objectName, filePath string, opts GetObjectOptions) error
	FPutObject(ctx context.Context, bucketName, objectName, filePath string, opts PutObjectOptions) (info UploadInfo, err error)
	GetBucketEncryption(ctx context.Context, bucketName string) (*SseConfiguration, error)
	GetBucketLifecycle(ctx context.Context, bucketName string) (*LifecycleConfiguration, error)
	GetBucketLocation(ctx context.Context, bucketName string) (string, error)
	GetBucketNotification(ctx context.Context, bucketName string) (bucketNotification NotificationConfiguration, err error)
	GetBucketObjectLockConfig(ctx context.Context, bucketName string) (mode *RetentionMode, validity *uint, unit *ValidityUnit, err error)
	GetBucketPolicy(ctx context.Context, bucketName string) (string, error)
	GetBucketReplication(ctx context.Context, bucketName string) (cfg ReplicationConfig, err error)
	GetBucketReplicationMetrics(ctx context.Context, bucketName string) (s ReplicationMetrics, err error)
	GetBucketTagging(ctx context.Context, bucketName string) (*TagsTags, error)
	GetBucketVersioning(ctx context.Context, bucketName string) (BucketVersioningConfiguration, error)
	GetObject(ctx context.Context, bucketName, objectName string, opts GetObjectOptions) (*Object, error)
	GetObjectACL(ctx context.Context, bucketName, objectName string) (*ObjectInfo, error)
	GetObjectLegalHold(ctx context.Context, bucketName, objectName string, opts GetObjectLegalHoldOptions) (status *LegalHoldStatus, err error)
	GetObjectLockConfig(ctx context.Context, bucketName string) (objectLock string, mode *RetentionMode, validity *uint, unit *ValidityUnit, err error)
	GetObjectRetention(ctx context.Context, bucketName, objectName, versionID string) (mode *RetentionMode, retainUntilDate *time.Time, err error)
	GetObjectTagging(ctx context.Context, bucketName, objectName string, opts GetObjectTaggingOptions) (*TagsTags, error)
	HealthCheck(hcDuration time.Duration) (context.CancelFunc, error)
	IsOffline() bool
	IsOnline() bool
	ListBuckets(ctx context.Context) ([]BucketInfo, error)
	ListIncompleteUploads(ctx context.Context, bucketName, objectPrefix string, recursive bool) <-chan ObjectMultipartInfo
	ListObjects(ctx context.Context, bucketName string, opts ListObjectsOptions) <-chan ObjectInfo
	ListenBucketNotification(ctx context.Context, bucketName, prefix, suffix string, events []string) <-chan NotificationInfo
	ListenNotification(ctx context.Context, prefix, suffix string, events []string) <-chan NotificationInfo
	MakeBucket(ctx context.Context, bucketName string, opts MakeBucketOptions) (err error)
	Presign(ctx context.Context, method string, bucketName string, objectName string, expires time.Duration, reqParams url.Values) (u *url.URL, err error)
	PresignedGetObject(ctx context.Context, bucketName string, objectName string, expires time.Duration, reqParams url.Values) (u *url.URL, err error)
	PresignedHeadObject(ctx context.Context, bucketName string, objectName string, expires time.Duration, reqParams url.Values) (u *url.URL, err error)
	PresignedPostPolicy(ctx context.Context, p *PostPolicy) (u *url.URL, formData map[string]string, err error)
	PresignedPutObject(ctx context.Context, bucketName string, objectName string, expires time.Duration) (u *url.URL, err error)
	PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts PutObjectOptions) (info UploadInfo, err error)
	PutObjectLegalHold(ctx context.Context, bucketName, objectName string, opts PutObjectLegalHoldOptions) error
	PutObjectRetention(ctx context.Context, bucketName, objectName string, opts PutObjectRetentionOptions) error
	PutObjectTagging(ctx context.Context, bucketName, objectName string, otags *TagsTags, opts PutObjectTaggingOptions) error
	PutObjectsSnowball(ctx context.Context, bucketName string, opts SnowballOptions, objs <-chan SnowballObject) (err error)
	RemoveAllBucketNotification(ctx context.Context, bucketName string) error
	RemoveBucket(ctx context.Context, bucketName string) error
	RemoveBucketEncryption(ctx context.Context, bucketName string) error
	RemoveBucketReplication(ctx context.Context, bucketName string) error
	RemoveBucketTagging(ctx context.Context, bucketName string) error
	RemoveBucketWithOptions(ctx context.Context, bucketName string, opts BucketOptions) error
	RemoveIncompleteUpload(ctx context.Context, bucketName, objectName string) error
	RemoveObject(ctx context.Context, bucketName, objectName string, opts RemoveObjectOptions) error
	RemoveObjectTagging(ctx context.Context, bucketName, objectName string, opts RemoveObjectTaggingOptions) error
	RemoveObjects(ctx context.Context, bucketName string, objectsCh <-chan ObjectInfo, opts RemoveObjectsOptions) <-chan RemoveObjectError
	ResetBucketReplication(ctx context.Context, bucketName string, olderThan time.Duration) (rID string, err error)
	ResetBucketReplicationOnTarget(ctx context.Context, bucketName string, olderThan time.Duration, tgtArn string) (rinfo ReplicationResyncTargetsInfo, err error)
	RestoreObject(ctx context.Context, bucketName, objectName, versionID string, req RestoreRequest) error
	SelectObjectContent(ctx context.Context, bucketName, objectName string, opts SelectObjectOptions) (*SelectResults, error)
	SetAppInfo(appName string, appVersion string)
	SetBucketEncryption(ctx context.Context, bucketName string, config *SseConfiguration) error
	SetBucketLifecycle(ctx context.Context, bucketName string, config *LifecycleConfiguration) error
	SetBucketNotification(ctx context.Context, bucketName string, config NotificationConfiguration) error
	SetBucketObjectLockConfig(ctx context.Context, bucketName string, mode *RetentionMode, validity *uint, unit *ValidityUnit) error
	SetBucketPolicy(ctx context.Context, bucketName, policy string) error
	SetBucketReplication(ctx context.Context, bucketName string, cfg ReplicationConfig) error
	SetBucketTagging(ctx context.Context, bucketName string, tags *TagsTags) error
	SetBucketVersioning(ctx context.Context, bucketName string, config BucketVersioningConfiguration) error
	SetObjectLockConfig(ctx context.Context, bucketName string, mode *RetentionMode, validity *uint, unit *ValidityUnit) error
	SetS3TransferAccelerate(accelerateEndpoint string)
	StatObject(ctx context.Context, bucketName, objectName string, opts StatObjectOptions) (ObjectInfo, error)
	SuspendVersioning(ctx context.Context, bucketName string) error
	TraceErrorsOnlyOff()
	TraceErrorsOnlyOn(outputStream io.Writer)
	TraceOff()
	TraceOn(outputStream io.Writer)
}
