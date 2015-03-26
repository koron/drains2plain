package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mattn/go-colorable"
)

func unescape(s []byte) ([]byte, error) {
	d := make([]byte, 0, len(s))
	for len(s) > 0 {
		if s[0] == '\\' && s[1] == 'x' {
			if len(s) < 4 {
				return nil, errors.New("short escape")
			}
			c, err := strconv.ParseUint(string(s[2:4]), 16, 8)
			if err != nil {
				return nil, err
			}
			d = append(d, byte(c))
			s = s[4:]
		} else {
			d = append(d, s[0])
			s = s[1:]
		}
	}
	return d, nil
}

func convert(b []byte) ([]byte, error) {
	s := bytes.SplitN(b, []byte(" "), 3)
	m, err := unescape(bytes.TrimSpace(s[2]))
	if err != nil {
		return nil, err
	}
	return []byte(m), nil
}

func main() {
	r := bufio.NewReader(os.Stdin)
	o := colorable.NewColorableStdout()
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		m, err := convert(b)
		if err != nil {
			log.Fatal(err)
		}
		o.Write(m)
	}
}
