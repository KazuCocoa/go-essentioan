package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(s []byte) (int, error) {
	n, e := r.r.Read(s)

	for i := 0; i < n; i++ {
		if (s[i] >= 'A' && s[i] < 'N') || (s[i] >= 'a' && s[i] < 'n') {
			s[i] += 13
		} else if (s[i] > 'M' && s[i] <= 'Z') || (s[i] > 'm' && s[i] <= 'z') {
			s[i] -= 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
