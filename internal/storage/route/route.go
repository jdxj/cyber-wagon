package route

import (
	"github.com/gin-gonic/gin"

	"github.com/jdxj/cyber-wagon/config"
	"github.com/jdxj/cyber-wagon/internal/storage/model"
)

var stg *model.Storage

func Init(cfg config.Storage) {
	stg = model.NewStorage(cfg)
}

func RegisterRoute(root gin.IRouter) {
	// todo: check token
	//root.Use()

	rFiles := root.Group("/files")
	{
		rFiles.POST("/", upload)
		rFiles.GET("/:file_id", download)
	}
}
