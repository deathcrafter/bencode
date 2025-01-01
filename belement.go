package bencode

import "fmt"

type BelementType int

const (
	TypeInvalid BelementType = iota
	TypeInt
	TypeString
	TypeList
	TypeDict
)

type Belement struct {
	Type  BelementType
	Value interface{}
}

var InvalidBelement = Belement{Type: TypeInvalid}

func getErrorByType(t BelementType, v Belement) error {
	if v.Type == TypeInvalid {
		return BencodeError{msg: "Belement is invalid"}
	}
	if v.Type != t {
		return BencodeError{msg: fmt.Sprintf("Value has type %d, expected %d", v.Type, t)}
	}
	return nil
}

func (v Belement) GetInt() (int, error) {
	if err := getErrorByType(TypeInt, v); err != nil {
		return 0, err
	} else {
		return v.Value.(int), nil
	}
}

func (v Belement) GetString() (string, error) {
	if err := getErrorByType(TypeString, v); err != nil {
		return "", err
	} else {
		return v.Value.(string), nil
	}
}

func (v Belement) GetList() ([]Belement, error) {
	if err := getErrorByType(TypeList, v); err != nil {
		return nil, err
	} else {
		return v.Value.([]Belement), nil
	}
}

func (v Belement) GetDict() (map[string]Belement, error) {
	if err := getErrorByType(TypeDict, v); err != nil {
		return nil, err
	} else {
		return v.Value.(map[string]Belement), nil
	}
}

func (v Belement) GetDictValue(key string) (Belement, error) {
	d, err := v.GetDict()
	if err != nil {
		return InvalidBelement, BencodeError{msg: "Belement is not a dict"}
	}

	val, ok := d[key]
	if !ok {
		return InvalidBelement, BencodeError{msg: fmt.Sprintf("Key %s not found in dict", key)}
	}
	return val, nil
}

func (v Belement) GetDictValueSafe(key string) (Belement, bool) {
	d, err := v.GetDict()
	if nil != err {
		return InvalidBelement, false
	}

	val, ok := d[key]
	if !ok {
		return InvalidBelement, false
	}
	return val, true
}
