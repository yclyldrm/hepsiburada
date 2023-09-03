package app

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.POST("/custom", HandleCustomData)
	router.GET("/sample", HandleSampleData)

	return router
}
