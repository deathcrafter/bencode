package bencode_test

import (
	"testing"

	"github.com/deathcrafter/bencode"
)

func TestDecoderInt(t *testing.T) {
	b, err := bencode.Decode([]byte("i123e"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeInt {
		t.Fatalf("Expected type %d, got %d", bencode.TypeInt, b.Type)
	}

	if b.Value.(int) != 123 {
		t.Fatalf("Expected value %d, got %d", 123, b.Value.(int))
	}

	t.Logf("Int val: %d", b.Value.(int))
}

func TestDecoderString(t *testing.T) {
	b, err := bencode.Decode([]byte("3:abc"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeString {
		t.Fatalf("Expected type %d, got %d", bencode.TypeString, b.Type)
	}

	if b.Value.(string) != "abc" {
		t.Fatalf("Expected value %s, got %s", "abc", b.Value.(string))
	}

	t.Logf("String val: %s", b.Value.(string))
}

func TestDecoderEmptyList(t *testing.T) {
	b, err := bencode.Decode([]byte("le"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeList {
		t.Fatalf("Expected type %d, got %d", bencode.TypeList, b.Type)
	}

	if len(b.Value.([]bencode.Belement)) != 0 {
		t.Fatalf("Expected empty list, got %v", b.Value.([]bencode.Belement))
	}

	t.Logf("List val: %v", b.Value.([]bencode.Belement))
}

func TestDecoderList(t *testing.T) {
	b, err := bencode.Decode([]byte("li1e3:abce"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeList {
		t.Fatalf("Expected type %d, got %d", bencode.TypeList, b.Type)
	}

	if len(b.Value.([]bencode.Belement)) != 2 {
		t.Fatalf("Expected list with 2 elements, got %v", b.Value.([]bencode.Belement))
	}

	if b.Value.([]bencode.Belement)[0].Type != bencode.TypeInt {
		t.Fatalf("Expected type %d, got %d", bencode.TypeInt, b.Value.([]bencode.Belement)[0].Type)
	}

	if b.Value.([]bencode.Belement)[0].Value.(int) != 1 {
		t.Fatalf("Expected value %d, got %d", 1, b.Value.([]bencode.Belement)[0].Value.(int))
	}

	if b.Value.([]bencode.Belement)[1].Type != bencode.TypeString {
		t.Fatalf("Expected type %d, got %d", bencode.TypeString, b.Value.([]bencode.Belement)[1].Type)
	}

	if b.Value.([]bencode.Belement)[1].Value.(string) != "abc" {
		t.Fatalf("Expected value %s, got %s", "abc", b.Value.([]bencode.Belement)[1].Value.(string))
	}

	t.Logf("List val: %v", b.Value.([]bencode.Belement))
}

func TestDecoderEmptyDict(t *testing.T) {
	b, err := bencode.Decode([]byte("de"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeDict {
		t.Fatalf("Expected type %d, got %d", bencode.TypeDict, b.Type)
	}

	if len(b.Value.(map[string]bencode.Belement)) != 0 {
		t.Fatalf("Expected empty dict, got %v", b.Value.(map[string]bencode.Belement))
	}

	t.Logf("Dict val: %v", b.Value.(map[string]bencode.Belement))
}

func TestDecoderDict(t *testing.T) {
	b, err := bencode.Decode([]byte("d3:abc3:defe"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeDict {
		t.Fatalf("Expected type %d, got %d", bencode.TypeDict, b.Type)
	}

	if len(b.Value.(map[string]bencode.Belement)) != 1 {
		t.Fatalf("Expected dict with 1 elements, got %v", b.Value.(map[string]bencode.Belement))
	}

	if b.Value.(map[string]bencode.Belement)["abc"].Type != bencode.TypeString {
		t.Fatalf("Expected type %d, got %d", bencode.TypeString, b.Value.(map[string]bencode.Belement)["abc"].Type)
	}

	if b.Value.(map[string]bencode.Belement)["abc"].Value.(string) != "def" {
		t.Fatalf("Expected value %s, got %s", "def", b.Value.(map[string]bencode.Belement)["abc"].Value.(string))
	}

	t.Logf("Dict val: %v", b.Value.(map[string]bencode.Belement))
}

func TestDecoderListDict(t *testing.T) {
	b, err := bencode.Decode([]byte("d3:vall3:abci1eee"))
	if err != nil {
		t.Fatal(err)
	}

	if b.Type != bencode.TypeDict {
		t.Fatalf("Expected type %d, got %d", bencode.TypeDict, b.Type)
	}

	if len(b.Value.(map[string]bencode.Belement)) != 1 {
		t.Fatalf("Expected dict with 1 elements, got %v", b.Value.(map[string]bencode.Belement))
	}

	if b.Value.(map[string]bencode.Belement)["val"].Type != bencode.TypeList {
		t.Fatalf("Expected type %d, got %d", bencode.TypeList, b.Value.(map[string]bencode.Belement)["val"].Type)
	}

	if len(b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)) != 2 {
		t.Fatalf("Expected list with 2 elements, got %v", b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement))
	}

	if b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[0].Type != bencode.TypeString {
		t.Fatalf("Expected type %d, got %d", bencode.TypeString, b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[0].Type)
	}

	if b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[0].Value.(string) != "abc" {
		t.Fatalf("Expected value %s, got %s", "abc", b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[0].Value.(string))
	}

	if b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[1].Type != bencode.TypeInt {
		t.Fatalf("Expected type %d, got %d", bencode.TypeInt, b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[1].Type)
	}

	if b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[1].Value.(int) != 1 {
		t.Fatalf("Expected value %d, got %d", 1, b.Value.(map[string]bencode.Belement)["val"].Value.([]bencode.Belement)[1].Value.(int))
	}

	t.Logf("Dict val: %v", b.Value.(map[string]bencode.Belement))
}

func TestDecoderIntInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("i123"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	t.Log(err)
}

func TestDecoderStringColonInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("3abc"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	t.Log(err)
}

func TestDecoderStringLengthInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("4:abc"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	t.Log(err)
}

func TestDecoderListInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("li1e3:abc"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	t.Log(err)
}

func TestDecoderDictEndInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("d3:abc3:def"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	t.Log(err)
}

func TestDecoderDictKeyInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("di3e3:defe"))
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	t.Log(err)
}

func TestDecoderDictValueInvalid(t *testing.T) {
	_, err := bencode.Decode([]byte("d3:abce"))
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	t.Log(err)
}
