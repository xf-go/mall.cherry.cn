package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mall.cherry.cn/server/pkg/e"
)

type Response struct {
	C       *gin.Context
	ErrCode int
	Data    interface{}
}

type responseJson struct {
	Code    int         `json:"code"` // TODO
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (res *Response) Out() {
	httpStatus := http.StatusOK
	json := &responseJson{
		Code:    e.SUCCESS,
		Data:    res.Data,
		Message: e.GetMessage(e.SUCCESS),
	}
	if res.ErrCode != e.SUCCESS {
		json.Code = res.ErrCode
		json.Message = e.GetMessage(res.ErrCode)
		json.Data = nil
	}
	res.C.JSON(httpStatus, json)
}
