package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// Encode writes the JSON encoding of v to the stream, followed by a newline character.
// It disables HTML escaping.
// Encode 编码数据为json
// in 要编码的数据的指针
// writer 接收结果的指针
// data := writer.Bytes() 可以将流转 []byte
// 将 w http.ResponseWriter 传入writer，可直接写入http响应流
func Encode(in any, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(in); err != nil {
		return errors.Wrap(err, "json encode failed")
	}
	return nil
}
