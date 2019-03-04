package prob0067

import "strings"

func addBinary(a string, b string) string {
	na, nb := len(a), len(b)

	if nb > na {
		a, b = b, a
		na, nb = nb, na
	}


	s := make([]string, na+1)
	var carry byte
	carry = '0'
	for na > 0 {

		var current string

		if nb > 0 {
			if a[na-1] + b[nb-1] + carry == '1' * 3 {
				carry = '1'
				current = "1"
			}else if a[na-1] + b[nb-1] + carry == (('1'*2) +'0') {
				carry = '1'
				current = "0"
			}else if a[na-1] + b[nb-1] + carry == (('1') +'0' * 2) {
				carry = '0'
				current = "1"
			}else {
				carry = '0'
				current = "0"
			}
			nb--
		}else {
			if a[na-1] + carry == '1' * 2 {
				carry = '1'
				current = "0"
			}else if a[na-1] + carry == '1' + '0' {
				carry = '0'
				current = "1"
			}else if a[na-1] + carry == '0' * 2 {
				carry = '0'
				current = "0"
			}
		}

		s[na] = current
		na--

	}

	if carry == '1' {
		s[0] = "1"

		return strings.Join(s, "")
	}

	return strings.Join(s[1:], "")
}