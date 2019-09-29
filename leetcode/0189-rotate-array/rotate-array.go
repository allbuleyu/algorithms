package prob0189

func rotate(nums []int, k int)  {
	k = k % len(nums)
	reverseArr(nums, k)
}

// public class Solution {
//    public void rotate(int[] nums, int k) {
//        k = k % nums.length;
//        int count = 0;
//        for (int start = 0; count < nums.length; start++) {
//            int current = start;
//            int prev = nums[start];
//            do {
//                int next = (current + k) % nums.length;
//                int temp = nums[next];
//                nums[next] = prev;
//                prev = temp;
//                current = next;
//                count++;
//            } while (start != current);
//        }
//    }
//}
//环状替换
func cycleReplace(nums []int, k int) {
	count := 0
	for start := 0; count  < len(nums); start++ {
		current := start
		prev := nums[start]

		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}


func extraArr(nums []int, k int) {
	ans := make([]int, 0)

	ans = append(ans, nums[len(nums)-k:]...)
	ans = append(ans, nums[:len(nums)-k]...)

	for i := 0; i < len(nums); i++ {
		nums[i] = ans[i]
	}
}

func reverseArr(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j;  {
		nums[i], nums[j] = nums[j], nums[i]

		i++
		j--
	}
}