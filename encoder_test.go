package bencode_test

import (
	"testing"

	"github.com/deathcrafter/bencode"
)

func TestEncoderInt(t *testing.T) {
	b := bencode.Belement{Type: bencode.TypeInt, Value: 123}
	e, err := b.Encode()

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "i123e" {
		t.Fatalf("Expected %s, got %s", "i123e", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}

func TestEncoderString(t *testing.T) {
	b := bencode.Belement{Type: bencode.TypeString, Value: "abc"}
	e, err := b.Encode()

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "3:abc" {
		t.Fatalf("Expected %s, got %s", "3:abc", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}

func TestEncoderBencodeList(t *testing.T) {
	b := bencode.Belement{Type: bencode.TypeList, Value: []bencode.Belement{{Type: bencode.TypeInt, Value: 123}, {Type: bencode.TypeString, Value: "abc"}}}
	e, err := b.Encode()

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "li123e3:abce" {
		t.Fatalf("Expected %s, got %s", "li123e3:abce", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}

func TestEncoderBencodeDict(t *testing.T) {
	b := bencode.Belement{
		Type: bencode.TypeDict,
		Value: map[string]bencode.Belement{
			"def": {Type: bencode.TypeString, Value: "abc"},
			"abc": {Type: bencode.TypeInt, Value: 123},
			"mno": {
				Type: bencode.TypeList,
				Value: []bencode.Belement{
					{Type: bencode.TypeInt, Value: 123},
					{Type: bencode.TypeString, Value: "abc"},
				},
			},
		},
	}
	e, err := b.Encode()

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "d3:abci123e3:def3:abc3:mnoli123e3:abcee" {
		t.Fatalf("Expected %s, got %s", "d3:abci123e3:def3:abc3:mnoli123e3:abcee", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}

func TestEncoderNativeList(t *testing.T) {
	b := []interface{}{123, "abc"}
	e, err := bencode.EncodeList(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "li123e3:abce" {
		t.Fatalf("Expected %s, got %s", "li123e3:abce", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}

func TestEncoderNativeDict(t *testing.T) {
	b := map[string]interface{}{"abc": 123, "def": "abc", "mno": []interface{}{123, "abc"}}
	e, err := bencode.EncodeDict(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "d3:abci123e3:def3:abc3:mnoli123e3:abcee" {
		t.Fatalf("Expected %s, got %s", "d3:abci123e3:def3:abc3:mnoli123e3:abcee", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}

func TestEncoderMixed(t *testing.T) {
	b := []interface{}{
		123,
		"abc",
		map[string]interface{}{
			"def": bencode.Belement{Type: bencode.TypeString, Value: "abc"},
			"abc": 123,
			"mno": []bencode.Belement{
				{Type: bencode.TypeInt, Value: 123},
				{Type: bencode.TypeString, Value: "abc"},
			},
		},
	}
	e, err := bencode.EncodeList(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(e) != "li123e3:abcd3:abci123e3:def3:abc3:mnoli123e3:abceee" {
		t.Fatalf("Expected %s, got %s", "li123e3:abcd3:abci123e3:def3:abc3:mnoli123e3:abceee", string(e))
	}

	t.Logf("Encoded: %s", string(e))
}
