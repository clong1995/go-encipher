package gob

import (
	"encoding/gob"
	"io"

	"github.com/pkg/errors"
)

// Encode writes the Gob representation of v to the given writer.
// The 'v' argument must be a pointer to the data being encoded.
// Encode 编码gob
// in 要编码的数据的指针
// writer 接收结果的指针
// data := writer.Bytes() 可以将流转 []byte
// 将 w http.ResponseWriter 传入writer，可直接写入http响应流
func Encode(in any, writer io.Writer) error {
	encoder := gob.NewEncoder(writer)
	if err := encoder.Encode(in); err != nil {
		return errors.Wrap(err, "gob encode failed")
	}
	return nil
}
