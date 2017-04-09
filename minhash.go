package minhash

import (
	"hash"
	"hash/fnv"
)

type newFn func(int) hash.Hash

type MinHash8 []uint8
type MinHash16 []uint16
type MinHash32 []uint32

func New8(sz int) hash.Hash {
	n := make(MinHash8, sz)
	for i := 0; i < sz; i++ {
		n[i] = 255
	}
	return &n
}

func New16(sz int) hash.Hash {
	n := make(MinHash16, sz)
	for i := 0; i < sz; i++ {
		n[i] = 65535
	}
	return &n
}

func New32(sz int) hash.Hash {
	n := make(MinHash32, sz)
	for i := 0; i < sz; i++ {
		n[i] = 4294967295
	}
	return &n
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

func (mh MinHash8) Write(data []byte) (int, error) {
	h := fnv.New32a()
	h.Write(data)
	s := make([]byte, 0)
	s = h.Sum(s)

	v1 := uint8(s[0])
	v2 := uint8(s[1])

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
	s := make([]byte, 0)
	s = h.Sum(s)

	v1 := uint16(s[0] << 8) + uint16(s[1])
	v2 := uint16(s[2] << 8) + uint16(s[3])

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
	s := h.Sum64()

	v1 := uint32(s)
	v2 := uint32(s >> 32)

	for i := 0; i < len(mh); i++ {
		x := v1 + uint32(i)*v2
		if x < mh[i] {
			mh[i] = x
		}
	}

	return len(data), nil
}

func (mh MinHash8) BlockSize() int { return 1 }
func (mh MinHash16) BlockSize() int { return 1 }
func (mh MinHash32) BlockSize() int { return 1 }

func (mh MinHash8) Size() int { return len(mh) }
func (mh MinHash16) Size() int { return len(mh) }
func (mh MinHash32) Size() int { return len(mh) }
