package gob

import (
	"encoding/gob"
	"io"

	"github.com/pkg/errors"
)

// Decode 从给定的 reader 中读取一个 Gob 编码的值，并将其存储在 v 指向的值中。
// 'v' 参数必须是一个指针。
func Decode(reader io.Reader, v any) error {
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(v); err != nil && err != io.EOF {
		return errors.Wrap(err, "gob 解码失败")
	}
	return nil
}
