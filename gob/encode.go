package gob

import (
	"encoding/gob"
	"io"

	"github.com/pkg/errors"
)

// Encode 将 v 的 Gob 表示形式写入给定的 writer。
// 'v' 参数必须是指向待编码数据的指针。
func Encode(v any, writer io.Writer) error {
	encoder := gob.NewEncoder(writer)
	if err := encoder.Encode(v); err != nil {
		return errors.Wrap(err, "gob 编码失败")
	}
	return nil
}
