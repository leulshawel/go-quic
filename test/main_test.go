package test

import (
	"go-quic/quic"
	"testing"
)

func BenchmarkGenerator(b *testing.B) {
	quic.GetConnectionById(8)
}
