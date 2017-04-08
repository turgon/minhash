package minhash

import (
	"bytes"
	"testing"
)

func TestMinHash8(t *testing.T) {
	var x, b []byte
	var l int = 4

	mh := New8(l)

	b = make([]byte, 0)
	x = mh.Sum(b)

	if !bytes.Equal(x, []byte{255, 255, 255, 255}) {
		t.Error("Sum returned unexpected results", x)
	}

	mh.Write([]byte("test"))

	mh.Reset()

	b = make([]byte, 0)
	x = mh.Sum(b)

	if !bytes.Equal(x, []byte{255, 255, 255, 255}) {
		t.Error("Sum returned unexpected results", x)
	}

	if mh.BlockSize() != 1 {
		t.Error("BlockSize returned unexpected result")
	}

	if mh.Size() != l {
		t.Error("Size returned unexpected result")
	}
}

func TestFuzz(t *testing.T) {
	b := make([]byte, 0)
	if Fuzz(b) != -1 {
		t.Error("Fuzz returned unexpected result")
	}

	b = append(b, byte(0), byte(0))

	if Fuzz(b) != 0 {
		t.Error("Fuzz returned unexpected result")
	}
}