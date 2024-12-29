package bencode

import (
	"fmt"
)

type BType int

const (
	TypeInvalid BType = -1
	TypeInt     BType = 0
	TypeString  BType = 1
	TypeList    BType = 2
	TypeDict    BType = 3
)

type Belement struct {
	Type  BType
	Value interface{}
}

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
