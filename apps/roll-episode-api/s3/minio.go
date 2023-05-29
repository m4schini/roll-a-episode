package s3

import (
	"context"
	"github.com/m4schini/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometheus/client_golang/prometheus"
	"io"
)

var log = logger.Named("s3").Sugar()

type MinioBucket struct {
	client     *minio.Client
	bucketName string
}

func (m *MinioBucket) Put(ctx context.Context, name, contentType string, size int64, objectReader io.Reader) error {
	info, err := m.client.PutObject(ctx, m.bucketName, name, objectReader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}
	log.Infof("Put object(%v) in bucket(%v): %v", name, m.bucketName, info)
	return nil
}

func NewMinioBucket(ctx context.Context, bucketName, endpoint, accessKey, accessSecret string, useSSL bool) (*MinioBucket, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	minioBucket := &MinioBucket{client: minioClient, bucketName: bucketName}

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if exists {
		return minioBucket, nil
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return nil, err
	}
	return minioBucket, nil
}

var _minioPutDurationHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace:   "s3",
	Subsystem:   "minio",
	Name:        "put",
	Help:        "",
	ConstLabels: nil,
})

func init() {
	prometheus.MustRegister(_minioPutDurationHistogram)
}
