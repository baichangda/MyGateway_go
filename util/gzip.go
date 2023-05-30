package util

import (
	"bytes"
	gzip2 "compress/gzip"
	"io"
)

func Gzip(b []byte) ([]byte, error) {
	reader, err := gzip2.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func UnGzip(b []byte) ([]byte, error) {
	res := bytes.Buffer{}
	writer := gzip2.NewWriter(&res)
	defer writer.Close()
	_, err := writer.Write(b)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}
