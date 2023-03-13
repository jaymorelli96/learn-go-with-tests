package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("when not value for the iterations number is given, uses default of 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", repeated, expected)
		}
	})

	t.Run("run 2 times", func(t *testing.T) {
		repeated := Repeat("b", 2)
		expected := "bb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
