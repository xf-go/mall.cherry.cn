package controller

import (
	"github.com/gin-gonic/gin"
"strconv"

	"mall.cherry.cn/server/model"
	"mall.cherry.cn/server/pkg/e"
	"mall.cherry.cn/server/pkg/app"
)

func GetItemsByCategoryId(c *gin.Context) {
	res := &app.Response{C: c}
	defer res.Out()

	id := c.Param("id")
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res.ErrCode = e.TEST
		return
	}

	categories, err := model.GetItemsByCategoryId(uint(cid))
	if err != nil {
		res.ErrCode = e.TEST
		return
	}

	res.Data = categories
}

func GetItemById(c *gin.Context) {
	res := &app.Response{C: c}
	defer res.Out()

	id := c.Param("id")
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res.ErrCode = e.TEST
		return
	}

	item, err := model.GetItemById(uint(pid))
	if err != nil {
		res.ErrCode = e.TEST
		return
	}

	res.Data = item
}

func GetItemsRecommend(c *gin.Context) {
	res := &app.Response{C: c}
	defer res.Out()

	res.Data = 1
}
	