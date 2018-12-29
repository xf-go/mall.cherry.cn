package controller

import (
	"github.com/gin-gonic/gin"
	"mall.cherry.cn/server/model"
	"mall.cherry.cn/server/pkg/app"
	"mall.cherry.cn/server/pkg/e"
)

// IndexGet 列表
func GetCategories(c *gin.Context) {
	res := &app.Response{C: c}
	defer res.Out()

	categories, err := model.GetCategories()
	if err != nil {
		res.ErrCode = e.TEST
	}
	res.Data = categories
}
