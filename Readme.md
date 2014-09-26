# poc

[![Build Status](https://travis-ci.org/nowk/poc.svg?branch=master)](https://travis-ci.org/nowk/poc)
[![GoDoc](https://godoc.org/github.com/nowk/poc?status.svg)](http://godoc.org/github.com/nowk/poc)

Pipe on channel

## Example

    p := poc.New()
    go p.Write([]byte("Hello World!"))

    b := make([]byte, 1024)
    n, err := p.Read(b)
    if err != nil {
      // handle error
    }

    // string(b[:n]) == "Hello World!"

## License

MIT