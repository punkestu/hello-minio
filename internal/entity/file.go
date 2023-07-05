package entity

import (
	"github.com/minio/minio-go/v7"
	"io"
	"os"
)

type File struct {
	File io.Reader
	Size int64
}

type FileResponse struct {
	File        []byte
	ContentType string
}

func NewFile(file io.Reader, size int64) *File {
	return &File{
		File: file,
		Size: size,
	}
}

func NewFileFromOS(file *os.File) (*File, error) {
	stat, err := file.Stat()
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return &File{
		File: file,
		Size: stat.Size(),
	}, nil
}

func NewFileFromMinio(file *minio.Object) (*File, error) {
	stat, err := file.Stat()
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return &File{
		File: file,
		Size: stat.Size,
	}, nil
}
