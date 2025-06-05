package main

import "testing"

func BenchmarkFooer(b *testing.B) {
	for b.Loop() {
		Fooer(1)
	}
}
