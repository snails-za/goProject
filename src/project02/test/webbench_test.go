package test

import (
	"project02/model"
	"testing"
)

func BenchmarkDbConnect(b *testing.B) {
	model.DbConnect()
}
