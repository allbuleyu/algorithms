package prob0415

func addStrings(num1 string, num2 string) string {
	// 进位
	var carry byte = '0'

	// 如果用byte表示,10的表示数字是58
	var tenByteNum byte = 58

	n1, n2 := len(num1)-1, len(num2)-1

	resBytes := make([]byte, 0, max(n2, n1)+2)

	for n1 >= 0 || n2 >= 0 {
		var currentByte byte
		var sum byte = 0
		if n1 >= 0 {
			sum += num1[n1]
		}

		if n2 >= 0 {
			sum += num2[n2]
		}

		sum += carry

		// 要覆盖所有情况,其实carry 也是字符串数字
		if sum >= 106 + '0' {
			currentByte = sum-tenByteNum - '0'
			carry = '1'
		}else if sum >= 144 {
			currentByte = sum - '0' - '0'
			carry = '0'
		}else if sum >= 106 {
			currentByte = sum - tenByteNum
			carry = '1'
		}else {
			currentByte = sum - '0'
			carry = '0'
		}


		resBytes = append(resBytes, currentByte)

		n1--
		n2--
	}
	if carry == '1' {
		resBytes = append(resBytes, '1')
	}

	for i, j := 0, len(resBytes)-1; i < j; i, j = i+1, j-1 {
		resBytes[i], resBytes[j] = resBytes[j], resBytes[i]
	}

	return string(resBytes)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}