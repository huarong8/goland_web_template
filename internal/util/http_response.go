package util

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "",
		Data: data,
	}
}

func FailureResponseCode(code int, err error) *Response {
	return &Response{
		Code: code,
		Msg:  err.Error(),
		Data: "",
	}
}

func FailureResponse(err error) *Response {
	return &Response{
		Code: -1,
		Msg:  err.Error(),
		Data: "",
	}
}
