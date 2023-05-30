package util

import (
	"bytes"
	gzip2 "compress/gzip"
	"github.com/pkg/errors"
	"io"
)

func Gzip(b []byte) ([]byte, error) {
	reader, err := gzip2.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer reader.Close()
	all, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return all, nil
}

func UnGzip(b []byte) ([]byte, error) {
	res := bytes.Buffer{}
	writer := gzip2.NewWriter(&res)
	defer writer.Close()
	_, err := writer.Write(b)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res.Bytes(), nil
}
