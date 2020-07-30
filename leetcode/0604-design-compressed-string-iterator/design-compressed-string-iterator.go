package prob0604

import "strconv"

type StringIterator struct {
	s string
	curIndex int
	remain int
}


func Constructor(compressedString string) StringIterator {
	curIndex := 0
	remain := 1
	i := 1
	for ; i < len(compressedString); i++ {
		if compressedString[i] < '0' || compressedString[i] > '9' {
			break
		}
	}

	if i > remain {
		remain, _ = strconv.Atoi(compressedString[1:i])
	}

	return StringIterator{
		s:compressedString,
		curIndex: curIndex,
		remain: remain,
	}
}


func (this *StringIterator) Next() byte {
	if !this.HasNext() {
		return ' '
	}

	cb := this.s[this.curIndex]
	this.remain--

	if this.remain == 0 {
		this.curIndex++

		for this.HasNext() {
			if this.s[this.curIndex] < '0' || this.s[this.curIndex] > '9' {
				break
			}

			this.curIndex++
		}

		if this.HasNext() {
			i := this.curIndex+1
			for i < len(this.s) {
				if this.s[i] < '0' || this.s[i] > '9' {
					break
				}

				i++
			}


			this.remain, _ = strconv.Atoi(this.s[this.curIndex+1:i])
		}
	}

	return cb
}


func (this *StringIterator) HasNext() bool {
	if this.curIndex == len(this.s) {
		return false
	}

	return true
}


/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */