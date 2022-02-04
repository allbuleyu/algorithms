package kit

// 压缩路径的加权 union-find 算法与数据结构
type weightUinonFind struct {
	n, count   int
	id, weight []int
}

func NewWUF(n int) *weightUinonFind {
	id := make([]int, n)
	weight := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		weight[i] = 1
	}

	return &weightUinonFind{n, n, id, weight}
}

func (wuf *weightUinonFind) Find(p int) int {
	x := p
	for x != wuf.id[x] {
		x = wuf.id[x]
	}
	// 压缩路径
	wuf.id[p] = x

	return x
}

func (wuf *weightUinonFind) Uinon(p, q int) {
	proot, qroot := wuf.Find(p), wuf.Find(q)
	if proot == qroot {
		return
	}

	if wuf.weight[proot] < wuf.weight[qroot] {
		wuf.id[proot] = qroot
		wuf.weight[qroot] += wuf.weight[proot]
	} else {
		wuf.id[qroot] = proot
		wuf.weight[proot] += wuf.weight[qroot]
	}
	wuf.count--
}

func (wuf *weightUinonFind) Connected(p, q int) bool {
	return wuf.Find(p) == wuf.Find(q)
}

func (wuf *weightUinonFind) Count() int {
	return wuf.count
}

// 压缩路径的 union-find 算法与数据结构
type unionFind struct {
	n  int
	id []int
}

func NewUF(n int) *unionFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &unionFind{n, id}
}

func (uf *unionFind) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *unionFind) union(p, q int) {
	proot, qroot := uf.find(p), uf.find(q)
	if proot == qroot {
		return
	}

	uf.id[proot] = qroot
}

// 找到p根节点
func (uf *unionFind) find(p int) int {
	x := p
	for x != uf.id[x] {
		x = uf.id[x]
	}

	// 压缩路径
	uf.id[p] = x
	return x
}
