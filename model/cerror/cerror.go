package cerror

type Error struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	Description   string `json:"description"`
	StatusCode    int    `json:"statusCode"`
	IsRecoverable bool   `json:"isRecoverable"`
}

func NewError(code string, message string, description string, StatusCode int, isRecoverable bool) *Error {
	return &Error{
		Code:          code,
		Message:       message,
		Description:   description,
		StatusCode:    StatusCode,
		IsRecoverable: isRecoverable,
	}
}

func NewFromError(err error) *Error {
	return &Error{
		Code:          "UNEXPECTED_ERROR",
		Message:       err.Error(),
		Description:   "",
		StatusCode:    500,
		IsRecoverable: false,
	}
}

func (e *Error) Error() string {
	return e.Message
}
