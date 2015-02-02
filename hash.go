package hash

func Bytes32(data []byte) uint32 {
	var v uint32 = 2166136261
	for _, c := range data {
		v = (v ^ uint32(c)) * 16777619
	}
	return v
}

func Bytes64(data []byte) uint64 {
	var v uint64 = 14695981039346656037
	for _, c := range data {
		v = (v ^ uint64(c)) * 1099511628211
	}
	return v
}

func Runes32(data []rune) uint32 {
	var v uint32 = 2166136261
	for _, c := range data {
		v = (v ^ uint32(c)) * 16777619
	}
	return v
}

func Runes64(data []rune) uint64 {
	var v uint64 = 14695981039346656037
	for _, c := range data {
		v = (v ^ uint64(c)) * 1099511628211
	}
	return v
}
