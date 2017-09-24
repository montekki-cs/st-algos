package unionfind

type Union struct {
	p	[]int
	rank	[]int
}

func (u *Union) Alloc(size int) {
	u.p = make([]int, size)
	u.rank = make([]int, size)

	for i, _ := range u.p {
		u.p[i] = i
	}

	for i, _ := range u.rank {
		u.rank[i] = 0
	}
}

func (u *Union) Find(x int) int {
	if u.p[x] == x {
		return x
	} else {
		return u.Find(u.p[x])
	}
}

func (u *Union) Union(x, y int) {
	if u.Find(x) == u.Find(y) {
		return
	}

	if u.rank[u.Find(x)] < u.rank[u.Find(y)] {
		u.p[u.Find(x)] = u.Find(y);
	} else {
		u.p[u.Find(y)] = u.Find(x);
	}

	if u.rank[u.Find(x)] == u.rank[u.Find(y)] {
		u.rank[x]++;
	}
}

func (u *Union) ConnComps() (res map[int][]int) {
	res = make(map[int][]int)

	for i, _ := range u.p {
		find := u.Find(i)
		_, ok := res[find]

		if !ok {
			res[find] = make([]int, 0, 256)
		}

		res[find] = append(res[find], i)
	}

	return res
}
