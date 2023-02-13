package types

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

type Response struct {
	Success bool        `json:"success"`
	Error   *Error      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func NewResponse(err *Error, res interface{}) *Response {
	return &Response{
		Success: err == nil,
		Error:   err,
		Result:  res,
	}
}

func NewResponseError(code int, v interface{}) *Response {
	message := ""
	if m, ok := v.(string); ok {
		message = m
	} else if m, ok := v.(error); ok {
		message = m.Error()
	}

	return NewResponse(NewError(code, message), nil)
}

func NewResponseResult(v interface{}) *Response {
	return NewResponse(nil, v)
}
