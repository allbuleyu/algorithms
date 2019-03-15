package kit

import (
	"fmt"
	"testing"
)

var bst BinarySearchTree
func fillTree(bst *BinarySearchTree)  {
	bst.Insert(8)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(13)
	bst.Insert(14)
}

func TestTree(t *testing.T)  {
	fillTree(&bst)

	inorderInts := bst.InOrder()
	fmt.Println("inorderInts:",inorderInts)

	preOrderInts := bst.PreOrder()
	fmt.Println("preOrderInts:",preOrderInts)

	postOrderInts := bst.PostOrder()
	fmt.Println("postOrderInts:",postOrderInts)

	max := bst.Max()
	fmt.Println("max:",max)

	min := bst.Min()
	fmt.Println("min:",min)

	bst.String()
}
