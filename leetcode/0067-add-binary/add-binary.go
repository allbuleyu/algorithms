package prob0067

import "strconv"

func addBinary(a string, b string) string {

	return bitManipulation(a, b)
}

// 位操作还是有问题
func bitManipulation(a, b string) string {
	m, n := len(a), len(b)
	if n > m {
		return bitManipulation(b, a)
	}
	ba, _ := strconv.ParseInt(a, 2, 64)
	bb, _ := strconv.ParseInt(b, 2, 64)

	for bb > 0 {
		carry := (ba & bb) << 1
		ba = ba ^ bb
		bb = carry
	}

	return strconv.FormatInt(ba, 2)
}

func bitByBit(a, b string) string {
	m, n := len(a), len(b)
	if n > m {
		return bitByBit(b, a)
	}
	ans := make([]byte, 0, m+1)
	carry := 0
	for i, j := m-1, n-1; i >= 0; {
		if a[i] == '1' {
			carry++
		}
		if j >= 0 && b[j] == '1' {
			carry++
		}

		if carry % 2 != 0 {
			ans = append(ans, '1')
		}else {
			ans = append(ans, '0')
		}

		carry /= 2
		i--
		j--
	}

	if carry == 1 {
		ans = append(ans, '1')
	}

	reverse(ans)
	return string(ans)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}