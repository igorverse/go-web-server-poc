package web

import "strconv"

type Response struct {
	Code  string `json:"code"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewResponse(code int, data any, err string) Response {
	if code < 300 {
		return Response{
			Code:  strconv.FormatInt(int64(code), 10),
			Data:  data,
			Error: "",
		}
	}

	return Response{
		Code:  strconv.FormatInt(int64(code), 10),
		Data:  nil,
		Error: err,
	}
}
