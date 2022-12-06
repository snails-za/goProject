package model

import "testing"

func BenchmarkDbConnect(b *testing.B) {
	DbConnect()
}
