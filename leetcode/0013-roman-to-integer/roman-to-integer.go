package prob0013

func romanToInt(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		if i + 1 < n {
			switch s[i:i+2] {
			case "IV", "IX", "XL", "XC", "CD", "CM":
				ans += m[s[i:i+2]]
				i++
				continue
			default:

			}
		}

		ans += m[string(s[i])]
	}

	return ans
}

// I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
var m = map[string]int{
	"I":1,
	"IV":4,
	"V":5,
	"IX":9,
	"X":10,
	"XL":40,
	"L":50,
	"XC":90,
	"C":100,
	"CD":400,
	"D":500,
	"CM":900,
	"M":1000,
}