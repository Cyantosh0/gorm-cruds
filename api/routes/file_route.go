package routes

import (
	"github.com/Cyantosh0/gorm-crud/api/controllers"
	"github.com/Cyantosh0/gorm-crud/config"
)

type FileRoute struct {
	handler        config.Router
	fileController controllers.FileController
}

func NewFileRoute(
	handler config.Router,
	fileController controllers.FileController,
) FileRoute {
	return FileRoute{
		handler:        handler,
		fileController: fileController,
	}
}

func (f FileRoute) Setup() {
	file := f.handler.Group("/")

	file.POST("/file", f.fileController.UploadFile)
}
