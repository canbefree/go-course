package test

import "testing"

func TestXXX(t *testing.T) {
	t.Log("hello world")
}

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func TestArrayAppend(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	slice := arr[1:3]

	slice = append(slice, 1)

	except := []int{2, 3, 1}

	for i := range slice {
		if slice[i] != except[i] {
			t.Errorf("slice not equal except: %v %v", slice, except)
		}
	}
}
