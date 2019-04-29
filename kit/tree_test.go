package kit

import (
	"fmt"
	"testing"
)

func TestNetTrees(t *testing.T) {
	tc := []int{1, Null, 2, 3}

	nodes := NetTrees(tc)

	fmt.Println(nodes.PostOrder())
}
