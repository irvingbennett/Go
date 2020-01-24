// Package circular manages a circular buffer
package circular

import (
	"fmt"
)

//Buffer is a circular ring buffer.
type Buffer struct {
	read, used int
	data       []byte
}

/*NewBuffer creates a new ring buffer of a certain size.*/
func NewBuffer(size int) *Buffer {
	var buff = Buffer{0, 0, make([]byte, size)}
	return &buff
}

/*ReadByte reads the oldest byte in the buffer,
cant read an empty buffer.*/
func (b *Buffer) ReadByte() (byte, error) {
	if b.used == 0 {
		return 0, fmt.Errorf("Buffer is empty")
	}
	read := b.data[b.read]
	b.read = (b.read + 1) % len(b.data)
	b.used--
	return read, nil
}

/*WriteByte writes to the buffer and woun't write to a full buffer.*/
func (b *Buffer) WriteByte(c byte) error {
	if b.used == len(b.data) {
		return fmt.Errorf("Buffer is full")
	}
	write := (b.read + b.used) % len(b.data)
	b.data[write] = c
	b.used++
	return nil
}

/*Reset clears the buffer of all data. it really doen't not clear the buffer but readbyte will show the Buffer as empty*/
func (b *Buffer) Reset() {
	b.used = 0
}

/*Overwrite writes to buffer even if it's full*/
func (b *Buffer) Overwrite(c byte) {
	if b.used == len(b.data) {
		b.ReadByte()
	}
	b.WriteByte(c)
}
