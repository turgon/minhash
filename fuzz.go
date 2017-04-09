package minhash

func Fuzz(data []byte) int {
	if len(data) < 1 {
		return -1
	}

	sz := int(data[0])
	data = data[1:]

	mh8 := New8(sz)
	mh8.Write(data)

	mh16 := New16(sz)
	mh16.Write(data)

	return 0
}
