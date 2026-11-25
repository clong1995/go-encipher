package codec

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	codecChars     = "haVdW7kwXq8fLbcNJe9yT2ROHB1FixMsj3mI0zAQKZvtu5nY6SolpCED4rgUPG"
	codecNonceSize = 4
)

var (
	errInvalidBase = errors.New("invalid codec string")

	codecDecodeTable = func() [256]byte {
		var table [256]byte
		for i := range table {
			table[i] = 0xFF
		}
		for i := range len(codecChars) {
			table[codecChars[i]] = byte(i)
		}
		return table
	}()

	codecBase = big.NewInt(62)
)

// XOR
func xor(data, key []byte) {
	if len(key) == 0 {
		return
	}
	for i := range data {
		data[i] ^= key[i%len(key)]
	}
}

// Base62 编码
func base62Encode(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	zeros := 0
	for _, b := range data {
		if b != 0 {
			break
		}
		zeros++
	}

	var n big.Int
	n.SetBytes(data)

	var rem big.Int
	out := make([]byte, 0, len(data)*2)

	for n.Sign() > 0 {
		n.QuoRem(&n, codecBase, &rem)
		out = append(out, codecChars[rem.Int64()])
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	result := make([]byte, zeros+len(out))
	for i := 0; i < zeros; i++ {
		result[i] = codecChars[0]
	}
	copy(result[zeros:], out)

	return string(result)
}

// Base62 解码
func base62Decode(str string) ([]byte, error) {
	if len(str) == 0 {
		return []byte{}, nil
	}

	zeros := 0
	for i := 0; i < len(str); i++ {
		if str[i] != codecChars[0] {
			break
		}
		zeros++
	}

	var n, digit big.Int

	for i := 0; i < len(str); i++ {
		v := codecDecodeTable[str[i]]
		if v == 0xFF {
			return nil, errInvalidBase
		}

		n.Mul(&n, codecBase)
		digit.SetInt64(int64(v))
		n.Add(&n, &digit)
	}

	bytes := n.Bytes()
	out := make([]byte, zeros+len(bytes))
	copy(out[zeros:], bytes)

	return out, nil
}

// 混淆编码
func codecEncode(data []byte) (string, error) {
	// 生成随机 nonce
	nonce := make([]byte, codecNonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// 复制原始数据
	out := make([]byte, len(data)+codecNonceSize)
	copy(out, data)

	// XOR
	xor(out[:len(data)], nonce)

	// nonce 拼接到末尾
	copy(out[len(data):], nonce)

	// Base62
	return base62Encode(out), nil
}

// 混淆解码
func codecDecode(str string) ([]byte, error) {
	data, err := base62Decode(str)
	if err != nil {
		return nil, err
	}

	if len(data) < codecNonceSize {
		return nil, errInvalidBase
	}

	// 提取 nonce
	dataLen := len(data) - codecNonceSize
	nonce := data[dataLen:]

	// XOR 还原
	xor(data[:dataLen], nonce)

	return data[:dataLen], nil
}
