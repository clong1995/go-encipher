package auth

// xor 函数对传入的字节切片进行原地XOR（异或）操作。
func xor(data, key []byte) {
	if len(key) == 0 {
		return
	}
	k := 0
	for i := range data {
		data[i] ^= key[k]
		k++
		if k == len(key) {
			k = 0
		}
	}
}
