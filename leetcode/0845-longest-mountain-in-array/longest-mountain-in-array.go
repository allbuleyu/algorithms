package prob0845


func longestMountain(A []int) int {
	return twoPointers(A)
}

func twoPointers(a []int) int {
	var ans, ascMax, descMax int
	for i := 1; i < len(a); {
		if a[i] == a[i-1] {
			i++
			continue
		}

		ascMax = 1
		for i < len(a) && a[i-1] < a[i] {
			ascMax++
			i++
		}

		//if ascMax >= 3 && ascMax > ans {
		//	ans = ascMax
		//}

		descMax = 1
		for i < len(a) && a[i-1] > a[i] {
			descMax++
			i++
		}

		mm := descMax + ascMax - 1
		if ascMax >= 2 && descMax >= 2 && mm >= 3 && mm > ans {
			ans = mm
		}
	}

	return ans
}

//public int longestMountain(int[] A) {
//        int N = A.length;
//        int ans = 0, base = 0;
//        while (base < N) {
//            int end = base;
//            // if base is a left-boundary
//            if (end + 1 < N && A[end] < A[end + 1]) {
//                // set end to the peak of this potential mountain
//                while (end + 1 < N && A[end] < A[end + 1]) end++;
//
//                // if end is really a peak..
//                if (end + 1 < N && A[end] > A[end + 1]) {
//                    // set end to the right-boundary of mountain
//                    while (end + 1 < N && A[end] > A[end + 1]) end++;
//                    // record candidate answer
//                    ans = Math.max(ans, end - base + 1);
//                }
//            }
//
//            base = Math.max(end, base + 1);
//        }
//
//        return ans;
//    }
func officialTwoPointers(a []int) int {
	n := len(a)-1
	var ans, base, end int
	for base <= n {
		end = base
		if end < n && a[end] < a[end+1] {
			for end < n && a[end] < a[end+1] {
				end++
			}

			if end < n && a[end] > a[end+1] {
				for end < n && a[end] > a[end+1] {
					end++
				}

				ans = max(ans, end-base+1)
			}
		}

		base = max(end, base+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}