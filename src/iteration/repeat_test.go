package iteration

import (
	"testing"
	"fmt"
	"strings"
)

func TestRepeat(t *testing.T) {
	t.Run("generic test", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("test with standard library", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := strings.Repeat("a", 5)

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Repeat("a", 5)
	}
}
