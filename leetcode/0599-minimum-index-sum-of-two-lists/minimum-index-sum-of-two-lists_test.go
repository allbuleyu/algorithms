package prob0599

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input []string
		input2 []string
		ans []string
	}{
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"KFC", "Shogun", "Burger King"},
			[]string{"Shogun"},
		},
		{
			[]string{"Shogun", "KFC", "Burger King"},
			[]string{"KFC", "Shogun", "Burger King"},
			[]string{"KFC","Shogun"},
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findRestaurant(tc.input, tc.input2), "输入:%v", tc)
	}
}