package prob0622

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	ast := assert.New(t)
	queue := Constructor(3)

	ast.Equal(true, queue.EnQueue(1))
	ast.Equal(true, queue.EnQueue(2))
	ast.Equal(2, queue.Rear())
	ast.Equal(1, queue.Front())
	ast.Equal(true, queue.EnQueue(3))
	ast.Equal(true, queue.IsFull())
	ast.Equal(false, queue.EnQueue(4))
	ast.Equal(true, queue.DeQueue())
	ast.Equal(true, queue.EnQueue(4))
	ast.Equal(2, queue.Front())
	ast.Equal(4, queue.Rear())
	ast.Equal(true, queue.DeQueue())
	ast.Equal(true, queue.DeQueue())
	ast.Equal(true, queue.DeQueue())
	ast.Equal(false, queue.DeQueue())

	// 第二次操作
	ast.Equal(true, queue.EnQueue(1))
	ast.Equal(true, queue.EnQueue(2))
	ast.Equal(2, queue.Rear())
	ast.Equal(1, queue.Front())
	ast.Equal(true, queue.EnQueue(3))
	ast.Equal(true, queue.IsFull())
	ast.Equal(false, queue.EnQueue(4))
	ast.Equal(true, queue.DeQueue())
	ast.Equal(true, queue.EnQueue(4))
	ast.Equal(2, queue.Front())
	ast.Equal(4, queue.Rear())
	ast.Equal(true, queue.DeQueue())
	ast.Equal(true, queue.DeQueue())
	ast.Equal(true, queue.DeQueue())
	ast.Equal(false, queue.DeQueue())
}