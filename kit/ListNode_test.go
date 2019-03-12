package kit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNodes(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6}

	head := CreateNodes(nums)

	tnums := TransferNodes(head)

	ast.Equal(nums, tnums)
}
