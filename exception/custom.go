package exception

import "jhr.com/datarelay/model"

type CustomError struct {
	Result model.Result
}

func (ce CustomError) Error() string {
	return ce.Result.Message
}
