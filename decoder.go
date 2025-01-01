package bencode

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func safeSubbyte(b []byte, start int, end int) []byte {
	if start < 0 {
		start = 0
	}
	if start >= len(b) {
		return []byte{}
	}
	if end < 0 || end >= len(b) {
		return b[start:]
	}
	return b[start:end]
}

func Decode(data []byte) (Belement, error) {
	var decode func([]byte) (Belement, []byte, error)
	decode = func(data []byte) (Belement, []byte, error) {
		belement := Belement{Type: TypeInvalid}

		if len(data) == 0 {
			return belement, nil, BencodeError{msg: "Empty value"}
		}

		for len(data) > 0 {
			switch data[0] {
			case 'i':
				end := bytes.IndexByte(data, 'e')
				if end == -1 {
					return belement, nil, BencodeError{msg: "Invalid integer format: missing end of element"}
				}

				v, err := strconv.Atoi(string(data[1:end]))
				if err != nil {
					return belement, nil, BencodeError{msg: fmt.Sprintf("Invalid integer format: %s", err.Error())}
				}

				belement = Belement{Type: TypeInt, Value: v}
				return belement, safeSubbyte(data, end+1, -1), nil
			case 'l':
				elements := make([]Belement, 0)
				metEnd := false

				data = safeSubbyte(data, 1, -1) // skip 'l'

				for len(data) > 0 {
					if data[0] == 'e' { // end of list
						data = safeSubbyte(data, 1, -1)
						metEnd = true
						break
					}

					elem, newData, err := decode(data)
					if err != nil {
						return belement, nil, err
					}
					data = newData

					elements = append(elements, elem)
				}

				if !metEnd {
					return belement, nil, BencodeError{msg: "Invalid list format: missing end of list"}
				}

				belement = Belement{Type: TypeList, Value: elements}
				return belement, data, nil
			case 'd':
				dict := make(map[string]Belement)
				metEnd := false
				key := ""

				data = safeSubbyte(data, 1, -1) // skip 'd'

				for len(data) > 0 {
					if data[0] == 'e' { // end of dict
						data = safeSubbyte(data, 1, -1)
						metEnd = true
						break
					}

					v, newData, err := decode(data)
					if err != nil {
						return belement, nil, err
					}
					data = newData

					if len(key) == 0 { // key is not set yet, v must be a string
						key, err = v.GetString()
						if err != nil {
							return belement, nil, BencodeError{msg: fmt.Sprintf("Invalid dict key: %s", err.Error())}
						}
					} else { // put in dict and reset key
						dict[key] = v
						key = ""
					}
				}

				if len(key) != 0 {
					return belement, nil, BencodeError{msg: "Invalid dict format: missing value"}
				}

				if !metEnd {
					return belement, nil, BencodeError{msg: "Invalid dict format: missing end of dict"}
				}

				belement = Belement{Type: TypeDict, Value: dict}
				return belement, []byte(data), nil
			default:
				// data must be a string
				raw := string(data)
				eol := strings.Index(raw, ":") // end of length
				if eol == -1 {
					return belement, nil, BencodeError{msg: "Invalid string format"}
				}

				length, err := strconv.Atoi(raw[:eol])
				if err != nil {
					return belement, nil, BencodeError{msg: fmt.Sprintf("Invalid string format. Invalid length: %s", err.Error())}
				}

				if eol+1+length > len(raw) {
					return belement, nil, BencodeError{msg: "Invalid string format. Length mismatch"}
				}

				str := raw[eol+1 : eol+1+length]

				belement = Belement{Type: TypeString, Value: str}
				return belement, safeSubbyte(data, eol+1+length, -1), nil
			}
		}

		return belement, nil, BencodeError{msg: "Unknown type"}
	}

	ret, _, err := decode(data)
	return ret, err
}

func DecodeReader(source Source) (Belement, error) {
	data := make([]byte, 0)
	for {
		slice := make([]byte, 1024)
		n, err := source.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return Belement{}, err
		}
		data = append(data, slice[:n]...)
	}

	return Decode(data)
}
