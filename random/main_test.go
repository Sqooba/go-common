package random

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestNewRandomDataReader(t *testing.T) {

	length := rand.Intn(100) + 100
	rdr := NewRandomDataReader(length)

	read, err := ioutil.ReadAll(rdr)
	assert.Nil(t, err)
	assert.Equal(t, length, len(read))
}
