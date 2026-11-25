package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// Encode 将 v 的 JSON 编码写入流中，并后跟一个换行符。
// 此函数会禁用 HTML 转义。
func Encode(v any, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
