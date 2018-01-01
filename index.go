package Elastic

import (
	Gin "github.com/gin-gonic/gin"
	"github.com/go-crazy/elastic/Routers"
)


func InitElastic(router *Gin.RouterGroup) {

		router.POST("/set", Routers.SetHandler)

		// router.POST("/get", func(c *Gin.Context)  {
		// 	Routers.GetHandler(c.Writer, c.Request)
		// })
		router.POST("/get", Routers.GetHandler)
		// http://blog.csdn.net/xsdxs/article/details/72849796
	    router.POST("/map", Routers.MappingHandler)
}

func main()  {
	engine := Gin.Default()
	wsElastic := engine.Group("test-elastic")
	InitElastic(wsElastic)
	engine.Run(":8081")
}