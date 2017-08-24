package countinversions

import "testing"

func TestCountInversions(t *testing.T) {
	cases := []struct {
		array []int
		want  int
	}{
		{[]int{1, 2, 3, 5, 4}, 1},
		{[]int{2, 4, 1, 3, 5}, 3},
	}
	for _, c := range cases {
		_, got := CountAndSort(c.array, 0)

		if got != c.want {
			t.Errorf("CountAndSort got %q, want %q", got, c.want)
		}
	}
}
