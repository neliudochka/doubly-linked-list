package main

import (
	"fmt"

	dll "github.com/neliudochka/doubly-linked-list"
)

func main() {
	ddl2 := dll.DoublyLinkedList{}

	err := ddl2.Append('1')
	if err != nil {
		fmt.Println(err.Error())
	}

	ddl2.Insert('5', 2)
	err = ddl2.Insert('m', 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	ddl2.Insert('3', 0)
	ddl2.Insert('4', 1)

	ddl2.PrintAll()
	fmt.Println("len: ", ddl2.Lenght())

	ddl3 := ddl2.Clone()

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
