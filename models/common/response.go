package res

import (
	"fast_gin/utils/valid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// ListResponse 用于列表页面的通用返回值
type ListResponse[T any] struct {
	List  []T `json:"list"`  // 列表数据
	Count any `json:"count"` // 返回的总数
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
func OkWithList[T any](List []T, count any, c *gin.Context) {
	if len(List) == 0 {
		List = []T{}
	}
	Result(SUCCESS, ListResponse[T]{
		List:  List,
		Count: count,
	}, "成功", c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailErrMessage(err error, obj any, c *gin.Context) {
	FailWithMessage(valid.GetValidMsg(err, obj), c)

}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
