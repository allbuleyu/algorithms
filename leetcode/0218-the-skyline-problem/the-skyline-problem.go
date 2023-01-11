package prob0218

import "sort"

func getSkyline(buildings [][]int) [][]int {
	return bruteForce1(buildings)
}

// Approach 1: Brute Force I
//
//Collect all the positions of the left and right edges from buildings, that's all the possible x where skyline key points are generated. For convenience, let's number these unique positions sequentially, representing these positions by indexes according to their location on the x-axis.
// public:
//
//	   vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
//	       // Sort the unique positions of all the edges.
//	       set<int> edgeSet;
//	       for (auto building : buildings) {
//	           int left = building[0], right = building[1];
//	           edgeSet.insert(left);
//	           edgeSet.insert(right);
//	       }
//	       vector<int> edges(edgeSet.begin(), edgeSet.end());
//	       // Use hash table 'edgeIndexMap' to record every {position : index} pairs in edges.
//	       map<int, int> edgeIndexMap;
//	       for (int i = 0; i < edges.size(); ++i) {
//	           edgeIndexMap[edges[i]] = i;
//	       }
//
//	       // Use 'heights' to record maximum height at each index.
//	       vector<int> heights(edges.size());
//
//	       // Iterate over all the buildings.
//	       for (auto building : buildings) {
//	           // Find the indexes of the left and right edges
//	           // and update the max height within the range [left_idx, right_idx)
//	           int left = building[0], right = building[1], height = building[2];
//	           int leftIndex = edgeIndexMap[left], rightIndex = edgeIndexMap[right];
//	           for (int idx = leftIndex; idx < rightIndex; ++idx) {
//	               heights[idx] = max(heights[idx], height);
//	           }
//	       }
//
//	       vector<vector<int>> answer;
//
//	       // Iterate over 'heights'.
//	       for (int i = 0; i < heights.size(); ++i) {
//	           // Add all the positions where the height changes to 'answer'.
//	           int currHeight = heights[i], currPos = edges[i];
//	           if (i == 0 || currHeight != heights[i - 1]) {
//	               answer.push_back({currPos, currHeight});
//	           }
//	       }
//	       return answer;
//	   }
//	};

// 使用edges 收集所有的buildings中的x轴的坐标,并排序后,记录所有 x 所对应的index
// 使用heights 记录所有edges中对应的最大高度.
func bruteForce1(buildings [][]int) [][]int {
	var res [][]int
	edgeSet := make([]int, 0)
	setMap := make(map[int]struct{})
	for _, b := range buildings {
		left, right := b[0], b[1]
		if _, ok := setMap[left]; !ok {
			edgeSet = append(edgeSet, left)
		}

		if _, ok := setMap[right]; !ok {
			edgeSet = append(edgeSet, right)
		}
	}
	sort.Ints(edgeSet)

	edgeIndexMap := make(map[int]int)
	for i, v := range edgeSet {
		edgeIndexMap[v] = i
	}

	heights := make([]int, len(edgeSet))
	for _, v := range buildings {
		left, right, height := v[0], v[1], v[2]
		leftIndex, rightIndex := edgeIndexMap[left], edgeIndexMap[right]

		// idx < rightIndex 不包括右边界
		for idx := leftIndex; idx < rightIndex; idx++ {
			heights[idx] = max(heights[idx], height)
		}
	}

	for i := 0; i < len(heights); i++ {
		curHeight := heights[i]
		curPos := edgeSet[i]
		if i == 0 || curHeight != heights[i-1] {
			res = append(res, []int{curPos, curHeight})
		}
	}
	return res
}

// Approach 2: Brute Force II, Sweep Line
//
// Another instinctive idea is to use a vertical line of infinite length to sweep over the ground from the left to right. The line stops by every edge and we shall record the maximum height among all the buildings that intersect with the line. As shown in the picture below, the right edge of a building doesn't count!
// class Solution {
// public:
//
//	   vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
//	       // Collect and sort the unique positions of all the edges.
//	       set<int> edgeSet;
//	       for (auto building : buildings) {
//	           edgeSet.insert(building[0]);
//	           edgeSet.insert(building[1]);
//	       }
//	       vector<int> positions(edgeSet.begin(), edgeSet.end());
//
//	       vector<vector<int>> answer;
//
//	       // For each position, draw an imaginary vertical line.
//	       for (auto position : positions) {
//	           int maxHeight = 0;
//	           // Check if any buildings intersect with the line at position.
//	           for (auto building : buildings) {
//	               int left = building[0], right = building[1], height = building[2];
//
//	               // Update 'maxHeight' if necessary.
//	               if (position >= left and position < right) {
//	                   maxHeight = max(maxHeight, height);
//	               }
//	           }
//
//	           // Add all the positions where the height changes to 'answer'.
//	           if (answer.empty() || answer[answer.size() - 1][1] != maxHeight) {
//	               answer.push_back({position, maxHeight});
//	           }
//	       }
//	       return answer;
//	   }
//	};
func bruteForce2(buildings [][]int) [][]int {
	var res [][]int
	edges := make([]int, 0)
	setMap := make(map[int]struct{})
	for _, v := range buildings {
		left, right := v[0], v[1]
		if _, ok := setMap[left]; !ok {
			edges = append(edges, left)
		}

		if _, ok := setMap[right]; !ok {
			edges = append(edges, right)
		}
	}
	sort.Ints(edges)

	heights := make([]int, len(edges))
	for i, x := range edges {
		for _, b := range buildings {
			left, right, height := b[0], b[1], b[2]
			if left <= x && x < right {
				heights[i] = max(heights[i], height)
			}
		}

		if i == 0 || heights[i-1] != heights[i] {
			res = append(res, []int{x, heights[i]})
		}
	}

	return res
}

func priorityQueue(buildings [][]int) [][]int {
	var res [][]int
	edges := make([][]int, 0)
	setMap := make(map[int]struct{})
	for i, v := range buildings {
		left, right := v[0], v[1]
		if _, ok := setMap[left]; !ok {
			edges = append(edges, []int{left, i})
		}

		if _, ok := setMap[right]; !ok {
			edges = append(edges, []int{right, i})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] <= edges[j][0]
	})

	pq := make([][]int, 0)
	for i, edge := range edges {
		curX := edge[0]

		for i < len(edges) && edges[i][0] == curX {
			b := edges[i][1]

			if buildings[b][0] == curX {
				right, height := buildings[b][1], buildings[b][2]

				pq = append(pq, []int{height, right})
			}
			i++
		}

		for len(pq) > 0 && pq[len(pq)-1][1] <= curX {
			pq = pq[:len(pq)-1]
		}

		curHeight := 0
		if len(pq) > 0 {
			curHeight = pq[len(pq)-1][0]
		}

		if len(res) == 0 || res[len(res)-1][1] <= curX {
			res = append(res, []int{curX, curHeight})
		}
	}

	return res
}

func helpDivideAndConquer(buildings [][]int) [][]int {
	//sort.Slice(buildings, func(i, j int) bool {
	//	if buildings[i][0] < buildings[j][0] {
	//		return true
	//	}
	//
	//	return buildings[i][1] < buildings[j][1]
	//})

	return divideAndConquer(buildings, 0, len(buildings)-1)
}

func divideAndConquer(buildings [][]int, left, right int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	if left == right {
		l, r, height := buildings[left][0], buildings[left][1], buildings[left][2]
		return [][]int{{l, height}, {r, 0}}
	}

	mid := left + (right-left)/2
	subLeft := divideAndConquer(buildings, left, mid)
	subRight := divideAndConquer(buildings, mid+1, right)

	return merge(subLeft, subRight)
}

func merge(leftSkyline, rightSkyline [][]int) [][]int {
	var leftPos, rightPos int
	var leftPrevHeight, rightPrevHeight int
	var curX, curY int
	var res [][]int

	for leftPos < len(leftSkyline) && rightPos < len(rightSkyline) {
		nextLeftX, nextRightX := leftSkyline[leftPos][0], rightSkyline[rightPos][0]

		if nextLeftX < nextRightX {
			leftPrevHeight = leftSkyline[leftPos][1]

			curX = nextLeftX
			curY = max(leftPrevHeight, rightPrevHeight)
			leftPos++
		} else if nextLeftX > nextRightX {
			rightPrevHeight = rightSkyline[rightPos][1]

			curX = nextRightX
			curY = max(leftPrevHeight, rightPrevHeight)
			rightPos++
		} else {
			leftPrevHeight = leftSkyline[leftPos][1]
			rightPrevHeight = rightSkyline[rightPos][1]

			curX = nextRightX
			curY = max(leftPrevHeight, rightPrevHeight)
			leftPos++
			rightPos++
		}

		if len(res) == 0 || res[len(res)-1][1] != curY {
			res = append(res, []int{curX, curY})
		}
	}

	for leftPos < len(leftSkyline) {
		res = append(res, leftSkyline[leftPos])
		leftPos++
	}

	for rightPos < len(rightSkyline) {
		res = append(res, rightSkyline[rightPos])
		rightPos++
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
