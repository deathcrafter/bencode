package bencode

import "fmt"

type Belement struct {
	Type  BType
	Value interface{}
}

func getErrorByType(t BType, v Belement) error {
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
		return Belement{}, BencodeError{msg: "Element is not a dict"}
	}

	val, ok := d[key]
	if !ok {
		return Belement{}, BencodeError{msg: fmt.Sprintf("Key %s not found in dict", key)}
	}
	return val, nil
}

func (v Belement) GetDictValueSafe(key string) (Belement, bool) {
	d, err := v.GetDict()
	if nil != err {
		return Belement{}, false
	}

	val, ok := d[key]
	if !ok {
		return Belement{}, false
	}
	return val, true
}
