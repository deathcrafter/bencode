package bencode_test

import (
	"testing"

	"github.com/deathcrafter/bencode"
)

func TestBelement(t *testing.T) {
	b := bencode.Belement{Type: bencode.TypeInt, Value: 123}

	if b.Type != bencode.TypeInt {
		t.Fatalf("Expected type %d, got %d", bencode.TypeInt, b.Type)
	}

	if v, err := b.GetInt(); v != 123 || err != nil {
		t.Fatalf("Expected value %d, got %d", 123, v)
	}

	t.Log(b)
}
