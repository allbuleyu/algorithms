package prob0707

import (
	"fmt"
	"testing"
)

func TestProb0707(t *testing.T) {
	myList := Constructor()

	myList.AddAtHead(1)
	myList.AddAtHead(2)
	myList.AddAtHead(3)

	myList.AddAtIndex(2, 100)
	fmt.Println(myList.ConvertToArr(), myList.Get(3))

}