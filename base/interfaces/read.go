package interfaces

import (
	"io"
)

// StringReader is a reader
type StringReader struct {
	data string
	n    int
}

// Read will return data
func (s *StringReader) Read(p []byte) (int, error) {
	data := []byte(s.data)
	if s.n >= len(data) {
		return 0, io.EOF
	}
	data = data[s.n:]
	n := copy(p, data)
	s.n += n
	return n, nil
}

// NewReader will return a reader
func NewReader(in string) *StringReader {
	s := new(StringReader)
	s.data = in
	return s
}

type limitReader struct {
	read  io.Reader
	index int64
	limit int64
}

// Read will limit read
func (l *limitReader) Read(p []byte) (int, error) {
	readAble := l.limit - l.index
	if readAble == 0 {
		return 0, io.EOF
	}
	if len(p) < int(readAble) {
		readAble = int64(len(p))
	}
	data := make([]byte, readAble)
	n, err := l.read.Read(data)
	if err != nil {
		return 0, err
	}
	data = data[:n]
	n = copy(p, data)
	l.index += int64(n)
	return n, nil
}

// LimitReader will return a limit reader
func LimitReader(r io.Reader, n int64) io.Reader {
	lr := new(limitReader)
	lr.read = r
	lr.limit = n
	return lr
}
