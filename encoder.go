package bencode

import (
	"fmt"
	"sort"
	"strconv"
)

func EncodeInt(v int) ([]byte, error) {
	ret := make([]byte, 0)
	ret = append(ret, byte('i'))
	ret = append(ret, []byte(strconv.Itoa(v))...)
	ret = append(ret, byte('e'))
	return ret, nil
}

func EncodeString(v string) ([]byte, error) {
	ret := make([]byte, 0)
	ret = append(ret, []byte(strconv.Itoa(len(v)))...)
	ret = append(ret, byte(':'))
	ret = append(ret, []byte(v)...)
	return ret, nil
}

func EncodeList(b []interface{}) ([]byte, error) {
	ret := make([]byte, 0)
	ret = append(ret, byte('l'))

	// Empty list
	if len(b) == 0 {
		ret = append(ret, byte('e'))
		return ret, nil
	}

	// since Go slices can contain only one type, we can switch over the type of the first element
	for _, v := range b {
		switch t := v.(type) {
		case int:
			e, err := EncodeInt(v.(int))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case string:
			e, err := EncodeString(v.(string))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case map[string]interface{}:
			e, err := EncodeDict(v.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case []interface{}:
			e, err := EncodeList(v.([]interface{}))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case Belement:
			e, err := (v.(Belement)).Encode()
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		default:
			return nil, BencodeError{msg: fmt.Sprintf("Expected list value of int, string or Belement. Recieved %T", t)}
		}
	}

	ret = append(ret, byte('e'))
	return ret, nil
}

func EncodeDict(b map[string]interface{}) ([]byte, error) {
	ret := make([]byte, 0)
	ret = append(ret, byte('d'))

	// Empty dict
	if len(b) == 0 {
		ret = append(ret, byte('e'))
		return ret, nil
	}

	// Bencode specs require dict keys to be sorted
	sorted_keys := make(sort.StringSlice, 0)
	for k := range b {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Sort(sorted_keys)

	for _, k := range sorted_keys {
		v := b[k]
		e, err := EncodeString(k)
		if err != nil {
			return nil, err
		}
		ret = append(ret, e...)

		switch t := v.(type) {
		case int:
			e, err := EncodeInt(v.(int))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case string:
			e, err := EncodeString(v.(string))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case Belement:
			e, err := (v.(Belement)).Encode()
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case []Belement:
			x := make([]interface{}, 0)
			for _, v := range v.([]Belement) {
				x = append(x, v)
			}
			e, err := EncodeList(x)
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case []interface{}:
			e, err := EncodeList(v.([]interface{}))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case map[string]Belement:
			x := make(map[string]interface{}, 0)
			for k, val := range v.(map[string]Belement) {
				x[k] = val
			}
			e, err := EncodeDict(x)
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		case map[string]interface{}:
			e, err := EncodeDict(v.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			ret = append(ret, e...)
		default:
			return nil, BencodeError{msg: fmt.Sprintf("Expected dict value of int, string or Belement. Recieved %T", t)}
		}
	}

	ret = append(ret, byte('e'))
	return ret, nil
}

func (v Belement) Encode() ([]byte, error) {
	switch v.Type {
	case TypeInt:
		r, _ := v.GetInt()
		return EncodeInt(r)
	case TypeString:
		r, _ := v.GetString()
		return EncodeString(r)
	case TypeList:
		x := make([]interface{}, 0)
		r, _ := v.GetList()
		for _, v := range r {
			x = append(x, v.Value)
		}
		return EncodeList(x)
	case TypeDict:
		x := make(map[string]interface{}, 0)
		r, _ := v.GetDict()
		for k, v := range r {
			x[k] = v
		}
		return EncodeDict(x)
	case TypeInvalid:
		return nil, BencodeError{msg: "Belement is invalid"}
	default:
		return nil, BencodeError{
			msg: fmt.Sprintf(
				"Unknown type: %d. Allowed types: { int: %d, string: %d, list: %d, dict: %d }",
				v.Type,
				TypeInt,
				TypeString,
				TypeList,
				TypeDict,
			),
		}
	}
}
