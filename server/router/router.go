package router

import (
	"github.com/gin-gonic/gin"
	"mall.cherry.cn/server/controller"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.New()

	setRouter(r)

	return r
}

func setRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	category := r.Group("/categories")
	{
		category.GET("/", controller.GetCategories)
	}

	items := r.Group("/items")
	{
		items.GET("/category/:id", controller.GetItemsByCategoryId)
		items.GET("/recommend", controller.GetItemsRecommend)
	}

	item := r.Group("/item")
	{
		item.GET("/:id", controller.GetItemById)
	}
}
