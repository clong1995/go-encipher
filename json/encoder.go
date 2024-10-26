package json

import (
	"encoding/json"
	"log"
)

// Encoder 编码数据为json
// in 要编码的数据的指针
// data 接收结果的指针
func Encoder(in any, data *[]byte) (err error) {
	d, err := json.Marshal(in)
	if err != nil {
		log.Println(err)
		return
	}
	*data = d
	return
}
