package minhash

import (
	"hash"
	"hash/fnv"
)

type MinHash8 []uint8

func New8(sz int) hash.Hash {
	n := make(MinHash8, sz)
	for i := 0; i < sz; i++ {
		n[i] = 255
	}
	return &n
}

func (mh MinHash8) Reset() {
	for i := 0; i < len(mh); i++ {
		mh[i] = 255
	}
}

func (mh MinHash8) Sum(in []byte) []byte {
	for i := 0; i < len(mh); i++ {
		in = append(in, byte(mh[i]))
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

func (mh MinHash8) BlockSize() int { return 1 }

func (mh MinHash8) Size() int { return len(mh) }
