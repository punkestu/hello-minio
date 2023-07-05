package api

import (
	"github.com/gofiber/fiber/v2"
	"hello-minio/internal/entity"
	"hello-minio/internal/repository/service/fstream"
	"net/http"
)

func InitUploadH(r *fiber.App) {
	g := r.Group("/file")
	g.Get(":fileName", downloadFile)
	g.Post("", uploadFile)
}

func downloadFile(c *fiber.Ctx) error {
	file, err := fstream.MinioBucket("functional").Download(c.Params("fileName"))
	if err != nil {
		println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	c.Set("content-type", file.ContentType)
	return c.Send(file.File)
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("myFile")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fileBuffer, err := file.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	key, err := fstream.MinioBucket("functional").Upload(
		file.Filename,
		entity.NewFile(fileBuffer, file.Size),
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(fiber.Map{
		"key": key,
	})
}
