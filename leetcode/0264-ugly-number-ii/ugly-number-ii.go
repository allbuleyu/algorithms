package prob0264


func nthUglyNumber(n int) int {
	var un, ugly int

	hash := make(map[int]bool)

	ugly = 1
	un = 1
	for un != n {
		if _, ok := hash[ugly]; !ok {
			hash[ugly] = true

			un++
		}


	}
}