package minhash

func Fuzz(data []byte) int {
	if len(data) < 2 {
		return -1
	}

	sz := int(data[0])
	data = data[1:]

	mh := New8(sz)
	mh.Write(data)

	return 0
}
