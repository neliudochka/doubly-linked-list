package main

import (
	"fmt"

	dll "github.com/neliudochka/doubly-linked-list"
)

func main() {
	ddl2 := dll.DoublyLinkedList{}

	ddl2.Append('1')
	ddl2.Append('2')
	ddl2.Append('3')
	ddl2.Append('4')
	ddl2.Append('5')
	ddl2.Append('m')

	ddl2.PrintAll()
	fmt.Println("len: ", ddl2.Lenght())

	ddl3, _ := ddl2.Clone()

	ddl3.Append('y')

	ddl3.Reverse()
	ddl3.Insert('y', 0)
	ddl3.PrintAll()

	ddl2.Extend(ddl3)
	ddl2.PrintAll()
	fmt.Println("len: ", ddl2.Lenght())
	ddl2.PrintAll()

	//fmt.Printf("clone: %p %p\n", &ddl2, &ddl3)
}
