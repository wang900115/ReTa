package pkgbigcache

import (
	"bytes"
	"encoding/gob"
)

// 序列化
func serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)
	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// 反序列化
func deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return value, nil

}
