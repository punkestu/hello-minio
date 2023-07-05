package bucket

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/savsgio/gotils/uuid"
	"hello-minio/internal/entity"
	"net/http"
	"path"
)

type Bucket struct {
	client *minio.Client
	name   string
}

func NewBucket(client *minio.Client, name string) *Bucket {
	return &Bucket{
		client: client,
		name:   name,
	}
}

func (b Bucket) Download(name string) (*entity.FileResponse, error) {
	if object, err := b.client.GetObject(context.Background(), b.name, name, minio.GetObjectOptions{}); err != nil {
		return nil, err
	} else {
		fileInfo, err := object.Stat()
		if err != nil {
			return nil, err
		}
		fileByte := make([]byte, fileInfo.Size)
		_, err = object.Read(fileByte)
		if err != nil {
			return nil, err
		}
		contentType := http.DetectContentType(fileByte)
		return &entity.FileResponse{
			File:        fileByte,
			ContentType: contentType,
		}, nil
	}
}

func (b Bucket) Upload(name string, file *entity.File) (*string, error) {
	object, err := b.client.PutObject(
		context.Background(),
		b.name, uuid.V4()+path.Ext(name),
		file.File, file.Size,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return nil, err
	} else {
		return &object.Key, nil
	}
}
