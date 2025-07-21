package model

type BaseRes[T any] struct {
	Code int    `json:"code" dc:"响应码，100000表示成功，其他表示异常"`
	Msg  string `json:"msg" dc:"响应信息"`
	Data T      `json:"data" dc:"响应数据"`
}
