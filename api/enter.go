package api

import "fast_gin/api/user_api"

type ApiGroup struct {
	UserApi user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
