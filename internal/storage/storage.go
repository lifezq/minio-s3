package storage

import (
	"context"
	"io"
	"net/url"
	"time"

	"gitlab.energy-envision.com/storage/client"
)

type LocalStorage struct {
	//TODO
}

func (l LocalStorage) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	panic("implement me")
}

func (l LocalStorage) ComposeObject(ctx context.Context, dst client.CopyDestOptions, srcs ...client.CopySrcOptions) (client.UploadInfo, error) {
	panic("implement me")
}

func (l LocalStorage) CopyObject(ctx context.Context, dst client.CopyDestOptions, src client.CopySrcOptions) (client.UploadInfo, error) {
	panic("implement me")
}

func (l LocalStorage) EnableVersioning(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) EndpointURL() *url.URL {
	panic("implement me")
}

func (l LocalStorage) FGetObject(ctx context.Context, bucketName, objectName, filePath string, opts client.GetObjectOptions) error {
	panic("implement me")
}

func (l LocalStorage) FPutObject(ctx context.Context, bucketName, objectName, filePath string, opts client.PutObjectOptions) (info client.UploadInfo, err error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketEncryption(ctx context.Context, bucketName string) (*client.SseConfiguration, error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketLifecycle(ctx context.Context, bucketName string) (*client.LifecycleConfiguration, error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketLocation(ctx context.Context, bucketName string) (string, error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketNotification(ctx context.Context, bucketName string) (bucketNotification client.NotificationConfiguration, err error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketObjectLockConfig(ctx context.Context, bucketName string) (mode *client.RetentionMode, validity *uint, unit *client.ValidityUnit, err error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketPolicy(ctx context.Context, bucketName string) (string, error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketReplication(ctx context.Context, bucketName string) (cfg client.ReplicationConfig, err error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketReplicationMetrics(ctx context.Context, bucketName string) (s client.ReplicationMetrics, err error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketTagging(ctx context.Context, bucketName string) (*client.TagsTags, error) {
	panic("implement me")
}

func (l LocalStorage) GetBucketVersioning(ctx context.Context, bucketName string) (client.BucketVersioningConfiguration, error) {
	panic("implement me")
}

func (l LocalStorage) GetObject(ctx context.Context, bucketName, objectName string, opts client.GetObjectOptions) (*client.Object, error) {
	panic("implement me")
}

func (l LocalStorage) GetObjectACL(ctx context.Context, bucketName, objectName string) (*client.ObjectInfo, error) {
	panic("implement me")
}

func (l LocalStorage) GetObjectLegalHold(ctx context.Context, bucketName, objectName string, opts client.GetObjectLegalHoldOptions) (status *client.LegalHoldStatus, err error) {
	panic("implement me")
}

func (l LocalStorage) GetObjectLockConfig(ctx context.Context, bucketName string) (objectLock string, mode *client.RetentionMode, validity *uint, unit *client.ValidityUnit, err error) {
	panic("implement me")
}

func (l LocalStorage) GetObjectRetention(ctx context.Context, bucketName, objectName, versionID string) (mode *client.RetentionMode, retainUntilDate *time.Time, err error) {
	panic("implement me")
}

func (l LocalStorage) GetObjectTagging(ctx context.Context, bucketName, objectName string, opts client.GetObjectTaggingOptions) (*client.TagsTags, error) {
	panic("implement me")
}

func (l LocalStorage) HealthCheck(hcDuration time.Duration) (context.CancelFunc, error) {
	panic("implement me")
}

func (l LocalStorage) IsOffline() bool {
	panic("implement me")
}

func (l LocalStorage) IsOnline() bool {
	panic("implement me")
}

func (l LocalStorage) ListBuckets(ctx context.Context) ([]client.BucketInfo, error) {
	panic("implement me")
}

func (l LocalStorage) ListIncompleteUploads(ctx context.Context, bucketName, objectPrefix string, recursive bool) <-chan client.ObjectMultipartInfo {
	panic("implement me")
}

func (l LocalStorage) ListObjects(ctx context.Context, bucketName string, opts client.ListObjectsOptions) <-chan client.ObjectInfo {
	panic("implement me")
}

func (l LocalStorage) ListenBucketNotification(ctx context.Context, bucketName, prefix, suffix string, events []string) <-chan client.NotificationInfo {
	panic("implement me")
}

func (l LocalStorage) ListenNotification(ctx context.Context, prefix, suffix string, events []string) <-chan client.NotificationInfo {
	panic("implement me")
}

func (l LocalStorage) MakeBucket(ctx context.Context, bucketName string, opts client.MakeBucketOptions) (err error) {
	panic("implement me")
}

func (l LocalStorage) Presign(ctx context.Context, method string, bucketName string, objectName string, expires time.Duration, reqParams url.Values) (u *url.URL, err error) {
	panic("implement me")
}

func (l LocalStorage) PresignedGetObject(ctx context.Context, bucketName string, objectName string, expires time.Duration, reqParams url.Values) (u *url.URL, err error) {
	panic("implement me")
}

func (l LocalStorage) PresignedHeadObject(ctx context.Context, bucketName string, objectName string, expires time.Duration, reqParams url.Values) (u *url.URL, err error) {
	panic("implement me")
}

func (l LocalStorage) PresignedPostPolicy(ctx context.Context, p *client.PostPolicy) (u *url.URL, formData map[string]string, err error) {
	panic("implement me")
}

func (l LocalStorage) PresignedPutObject(ctx context.Context, bucketName string, objectName string, expires time.Duration) (u *url.URL, err error) {
	panic("implement me")
}

func (l LocalStorage) PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts client.PutObjectOptions) (info client.UploadInfo, err error) {
	panic("implement me")
}

func (l LocalStorage) PutObjectLegalHold(ctx context.Context, bucketName, objectName string, opts client.PutObjectLegalHoldOptions) error {
	panic("implement me")
}

func (l LocalStorage) PutObjectRetention(ctx context.Context, bucketName, objectName string, opts client.PutObjectRetentionOptions) error {
	panic("implement me")
}

func (l LocalStorage) PutObjectTagging(ctx context.Context, bucketName, objectName string, otags *client.TagsTags, opts client.PutObjectTaggingOptions) error {
	panic("implement me")
}

func (l LocalStorage) PutObjectsSnowball(ctx context.Context, bucketName string, opts client.SnowballOptions, objs <-chan client.SnowballObject) (err error) {
	panic("implement me")
}

func (l LocalStorage) RemoveAllBucketNotification(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) RemoveBucket(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) RemoveBucketEncryption(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) RemoveBucketReplication(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) RemoveBucketTagging(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) RemoveBucketWithOptions(ctx context.Context, bucketName string, opts client.BucketOptions) error {
	panic("implement me")
}

func (l LocalStorage) RemoveIncompleteUpload(ctx context.Context, bucketName, objectName string) error {
	panic("implement me")
}

func (l LocalStorage) RemoveObject(ctx context.Context, bucketName, objectName string, opts client.RemoveObjectOptions) error {
	panic("implement me")
}

func (l LocalStorage) RemoveObjectTagging(ctx context.Context, bucketName, objectName string, opts client.RemoveObjectTaggingOptions) error {
	panic("implement me")
}

func (l LocalStorage) RemoveObjects(ctx context.Context, bucketName string, objectsCh <-chan client.ObjectInfo, opts client.RemoveObjectsOptions) <-chan client.RemoveObjectError {
	panic("implement me")
}

func (l LocalStorage) ResetBucketReplication(ctx context.Context, bucketName string, olderThan time.Duration) (rID string, err error) {
	panic("implement me")
}

func (l LocalStorage) ResetBucketReplicationOnTarget(ctx context.Context, bucketName string, olderThan time.Duration, tgtArn string) (rinfo client.ReplicationResyncTargetsInfo, err error) {
	panic("implement me")
}

func (l LocalStorage) RestoreObject(ctx context.Context, bucketName, objectName, versionID string, req client.RestoreRequest) error {
	panic("implement me")
}

func (l LocalStorage) SelectObjectContent(ctx context.Context, bucketName, objectName string, opts client.SelectObjectOptions) (*client.SelectResults, error) {
	panic("implement me")
}

func (l LocalStorage) SetAppInfo(appName string, appVersion string) {
	panic("implement me")
}

func (l LocalStorage) SetBucketEncryption(ctx context.Context, bucketName string, config *client.SseConfiguration) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketLifecycle(ctx context.Context, bucketName string, config *client.LifecycleConfiguration) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketNotification(ctx context.Context, bucketName string, config client.NotificationConfiguration) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketObjectLockConfig(ctx context.Context, bucketName string, mode *client.RetentionMode, validity *uint, unit *client.ValidityUnit) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketPolicy(ctx context.Context, bucketName, policy string) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketReplication(ctx context.Context, bucketName string, cfg client.ReplicationConfig) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketTagging(ctx context.Context, bucketName string, tags *client.TagsTags) error {
	panic("implement me")
}

func (l LocalStorage) SetBucketVersioning(ctx context.Context, bucketName string, config client.BucketVersioningConfiguration) error {
	panic("implement me")
}

func (l LocalStorage) SetObjectLockConfig(ctx context.Context, bucketName string, mode *client.RetentionMode, validity *uint, unit *client.ValidityUnit) error {
	panic("implement me")
}

func (l LocalStorage) SetS3TransferAccelerate(accelerateEndpoint string) {
	panic("implement me")
}

func (l LocalStorage) StatObject(ctx context.Context, bucketName, objectName string, opts client.GetObjectOptions) (client.ObjectInfo, error) {
	panic("implement me")
}

func (l LocalStorage) SuspendVersioning(ctx context.Context, bucketName string) error {
	panic("implement me")
}

func (l LocalStorage) TraceErrorsOnlyOff() {
	panic("implement me")
}

func (l LocalStorage) TraceErrorsOnlyOn(outputStream io.Writer) {
	panic("implement me")
}

func (l LocalStorage) TraceOff() {
	panic("implement me")
}

func (l LocalStorage) TraceOn(outputStream io.Writer) {
	panic("implement me")
}
