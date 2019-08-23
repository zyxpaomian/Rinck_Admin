package error

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func New(msg string) *MyError {
	return &MyError{msg: msg}
}

func DBError() *MyError {
	return New("数据库错误")
}

func AuthError() *MyError {
	return New("用户认证错误")
}