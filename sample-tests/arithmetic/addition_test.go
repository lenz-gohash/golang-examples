package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		x, y int
		want int
	}{
		"basic":    {x: 1, y: 2, want: 3},
		"zero":     {x: 0, y: 0, want: 0},
		"overflow": {x: 1<<31 - 1, y: 1, want: 1<<31 - 1}, // 2147483647 + 1 = -2147483648 (overflow)
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Add(tc.x, tc.y)
			if tc.want != got {
				t.Fatalf("expected: %d, got: %d", tc.want, got)
			}
		})
	}
}
