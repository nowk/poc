package poc

import "testing"
import "github.com/nowk/assert"

func TestPoc(t *testing.T) {
	p := New()
	go p.Write([]byte("Hello World!"))

	b := make([]byte, 1024)
	n, err := p.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World!", string(b[:n]))
}

func TestReadBuffers(t *testing.T) {
	p := New()
	go p.Write([]byte("Hello World!"))

	for _, c := range []byte("Hello World!") {
		b := make([]byte, 1)
		n, err := p.Read(b)
		assert.Nil(t, err)
		assert.Equal(t, 1, n)
		assert.Equal(t, string(c), string(b[:n]))
	}
}
