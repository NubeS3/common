//+build windows linux

package utils

const (
	DbError = iota
	Duplicated
	NotFound
	Invalid
	FsError
	Expired
	Timeout
	Other
)

type ErrorType int

type ModelError struct {
	Msg     string
	ErrType ErrorType
}

func (m *ModelError) Error() string {
	return m.Msg
}
