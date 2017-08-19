package geo

import (
	"math/rand"
	"testing"
)

func checkLength(f IFeature, expected int, b *testing.B) {
	got := f.Length()
	if got != expected {
		b.Fatalf("wrong length: expected %d, got %d", got, expected)
	}
}

func rnd() float64 {
	return rand.Float64() * 180.0
}

func fill(line ILine, n int) ILine {
	for i := 0; i < n; i++ {
		line.Append(rnd(), rnd(), rnd())
	}
	return line
}

// --------------------------------------------------------------------

func Benchmark_AddLast_Line1(b *testing.B) {
	line := NewLine1()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		AddLast_Line1(line, float64(b.N))
	}

	checkLength(line, b.N, b)
}

func Benchmark_AddLast_Line2(b *testing.B) {
	line := NewLine2()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		AddLast_Line2(line, float64(b.N))
	}

	checkLength(line, b.N, b)
}

func Benchmark_AddLast_Line3(b *testing.B) {
	line := NewLine3()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		AddLast_Line3(line, float64(b.N))
	}

	checkLength(line, b.N, b)
}

func Benchmark_AddLast_Line4(b *testing.B) {
	line := NewLine4()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		AddLast_Line4(line, float64(b.N))
	}

	checkLength(line, b.N, b)
}

// --------------------------------------------------------------------

func Benchmark_RemoveFirst_Line1(b *testing.B) {
	line := NewLine1()
	line = fill(line, b.N).(*Line1)
	checkLength(line, b.N, b)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		RemoveFirst_Line1(line)
	}

	checkLength(line, 0, b)
}

func Benchmark_RemoveFirst_Line2(b *testing.B) {
	line := NewLine2()
	line = fill(line, b.N).(*Line2)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		RemoveFirst_Line2(line)
	}

	checkLength(line, 0, b)
}

func Benchmark_RemoveFirst_Line3(b *testing.B) {
	line := NewLine3()
	line = fill(line, b.N).(*Line3)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		RemoveFirst_Line3(line)
	}

	checkLength(line, 0, b)
}

func Benchmark_RemoveFirst_Line4(b *testing.B) {
	line := NewLine4()
	line = fill(line, b.N).(*Line4)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		RemoveFirst_Line4(line)
	}

	checkLength(line, 0, b)
}
