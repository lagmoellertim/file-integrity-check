package main

import "math/rand"

type byteGen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *byteGen) RandByte() uint8 {
	if b.remaining <= 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache & 0xff
	b.cache >>= 8
	b.remaining -= 8

	return uint8(result)
}

func NewRandByteGen(seed int64) *byteGen {
	return &byteGen{src: rand.NewSource(seed)}
}
