package unionfind

import "testing"
import "fmt"

func TestUnionFind(t *testing.T) {
	cases := []struct {
		size	int
		unions	[]int
	}{
		{7, []int{1,3,4,5}},
	}

	for _, c := range cases {
		heap := new(Union)
		heap.Alloc(c.size)

		for _, el := range c.unions {
			heap.Union(1, el)
		}


		for _, u := range heap.p {
			fmt.Print(u, " ")
		}

		fmt.Println("\n")
		fmt.Println(heap.ConnComps())
		heap.Union(6, 3)
		fmt.Println(heap.ConnComps())
	}
}
