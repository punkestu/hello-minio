package fstream

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"hello-minio/internal/repository/bucket"
)

const (
	endpoint        = "URL"
	accessKeyID     = "ACCESSKEYID"
	secretAccessKey = "SECRETACCESSKEY"
)

var minioClient *minio.Client
var mBuckets = map[string]*bucket.Bucket{}

func Init() (err error) {
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	return
}

type MinioListBucketOptions struct {
	Verbose bool
}

func MinioListBucket(opt MinioListBucketOptions) error {
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		return err
	}
	for i, b := range buckets {
		if opt.Verbose {
			println(i, b.Name)
		}
		mBuckets[b.Name] = bucket.NewBucket(minioClient, b.Name)
	}
	return nil
}

func MinioCreateBucket(bucketName string) error {
	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
		Region:        "",
		ObjectLocking: false,
	})
	if err != nil {
		return err
	}
	mBuckets[bucketName] = bucket.NewBucket(minioClient, bucketName)
	return nil
}

func MinioBucket(bucketName string) *bucket.Bucket {
	return mBuckets[bucketName]
}
