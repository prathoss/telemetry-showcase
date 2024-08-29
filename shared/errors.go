package shared

var _ error = ErrNotFound{}

func NewErrNotFound(msg string) ErrNotFound {
	return ErrNotFound{msg}
}

type ErrNotFound struct {
	msg string
}

func (e ErrNotFound) Error() string {
	return e.msg
}
