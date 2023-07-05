package main

import (
	"github.com/gofiber/fiber/v2"
	"hello-minio/internal/handler/api"
	"hello-minio/internal/repository/service/fstream"
)

func main() {
	err := fstream.Init()
	if err != nil {
		return
	}
	err = fstream.MinioListBucket(fstream.MinioListBucketOptions{})
	if err != nil {
		return
	}
	app := fiber.New()
	api.InitUploadH(app)
	err = app.Listen(":8080")
	if err != nil {
		println(err.Error())
		return
	}
}
