package poc

import "io"
import . "gopkg.in/nowk/go-calm.v1"

// Poc is a struct around a single channel that gets written to and read from
// synchronously. Similar to io.Pipe but on a single channel.
type Poc struct {
	ch  chan []byte
	buf []byte
}

// New returns a new Poc
func New() (p *Poc) {
	p = &Poc{
		ch: make(chan []byte),
	}

	return
}

// Close closes the channel
func (p *Poc) Close() (err error) {
	close(p.ch)
	return
}

// Write implements io.Writer. It writes incoming bytes to the channel
func (p *Poc) Write(b []byte) (int, error) {
	n := len(b)
	if err := Calm(func() {
		p.ch <- b
	}); err != nil {
		return 0, err
	}

	return n, nil
}

// Read implements io.Reader. It reads any bytes written to the channel. It will
// buffer any excess bytes that cannot be read into d
func (p *Poc) Read(d []byte) (int, error) {
	// if there is buffered bytes, read till empty before taking in from the
	// channel
	j := len(p.buf)
	if j > 0 {
		i := copy(d, p.buf)
		if i < j {
			p.buf = p.buf[i:]
			return i, nil
		}

		p.buf = nil
		return i, nil
	}

	b := <-p.ch
	n := len(b)
	if n == 0 {
		return 0, io.EOF
	}

	// attempt to copy recieved bytes to d, if the bytes can't not all fit into d
	// save it to buf for the next read
	i := copy(d, b)
	if i < n {
		p.buf = b[i:]
	}

	return i, nil
}
