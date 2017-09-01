package geo

import (
	"testing"
)

const Size int = 10000

func checkLength(f Line, expected int, b *testing.B) {
	got := f.Length()
	if got != expected {
		b.Fatalf("wrong length: expected %d, got %d", expected, got)
	}
}

// --------------------------------------------------------------------

func Benchmark_AddLast_Line1(b *testing.B) {
	line := NewLine1()

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.AddLast(1.0, 2.0, 3.0)
	}
	b.StopTimer()

	checkLength(line, b.N, b)
}

func Benchmark_AddLast_Line2(b *testing.B) {
	line := NewLine2()

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.AddLast(1.0, 2.0, 3.0)
	}
	b.StopTimer()

	checkLength(line, b.N, b)
}

func Benchmark_AddLast_Line3(b *testing.B) {
	line := NewLine3()

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.AddLast(1.0, 2.0, 3.0)
	}
	b.StopTimer()

	checkLength(line, b.N, b)
}

func Benchmark_AddLast_Line4(b *testing.B) {
	line := NewLine4()

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.AddLast(1.0, 2.0, 3.0)
	}
	b.StopTimer()

	checkLength(line, b.N, b)
}

// --------------------------------------------------------------------

func Benchmark_RemoveFirst_Line1(b *testing.B) {
	line := NewLine1()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.RemoveFirst()
	}
	b.StopTimer()

	checkLength(line, 0, b)
}

func Benchmark_RemoveFirst_Line2(b *testing.B) {
	line := NewLine2()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.RemoveFirst()
	}
	b.StopTimer()

	checkLength(line, 0, b)
}

func Benchmark_RemoveFirst_Line3(b *testing.B) {
	line := NewLine3()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.RemoveFirst()
	}
	b.StopTimer()

	checkLength(line, 0, b)
}

func Benchmark_RemoveFirst_Line4(b *testing.B) {
	line := NewLine4()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.RemoveFirst()
	}
	b.StopTimer()

	checkLength(line, 0, b)
}

// --------------------------------------------------------------------

func Benchmark_Reproject_Line1(b *testing.B) {
	line := NewLine1()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.Reproject()
	}
	b.StopTimer()

	checkLength(line, b.N, b)
}

func Benchmark_Reproject_Line2(b *testing.B) {
	line := NewLine2()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.Reproject()
	}
	b.StopTimer()
}

func Benchmark_Reproject_Line3(b *testing.B) {
	line := NewLine3()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.Reproject()
	}
	b.StopTimer()
}

func Benchmark_Reproject_Line4(b *testing.B) {
	line := NewLine4()
	line.Resize(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		line.Reproject()
	}
	b.StopTimer()
}
