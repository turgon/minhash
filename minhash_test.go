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

func TestLessThan8(t *testing.T) { testLessThan(New8(8), New8(8), t) }
func TestLessThan16(t *testing.T) { testLessThan(New16(8), New16(8), t) }
func TestLessThan32(t *testing.T) { testLessThan(New32(8), New32(8), t) }
func TestLessThan64(t *testing.T) { testLessThan(New64(8), New64(8), t) }

func testLessThan(first, second MinHasher, t *testing.T) {
	first.Write([]byte("testing"))

	if !first.LessThan(second) {
		t.Error("LessThan failed")
	}

	if second.LessThan(first) {
		t.Error("LessThan failed")
	}
}

func TestSimilarity8(t *testing.T) { testSimilarity(New8(10), New8(10), t) }
func TestSimilarity16(t *testing.T) { testSimilarity(New16(10), New16(10), t) }
func TestSimilarity32(t *testing.T) { testSimilarity(New32(10), New32(10), t) }
func TestSimilarity64(t *testing.T) { testSimilarity(New64(10), New64(10), t) }

func testSimilarity(first, second MinHasher, t *testing.T) {

	first.Write([]byte("test"))
	if match := first.Similarity(second); match != 0 {
		t.Error("Similarity returned unexpected results:", match)
	}
	second.Write([]byte("test"))
	if match := first.Similarity(second); match != 10 {
		t.Error("Similarity returned unexpected results:", match)
	}
	second.Write([]byte("other thing"))
	if match := first.Similarity(second); match == 10 || match == 0 {
		t.Error("Similarity returned unexpected results:", match)
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

func BenchmarkSimilarity8(b *testing.B) { benchmarkSimilarity(New8(128), New8(128), b) }
func BenchmarkSimilarity16(b *testing.B) { benchmarkSimilarity(New16(128), New16(128), b) }
func BenchmarkSimilarity32(b *testing.B) { benchmarkSimilarity(New32(128), New32(128), b) }
func BenchmarkSimilarity64(b *testing.B) { benchmarkSimilarity(New64(128), New64(128), b) }

var answer int

func benchmarkSimilarity(first, second MinHasher, b *testing.B) {
	var r int

	b.StopTimer()
	r, _ = first.Write([]byte("testing"))
	r, _ = second.Write([]byte("testing"))
	r, _ = second.Write([]byte("testing2"))
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		r = first.Similarity(second)
	}
	answer = r
}

