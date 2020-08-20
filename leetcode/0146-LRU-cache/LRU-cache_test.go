package prob0146

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)
var null = math.MinInt32
func TestConstructor(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		methods []string
		params [][]int
		ans []int
	}{
		{
			[]string{"LRUCache","get","put","get","put","put","get","get"},
			[][]int{[]int{2},[]int{2},[]int{2,6},[]int{1},[]int{1,5},[]int{1,2},[]int{1},[]int{2}},
			[]int{null,-1,null,-1,null,null,2,6},
		},
		{
			[]string{"LRUCache","put","put","get","get","put","get","get","get"},
			[][]int{[]int{2},[]int{2,1},[]int{3,2},[]int{3},[]int{2},[]int{4,3},[]int{2},[]int{3},[]int{4}},
			[]int{null,null,null,2,1,null,1,-1,3},
		},



	}


	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		lru := Constructor(tc.params[0][0])

		v := reflect.ValueOf(&lru)
		ans := make([]int, 1, len(tc.methods))
		ans[0] = null
		for i := 1; i < len(tc.methods); i++ {
			vp := make([]reflect.Value, 0)
			for _, p := range tc.params[i] {
				vp = append(vp, reflect.ValueOf(p))
			}
			ss := []byte(tc.methods[i])
			ss[0] = ss[0]-'a'+'A'

			r := v.MethodByName(string(ss)).Call(vp)

			if len(r) == 0 {
				ans = append(ans, null)
				continue
			}

			for ii := 0; ii < len(r); ii++ {
				ans = append(ans, int(r[ii].Int()))
			}
		}

		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}