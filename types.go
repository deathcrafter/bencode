package bencode

import (
	"fmt"
)

type BencodeError struct {
	msg  string
	data interface{}
}

func (e BencodeError) Error() string {
	switch e.data.(type) {
	case nil:
		return e.msg
	default:
		return fmt.Sprintf(e.msg, e.data)
	}
}

type Source interface {
	Read(b []byte) (n int, err error)
}
