package prob0707

import (
	"fmt"
	"testing"
)

func TestProb0707(t *testing.T) {
	// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
	//[[],[1],[3],[1,2],[1],[1],[1]]
	myList := Constructor()

	myList.AddAtHead(1)
	myList.AddAtTail(3)
	myList.AddAtIndex(1, 2)

	myList.DeleteAtIndex(1)

	//myList.AddAtIndex(2, 100)
	//myList.DeleteAtIndex(1)
	fmt.Println(myList.ConvertToArr())

}