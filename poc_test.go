package poc

import "testing"
import "gopkg.in/nowk/assert.v2"

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
	go func() {
		p.Write([]byte("Hello World!"))
		p.Write([]byte("Wat!"))
	}()

	for _, c := range []byte("Hello World!") {
		b := make([]byte, 1)
		n, err := p.Read(b)
		assert.Nil(t, err)
		assert.Equal(t, 1, n)
		assert.Equal(t, string(c), string(b[:n]))
	}

	b := make([]byte, 1024)
	n, err := p.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "Wat!", string(b[:n]))
}
