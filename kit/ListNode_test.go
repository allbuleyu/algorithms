package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNodes(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6}

	head := CreateNodes(nums)

	tnums := TransferNodes(head)

	ast.Equal(nums, tnums)
}
