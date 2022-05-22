package util

import "douyin/entity"

var TokenFailResponse = entity.Response{
	StatusCode: 400,
	StatusMsg:  "未登录或登录信息过期",
}

var SuccessResponse = entity.Response{
	StatusCode: 0,
	StatusMsg:  "成功",
}

var ServerErrorResponse = entity.Response{
	StatusCode: 500,
	StatusMsg:  "服务器内部错误",
}

var ParamErrorResponse = entity.Response{
	StatusCode: 400,
	StatusMsg:  "参数错误",
}
