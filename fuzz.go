package minhash

import (
	"hash"
)

func test(mh hash.Hash, data []byte) bool {
	b := make([]byte, 0)
	b = mh.Sum(b)

	mh.Write(data)

	a := make([]byte, 0)
	a = mh.Sum(a)

	// return a.LessThan(b)
	return true
}

func Fuzz(data []byte) int {
	if len(data) < 1 {
		return -1
	}

	sz := int(data[0])
	data = data[1:]

	mh8 := New8(sz)
	if !test(mh8, data) {
		panic("fail")
	}

	mh16 := New16(sz)
	test(mh16, data)

	mh32 := New32(sz)
	test(mh32, data)

	mh64 := New64(sz)
	test(mh64, data)

	return 0
}
