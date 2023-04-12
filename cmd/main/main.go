package main

import (
	"fmt"

	dll "github.com/neliudochka/doubly-linked-list"
)

func main() {
	dll2 := dll.DoublyLinkedList{}

	fmt.Println("list:")
	dll2.Append('a')
	dll2.Append('a')
	dll2.Insert('m', 1)
	dll2.Insert('f', 3)
	dll2.Insert('p', 0)
	dll2.Append('і')
	dll2.Insert('я', 0)
	dll2.PrintAll()
	fmt.Println("lenght:", dll2.Lenght())

	el, _ := dll2.Delete(3)
	fmt.Printf("\ndeleted element '%c' with index 3 :\n", el)
	dll2.PrintAll()

	fmt.Println("\ndeleted all 'a' elements :")
	dll2.DeleteAll('a')
	dll2.PrintAll()

	fmt.Println("\nget element with index 1 :")
	el, _ = dll2.Get(1)
	fmt.Printf("get: %c\n", el)

	fmt.Println("\ncloned the list and reversed it :")
	dll3 := dll2.Clone()
	dll3.Reverse()
	fmt.Println("dll2 (the original) list: ")
	dll2.PrintAll()
	fmt.Println("\ndll3 (copied and reversed) list: ")
	dll3.PrintAll()

	el = 'f'
	i, _ := dll2.FindFirst(el)
	fmt.Printf("\nfound first element %c in the dll2 list on the %v index\n", el, i)
	i, _ = dll3.FindLast(el)
	fmt.Printf("found last element %c in the dll3 list on the %v index\n", el, i)

	fmt.Println("\ncleared dll2 list :")
	dll2.Clear()
	dll2.PrintAll()

	fmt.Println("\nextended dll3 with dll4 list :")
	fmt.Println("dll4 list:")
	dll4 := dll.DoublyLinkedList{}
	dll4.Append('t')
	dll4.Append('w')
	dll4.Append('o')
	dll4.PrintAll()

	fmt.Println("\nextended dll3 list:")
	dll3.Extend(dll4)
	dll3.PrintAll()
}
