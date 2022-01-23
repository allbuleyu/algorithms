package kit

func findKthLargest(nums []int, k int) int {
	q := NewQ(k)
	for i := 0; i < len(nums); i++ {
		q.insert(nums[i])
	}

	return q.q[1]
}

type pq struct {
	n int
	l int
	q []int
}

func NewQ(n int) *pq {
	return &pq{n: n + 1, l: 1, q: make([]int, n+1)}
}

func (q *pq) insert(x int) {
	if q.l != q.n {
		q.q[q.l] = x
		q.swim(q.l)
		q.l++
		return
	}

	if q.q[1] < x {
		q.q[1] = x
		q.sink(1)
	}
}

func (q *pq) del() int {
	if q.isEmpty() {
		return 0
	}

	x := q.q[1]
	q.q[1] = q.q[q.l-1]
	q.l--
	q.sink(1)

	return x
}

func (q *pq) less(i, j int) bool {
	return q.q[i] > q.q[j]
}

func (q *pq) swap(i, j int) {
	q.q[i], q.q[j] = q.q[j], q.q[i]
}

func (q *pq) isEmpty() bool {
	return q.l == 1
}

func (q *pq) swim(k int) {
	for k > 1 && q.less(k/2, k) {
		q.swap(k/2, k)
		k = k / 2
	}
}

func (q *pq) sink(k int) {
	for k*2 < q.l {
		j := k * 2
		if j < q.l-1 && q.less(j, j+1) {
			j++
		}

		if !q.less(k, j) {
			break
		}

		q.swap(k, j)
		k = j
	}
}
