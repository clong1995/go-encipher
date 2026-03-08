package gob

import (
	"encoding/gob"
	"io"

	"github.com/pkg/errors"
)

// Decode reads a Gob-encoded value from the given reader and stores it in the value pointed to by v.
// The 'v' argument must be a pointer.
// Decode 解码gob
// reader 要解码的数据流,
// out 接收结果的指针
// bytes.NewBuffer(encoderData) 可以将encoderData二进制转流
// 将 r *http.Request reader,可直接解http请求流
func Decode(reader io.Reader, out any) error {
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(out); err != nil && err != io.EOF {
		return errors.Wrap(err, "gob decode failed")
	}
	return nil
}
