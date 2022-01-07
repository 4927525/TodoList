package serializer

import "net/http"

// 响应数据格式
type Response struct {
	Status uint        `json:"status" example:"1"`
	Msg    string      `json:"msg" example:"msg"`
	Data   interface{} `json:"data" example:"数据"`
	Error  string      `json:"error" example:"错误"`
}

// token相应格式
type TokenResponse struct {
	User interface{} `json:"user"`
	Data string      `json:"data" example:"token"`
}

type DataList struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

// 构建通用分页
func BuildListResponse(items []interface{}, count int) Response {
	return Response{
		Status: http.StatusOK,
		Msg:    "查询成功",
		Data: DataList{
			List:  items,
			Total: count,
		},
	}
}
