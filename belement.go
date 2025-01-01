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

type BelementInt interface {
	GetInt() (int, error)
}

type BelementString interface {
	GetString() (string, error)
}

type BelementList interface {
	GetList() ([]Belement, error)
	GetListValue(int) (Belement, error)
	GetListInt(int) (int, error)
	GetListString(int) (string, error)
	GetListList(int) ([]Belement, error)
	GetListDict(int) (map[string]Belement, error)
}

type BelementIntList interface {
	GetIntList() ([]int, error)
}

type BelementStringList interface {
	GetStringList() ([]string, error)
}

type BelementAnyList interface {
	GetAnyList() ([]any, error)
}

type BelementDict interface {
	GetDict() (map[string]Belement, error)
	GetDictValue(string) (Belement, error)
	GetDictInt(string) (int, error)
	GetDictString(string) (string, error)
	GetDictList(string) ([]Belement, error)
	GetDictDict(string) (map[string]Belement, error)
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

func (v Belement) GetListValue(index int) (Belement, error) {
	l, err := v.GetList()
	if err != nil {
		return InvalidBelement, BencodeError{msg: "Belement is not a list"}
	}
	if index >= len(l) {
		return InvalidBelement, BencodeError{msg: fmt.Sprintf("Index %d out of range", index)}
	}
	return l[index], nil
}

func (v Belement) GetListInt(index int) (int, error) {
	val, err := v.GetListValue(index)
	if err != nil {
		return 0, err
	}
	return val.GetInt()
}

func (v Belement) GetListString(index int) (string, error) {
	val, err := v.GetListValue(index)
	if err != nil {
		return "", err
	}
	return val.GetString()
}

func (v Belement) GetListList(index int) ([]Belement, error) {
	val, err := v.GetListValue(index)
	if err != nil {
		return nil, err
	}
	return val.GetList()
}

func (v Belement) GetListDict(index int) (map[string]Belement, error) {
	val, err := v.GetListValue(index)
	if err != nil {
		return nil, err
	}
	return val.GetDict()
}

func (v Belement) GetIntList() ([]int, error) {
	l, err := v.GetList()
	if err != nil {
		return nil, err
	}
	x := make([]int, 0)
	for _, v := range l {
		i, err := v.GetInt()
		if err != nil {
			return nil, err
		}
		x = append(x, i)
	}
	return x, nil
}

func (v Belement) GetStringList() ([]string, error) {
	l, err := v.GetList()
	if err != nil {
		return nil, err
	}
	x := make([]string, 0)
	for _, v := range l {
		s, err := v.GetString()
		if err != nil {
			return nil, err
		}
		x = append(x, s)
	}
	return x, nil
}

func (v Belement) GetAnyList() ([]any, error) {
	l, err := v.GetList()
	if err != nil {
		return nil, err
	}
	x := make([]any, 0)
	for _, v := range l {
		x = append(x, v.Value)
	}
	return x, nil
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

func (v Belement) GetDictInt(key string) (int, error) {
	val, err := v.GetDictValue(key)
	if err != nil {
		return 0, err
	}
	return val.GetInt()
}

func (v Belement) GetDictString(key string) (string, error) {
	val, err := v.GetDictValue(key)
	if err != nil {
		return "", err
	}
	return val.GetString()
}

func (v Belement) GetDictList(key string) ([]Belement, error) {
	val, err := v.GetDictValue(key)
	if err != nil {
		return nil, err
	}
	return val.GetList()
}

func (v Belement) GetDictDict(key string) (map[string]Belement, error) {
	val, err := v.GetDictValue(key)
	if err != nil {
		return nil, err
	}
	return val.GetDict()
}
