package util

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"os"
)

func ReadSplit(p string, split byte, fn func(data []byte) error) error {
	open, err := os.Open(p)
	if err != nil {
		return errors.WithStack(err)
	}
	defer open.Close()
	reader := bufio.NewReader(open)
	for {
		last := false
		bytes, err := reader.ReadBytes(split)
		if err == nil {
			bytes = bytes[:len(bytes)-1]
		} else {
			if err == io.EOF {
				last = true
			} else {
				return errors.WithStack(err)
			}
		}

		err = fn(bytes)
		if err != nil {
			return err
		}

		if last {
			break
		}
	}

	return nil
}
