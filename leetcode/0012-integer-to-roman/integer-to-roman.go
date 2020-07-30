package prob0012

func intToRoman(num int) string {
	return greedy(num)
}

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
func greedy(num int) string {
	ans := ""
	for i := 0; i < len(values) && num > 0; i++ {
		if values[i] > num {
			continue
		}

		for num >= values[i] {
			ans += symbols[i]
			num -= values[i]
		}
	}

	return ans
}

var mm = map[int]string {
	1:"I",
	4:"IV",
	5:"V",
	9:"IX",
	10:"X",
	40:"XL",
	50:"L",
	90:"XC",
	100:"C",
	400:"CD",
	500:"D",
	900:"CM",
	1000:"M",
}
var t = []int{1000, 100, 10, 1}

func f1(num int) string {
	ans := ""

	for i := 0; i < len(t) && num > 0; i++ {
		x := num / t[i]
		for x > 0 {
			switch x {
			case 4, 5, 9:
				ans += mm[x*t[i]]
				x -= x
			default:
				if x > 5 {
					ans += mm[5*t[i]]
					x -= 5
				}else {
					ans += mm[1*t[i]]
					x -= 1
				}
			}
		}

		num %= t[i]
	}

	return ans
}