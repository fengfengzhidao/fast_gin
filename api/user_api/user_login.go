package user_api

import "github.com/gin-gonic/gin"

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// UserLoginView 登录，返回token，用户信息需要从token中解码
// @Tags 用户管理
// @Summary 登录，返回token
// @Description 登录，返回token
// @Param data body UserLoginRequest  true  "登录的一些参数"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLoginView(c *gin.Context) {

}
