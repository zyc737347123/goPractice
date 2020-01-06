package structs

import "io"

// Rot13Reader is a reader to translate rot13 bytes
type Rot13Reader struct {
	R io.Reader
}

func (r Rot13Reader) Read(b []byte) (int, error) {
	n, err := r.R.Read(b)
	for i := 0; i < n; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = ((b[i] - 'A' + 13) % 26) + 'A'
		}
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = ((b[i] - 'a' + 13) % 26) + 'a'
		}
	}
	return n, err
}
