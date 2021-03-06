package minhash

import (
	"hash"
	"hash/fnv"
)

type MinHash8 []uint8
type MinHash16 []uint16
type MinHash32 []uint32
type MinHash64 []uint64

type MinHasher interface {
	hash.Hash
	Similarity(MinHasher) int
	LessThan(MinHasher) bool
}

func New8(sz int) MinHash8 {
	n := make(MinHash8, sz)
	n.Reset()
	return n
}

func New16(sz int) MinHash16 {
	n := make(MinHash16, sz)
	n.Reset()
	return n
}

func New32(sz int) MinHash32 {
	n := make(MinHash32, sz)
	n.Reset()
	return n
}

func New64(sz int) MinHash64 {
	n := make(MinHash64, sz)
	n.Reset()
	return n
}

func (mh MinHash8) Reset() {
	for i := 0; i < len(mh); i++ {
		mh[i] = 255
	}
}

func (mh MinHash16) Reset() {
	for i := 0; i < len(mh); i++ {
		mh[i] = 65535
	}
}

func (mh MinHash32) Reset() {
	for i := 0; i < len(mh); i++ {
		mh[i] = 4294967295
	}
}

func (mh MinHash64) Reset() {
	for i := 0; i < len(mh); i++ {
		mh[i] = 18446744073709551615
	}
}

func (mh MinHash8) Sum(in []byte) []byte {
	for i := 0; i < len(mh); i++ {
		in = append(in, byte(mh[i]))
	}
	return in
}

func (mh MinHash16) Sum(in []byte) []byte {
	for i := 0; i < len(mh); i++ {
		in = append(in, byte(mh[i] >> 8), byte(mh[i]))
	}
	return in
}

func (mh MinHash32) Sum(in []byte) []byte {
	for i := 0; i < len(mh); i++ {
		in = append(in, byte(mh[i] >> 24), byte(mh[i] >> 16), byte(mh[i] >> 8), byte(mh[i]))
	}
	return in
}

func (mh MinHash64) Sum(in []byte) []byte {
	for i := 0; i < len(mh); i++ {
		in = append(in, byte(mh[i] >> 56), byte(mh[i] >> 48), byte(mh[i] >> 40), byte(mh[i] >> 32), byte(mh[i] >> 24), byte(mh[i] >> 16), byte(mh[i] >> 8), byte(mh[i]))
	}
	return in
}

func (mh MinHash8) Write(data []byte) (int, error) {
	h := fnv.New32a()
	h.Write(data)

	v1 := uint8(h.Sum32())
	v2 := uint8(h.Sum32() >> 8)

	for i := 0; i < len(mh); i++ {
		x := v1 + uint8(i)*v2
		if x < mh[i] {
			mh[i] = x
		}
	}

	return len(data), nil
}

func (mh MinHash16) Write(data []byte) (int, error) {
	h := fnv.New32a()
	h.Write(data)

	v1 := uint16(h.Sum32())
	v2 := uint16(h.Sum32() >> 16)

	for i := 0; i < len(mh); i++ {
		x := v1 + uint16(i)*v2
		if x < mh[i] {
			mh[i] = x
		}
	}

	return len(data), nil
}

func (mh MinHash32) Write(data []byte) (int, error) {
	h := fnv.New64a()
	h.Write(data)

	v1 := uint32(h.Sum64())
	v2 := uint32(h.Sum64() >> 32)

	for i := 0; i < len(mh); i++ {
		x := v1 + uint32(i)*v2
		if x < mh[i] {
			mh[i] = x
		}
	}

	return len(data), nil
}

func (mh MinHash64) Write(data []byte) (int, error) {
	h := fnv.New64a()
	h.Write(data)

	v1 := h.Sum64()

	h.Write(data)

	v2 := h.Sum64()

	for i := 0; i < len(mh); i++ {
		x := v1 + uint64(i)*v2
		if x < mh[i] {
			mh[i] = x
		}
	}

	return len(data), nil
}

func (mh MinHash8) BlockSize() int { return 1 }
func (mh MinHash16) BlockSize() int { return 1 }
func (mh MinHash32) BlockSize() int { return 1 }
func (mh MinHash64) BlockSize() int { return 1 }

func (mh MinHash8) Size() int { return len(mh) }
func (mh MinHash16) Size() int { return len(mh) }
func (mh MinHash32) Size() int { return len(mh) }
func (mh MinHash64) Size() int { return len(mh) }

func (mh MinHash8) LessThan(o MinHasher) bool {
	for i := 0; i < o.Size(); i++ {
		if mh[i] > o.(MinHash8)[i] {
			return false
		}
	}
	return true
}

func (mh MinHash16) LessThan(o MinHasher) bool {
	for i := 0; i < o.Size(); i++ {
		if mh[i] > o.(MinHash16)[i] {
			return false
		}
	}
	return true
}

func (mh MinHash32) LessThan(o MinHasher) bool {
	for i := 0; i < o.Size(); i++ {
		if mh[i] > o.(MinHash32)[i] {
			return false
		}
	}
	return true
}

func (mh MinHash64) LessThan(o MinHasher) bool {
	for i := 0; i < o.Size(); i++ {
		if mh[i] > o.(MinHash64)[i] {
			return false
		}
	}
	return true
}

func (mh MinHash8) Similarity(o MinHasher) int {
	x := 0
	var p MinHash8 = o.(MinHash8)
	for i := 0; i < p.Size(); i++ {
		if mh[i] == p[i] {
			x += 1
		}
	}
	return x
}

func (mh MinHash16) Similarity(o MinHasher) int {
	x := 0
	var p MinHash16 = o.(MinHash16)
	for i := 0; i < p.Size(); i++ {
		if mh[i] == p[i] {
			x += 1
		}
	}
	return x
}

func (mh MinHash32) Similarity(o MinHasher) int {
	x := 0
	var p MinHash32 = o.(MinHash32)
	for i := 0; i < p.Size(); i++ {
		if mh[i] == p[i] {
			x += 1
		}
	}
	return x
}

func (mh MinHash64) Similarity(o MinHasher) int {
	x := 0
	var p MinHash64 = o.(MinHash64)
	for i := 0; i < p.Size(); i++ {
		if mh[i] == p[i] {
			x += 1
		}
	}
	return x
}
