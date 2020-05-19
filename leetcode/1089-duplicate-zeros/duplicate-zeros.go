package prob1089


//Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.
//
//Note that elements beyond the length of the original array are not written.
//
//Do the above modifications to the input array in place, do not return anything from your function.
//
//
//
//Example 1:
//
//Input: [1,0,2,3,0,4,5,0]
//Output: null
//Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
//Example 2:
//
//Input: [1,2,3]
//Output: null
//Explanation: After calling your function, the input array is modified to: [1,2,3]
//
//
//Note:
//
//1 <= arr.length <= 10000
//0 <= arr[i] <= 9

func duplicateZeros(arr []int)  {
	helpFunc2(arr)
}

func helpFunc1(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			if i == len(arr) - 1 {
				break
			}

			for j := len(arr)-1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}

			arr[i+1] = 0

			i++
		}
	}
}

func helpFunc2(arr []int) {
	zero := make([]int, 0)
	for i := 0; i < len(arr) - len(zero); i++ {
		if arr[i] == 0 {

			// 此处没有足够的空间复制最后一个0,所以这里直接处理掉
			if i + len(zero) + 1 == len(arr) {
				arr[len(arr)-1] = 0
				break
			}
			zero = append(zero, i)
		}
	}

	for i := len(arr) - len(zero) - 1; i >= 0 && len(zero) > 0; i-- {
		z := zero[len(zero)-1]

		if i == z {
			arr[i + len(zero)] = 0
			zero = zero[:len(zero)-1]
			arr[i + len(zero)] = 0
		}else {
			arr[i + len(zero)] = arr[i]
		}
	}
}