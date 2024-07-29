package hello

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndian(t *testing.T) {
	var bs = []byte{0, 0, 0, 0}
	var data uint32 = 257

	// BigEndian encode
	binary.BigEndian.PutUint32(bs, data)
	assert.Equal(t, []byte{0, 0, 1, 1}, bs)
	// BigEndian decode
	assert.Equal(t, binary.BigEndian.Uint32(bs), data)

	// LittleEndian encode
	binary.LittleEndian.PutUint32(bs, data)
	assert.Equal(t, []byte{1, 1, 0, 0}, bs)
	// LittleEndian decode
	assert.Equal(t, binary.LittleEndian.Uint32(bs), data)
}
