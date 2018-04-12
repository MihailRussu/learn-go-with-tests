package arrays

import "testing"

var numbers = []int{1, 2, 3, 4, 5}

func TestSum(t *testing.T) {
	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(numbers)
	}
}
