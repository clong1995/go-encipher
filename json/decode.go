package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// Decode 从输入中读取下一个 JSON 编码的值，并将其存储在 v 指向的值中。
func Decode(reader io.Reader, v any) error {
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(v); err != nil && err != io.EOF {
		return errors.Wrap(err, "json 解码失败")
	}
	return nil
}
