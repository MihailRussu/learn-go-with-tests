package hello_world

import "testing"

const name = "Batman"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(name, "")
		want := "Hello, " + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello(name, "Spanish")
		want := "Hola, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello(name, "French")
		want := "Bonjour, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello(name, "Russian")
		want := "Привет, " + name
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkHello(b *testing.B) {
	var benchmarks = [4]string{"English", "Spanish", "French", "Russian"}

	for i := 0; i < len(benchmarks); i++ {
		b.Run("benchmarking "+benchmarks[i], func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Hello(name, benchmarks[i])
			}
		})
	}

}
