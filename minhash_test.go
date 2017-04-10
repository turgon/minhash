package minhash

import (
	"hash"
	"testing"
)

func TestMinHash8(t *testing.T) { testMinHash(New8(4), 4, t) }
func TestMinHash16(t *testing.T) { testMinHash(New16(2), 2, t) }
func TestMinHash32(t *testing.T) { testMinHash(New32(1), 1, t) }
func TestMinHash64(t *testing.T) { testMinHash(New64(1), 1, t) }

func testMinHash(mh hash.Hash, l int, t *testing.T) {
	var x, b []byte

	b = make([]byte, 0)
	x = mh.Sum(b)

	for i := 0; i < len(x); i++ {
		if x[i] != 255 {
			t.Error("Sum returned unexpected results", x)
		}
	}

	mh.Write([]byte("test"))

	mh.Reset()

	b = make([]byte, 0)
	x = mh.Sum(b)

	for i := 0; i < len(x); i++ {
		if x[i] != 255 {
			t.Error("Sum returned unexpected results", x)
		}
	}

	if mh.BlockSize() != 1 {
		t.Error("BlockSize returned unexpected result")
	}

	if mh.Size() != l {
		t.Error("Size returned unexpected result")
	}
}

func TestLess(t *testing.T) {
	l := 8
	mh8 := New8(l)

	mh8.Write([]byte("testing"))

	if !mh8.LessThan(New8(l)) {
		t.Error("Less failed")
	}

	if New8(l).LessThan(mh8) {
		t.Error("Less failed")
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

func BenchmarkNewMinHash8(b *testing.B) {
	var mh hash.Hash
	for n := 0; n < b.N; n++ {
		mh = New8(128)
	}
	resulth = mh
}

func BenchmarkNewMinHash16(b *testing.B) {
	var mh hash.Hash
	for n := 0; n < b.N; n++ {
		mh = New16(128)
	}
	resulth = mh
}

func BenchmarkNewMinHash32(b *testing.B) {
	var mh hash.Hash
	for n := 0; n < b.N; n++ {
		mh = New32(128)
	}
	resulth = mh
}

func BenchmarkNewMinHash64(b *testing.B) {
	var mh hash.Hash
	for n := 0; n < b.N; n++ {
		mh = New64(128)
	}
	resulth = mh
}

func BenchmarkMinHash8(b *testing.B) { benchmarkMinHash(New8(128), b) }
func BenchmarkMinHash16(b *testing.B) { benchmarkMinHash(New16(128), b) }
func BenchmarkMinHash32(b *testing.B) { benchmarkMinHash(New32(128), b) }
func BenchmarkMinHash64(b *testing.B) { benchmarkMinHash(New64(128), b) }

var result int

func benchmarkMinHash(mh hash.Hash, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		mh.Reset()
		r, _ = mh.Write([]byte("testing"))
	}
	result = r
}

