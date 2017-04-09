package minhash

import (
	"hash"
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

func TestMinHash16(t *testing.T) {
	var x, b []byte
	var l int = 2

	mh := New16(l)

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

func TestMinHash32(t *testing.T) {
	var x, b []byte
	var l int = 1

	mh := New32(l)

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

var resulth hash.Hash

func BenchmarkNewMinHash8(b *testing.B) { benchmarkNewMinHash(New8, b) }
func BenchmarkNewMinHash16(b *testing.B) { benchmarkNewMinHash(New16, b) }
func BenchmarkNewMinHash32(b *testing.B) { benchmarkNewMinHash(New32, b) }

func benchmarkNewMinHash(mhf newFn, b *testing.B) {
	var mh hash.Hash
	for n := 0; n < b.N; n++ {
		mh = mhf(128)
	}
	resulth = mh
}

func BenchmarkMinHash8(b *testing.B) { benchmarkMinHash(New8(128), b) }
func BenchmarkMinHash16(b *testing.B) { benchmarkMinHash(New16(128), b) }
func BenchmarkMinHash32(b *testing.B) { benchmarkMinHash(New32(128), b) }

var result int

func benchmarkMinHash(mh hash.Hash, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		mh.Reset()
		r, _ = mh.Write([]byte("testing"))
	}
	result = r
}

