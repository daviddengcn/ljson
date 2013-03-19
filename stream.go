// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ljson

import (
	"io"
)

// A Decoder reads and decodes JSON objects from an input stream.
type Decoder struct {
	r    io.Reader
	buf  []byte
	d    decodeState
	scan scanner
	err  error
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the JSON values requested.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode reads the next JSON-encoded value from its
// input and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about
// the conversion of JSON into a Go value.
func (dec *Decoder) Decode(v interface{}) error {
	if dec.err != nil {
		return dec.err
	}

	n, err := dec.readValue()
	if err != nil {
		return err
	}

	// Don't save err from unmarshal into dec.err:
	// the connection is still usable since we read a complete JSON
	// object from it before the error happened.
	dec.d.init(dec.buf[0:n])
	err = dec.d.unmarshal(v)

	// Slide rest of data down.
	rest := copy(dec.buf, dec.buf[n:])
	dec.buf = dec.buf[0:rest]

	return err
}

// readValue reads a JSON value into dec.buf.
// It returns the length of the encoding.
func (dec *Decoder) readValue() (int, error) {
	dec.scan.reset()

	scanp := 0
	var err error
Input:
	for {
		// Look in the buffer for a new value.
		for i, c := range dec.buf[scanp:] {
			dec.scan.bytes++
			v := dec.scan.step(&dec.scan, int(c))
			if v == scanEnd {
				scanp += i
				break Input
			}
			// scanEnd is delayed one byte.
			// We might block trying to get that byte from src,
			// so instead invent a space byte.
			if v == scanEndObject && dec.scan.step(&dec.scan, ' ') == scanEnd {
				scanp += i + 1
				break Input
			}
			if v == scanError {
				dec.err = dec.scan.err
				return 0, dec.scan.err
			}
		}
		scanp = len(dec.buf)

		// Did the last read have an error?
		// Delayed until now to allow buffer scan.
		if err != nil {
			if err == io.EOF {
				if dec.scan.step(&dec.scan, ' ') == scanEnd {
					break Input
				}
				if nonSpace(dec.buf) {
					err = io.ErrUnexpectedEOF
				}
			}
			dec.err = err
			return 0, err
		}

		// Make room to read more into the buffer.
		const minRead = 512
		if cap(dec.buf)-len(dec.buf) < minRead {
			newBuf := make([]byte, len(dec.buf), 2*cap(dec.buf)+minRead)
			copy(newBuf, dec.buf)
			dec.buf = newBuf
		}

		// Read.  Delay error for next iteration (after scan).
		var n int
		n, err = dec.r.Read(dec.buf[len(dec.buf):cap(dec.buf)])
		dec.buf = dec.buf[0 : len(dec.buf)+n]
	}
	return scanp, nil
}

func nonSpace(b []byte) bool {
	for _, c := range b {
		if !isSpace(rune(c)) {
			return true
		}
	}
	return false
}
