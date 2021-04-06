package errors

var (
	Error ErrorInterface = &Errors{}
)

type ErrorInterface interface {
	CheckPanicError(err error, message ...string)
	CheckFatalError(err error, message ...string)
}

type Errors struct {
	//
}
