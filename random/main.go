package random

import (
	"io"
	"math/rand"
)

// NewRandomDataReader creates an *io.ReadCloser returning
// random data of the specified length.
func NewRandomDataReader(length int) *RandomDataReader {
	rdr := &RandomDataReader{
		Length:      length,
		randomChars: "ABC",
	}
	return rdr
}

type RandomDataReader struct {
	Length      int
	count       int
	randomChars string
}

func (r *RandomDataReader) Read(b []byte) (int, error) {
	c := 0
	for i := 0; i < len(b) && r.count < r.Length; i++ {
		b[i] = r.randomChars[i%len(r.randomChars)]
		c++
		r.count++
	}
	if c == 0 {
		return 0, io.EOF
	}
	return c, nil
}

func (r *RandomDataReader) Close() error {
	return nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var moreLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_.,:;")

// GenerateRandomString generates a random string of the length provided as input.
func GenerateRandomString(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GenerateStrongerRandomString generates a random string of the length provided as input.
func GenerateStrongerRandomString(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = moreLetterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}