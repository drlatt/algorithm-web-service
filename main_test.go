package main

import "testing"

// Benchmark tests
func BenchmarkSimpleFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFib(10)
	}
}
func BenchmarkOptimizedFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optimizedFib(10)
	}
}

func BenchmarkSimpleFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFib(20)
	}
}

func BenchmarkOptimizedFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optimizedFib(20)
	}
}
func BenchmarkSimpleFibneg20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFib(-20)
	}
}

func BenchmarkOptimizedFibneg20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optimizedFib(-20)
	}
}
func BenchmarkSimpleAckermann33(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleAckermann(3, 3)
	}
}

func BenchmarkSimpleFactorial1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFactorial(1000)
	}
}
func BenchmarkOptimizedFactorial1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optimizedFactorial(1000)
	}
}
func BenchmarkSimpleFactorial5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleFactorial(5000)
	}
}
func BenchmarkOptimizedFactorial5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optimizedFactorial(5000)
	}
}

// Unittests

func TestSimpleFib(t *testing.T) {
	data := []struct {
		n        int
		expected int64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34}, {10, 55}, {11, 89}, {12, 144}, {13, 233}, {14, 377}, {15, 610},
		{-1, 1}, {-2, -1}, {-3, 2}, {-4, -3}, {-6, -8}, {-8, -21},
	}

	for _, v := range data {
		if result := simpleFib(v.n); result != v.expected {
			t.Errorf("n: %d, Result: %d, Expected: %d", v.n, result, v.expected)
		}
	}
}

func TestOptimizedFib(t *testing.T) {
	data := []struct {
		n        int
		expected int64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34}, {10, 55}, {11, 89}, {12, 144}, {13, 233}, {14, 377}, {15, 610},
		{-1, 1}, {-2, -1}, {-3, 2}, {-4, -3}, {-6, -8}, {-8, -21},
	}

	for _, v := range data {
		if result := optimizedFib(v.n); result != v.expected {
			t.Errorf("n: %d, Result: %d, Expected: %d", v.n, result, v.expected)
		}
	}
}

func TestSimpleAckermann(t *testing.T) {
	type dataStruct struct {
		m, n int
	}

	data := make(map[int64]dataStruct)
	data[1] = dataStruct{0, 0}
	data[9] = dataStruct{2, 3}
	data[61] = dataStruct{3, 3}
	data[253] = dataStruct{3, 5}
	data[125] = dataStruct{3, 4}

	for k, v := range data {
		if result := simpleAckermann(v.m, v.n); result != k {
			t.Errorf("m, n: %d, %d, Result: %d, Expected: %d", v.m, v.n, result, k)
		}
	}
}

func TestSimpleFactorial(t *testing.T) {
	data := []struct {
		n        int
		expected uint64
	}{
		{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}, {6, 720}, {7, 5040}, {8, 40320},
	}

	for _, v := range data {
		if result := simpleFactorial(v.n); result != v.expected {
			t.Errorf("n: %d, Result: %d, Expected: %d", v.n, result, v.expected)
		}
	}
}
func TestOptimizedFactorial(t *testing.T) {
	data := []struct {
		n        int
		expected uint64
	}{
		{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}, {6, 720}, {7, 5040}, {8, 40320},
	}

	for _, v := range data {
		if result := optimizedFactorial(v.n); result != v.expected {
			t.Errorf("n: %d, Result: %d, Expected: %d", v.n, result, v.expected)
		}
	}
}
