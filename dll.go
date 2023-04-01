package dll

import "fmt"

type Node struct {
	value rune
	prev  *Node
	next  *Node
}

// Вивести всі елементи списку
func (node *Node) Print() {
	fmt.Printf("current: %p, value: %c, prev: %p, next: %p \n", node, node.value, node.prev, node.next)
}

type DoublyLinkedList struct {
	size int
	head *Node
	tail *Node
}

// Вивести всі елементи списку
func (dll *DoublyLinkedList) PrintAll() {
	if dll.head == nil {
		fmt.Printf("List is empty\n")
		return
	}
	fmt.Printf("<---head---> \n")
	currentNode := dll.head
	for currentNode != nil {
		currentNode.Print()
		currentNode = currentNode.next
	}
	fmt.Printf("<---tail---> \n")
}

// Операція визначення довжини списку.
// Якщо список непорожній, то ця операція повинна повертати кількість елементів у списку.
// Якщо список порожній, то ця операція повинна повертати 0.
func (dll *DoublyLinkedList) Lenght() int {
	return dll.size
}

// Операцію додавання елементу в кінець списку.
// нічо не повертає
func (dll *DoublyLinkedList) Append(element rune) {
	var newNode = Node{value: element}
	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = &newNode
		dll.tail = &newNode
	}
	dll.size++
}

// Операцію вставки елементу на довільну позицію у списку. Нумерація елементів списку починається з 0.
// -> None
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за кількість елементів у списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Insert(element rune, index int) {
	switch {
	case index < 0:
		fmt.Println("Error: index < 0")
	case index >= dll.size:
		fmt.Printf("Error: index > кількість елементів у списку (%v)", dll.size)
	case index == 0:
		newNode := Node{value: element, next: dll.head}
		dll.head.prev = &newNode
		dll.head = &newNode
		dll.size++
	default:
		nextNode := dll.head
		for count := 0; count != index; count++ {
			nextNode = nextNode.next
		}
		prevNode := nextNode.prev
		newNode := Node{value: element, prev: prevNode, next: nextNode}
		prevNode.next = &newNode
		nextNode.prev = &newNode
		dll.size++
	}
}

// Операцію видалення елементу зі списку на вказаній позиції.
// Метод повинен повертати значення того елементу, який видаляється. Нумерація елементів списку починається з 0.
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за індекс останнього елементу списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Delete(index int) rune {
	switch {
	case index < 0:
		fmt.Println("Error: index < 0")
	case index >= dll.size:
		fmt.Printf("Error: index > кількість елементів у списку (%v)", dll.size)
	case dll.size == 1:
		element := dll.head.value
		dll.head = nil
		dll.tail = nil
		dll.size--
		return element
	case index == 0:
		element := dll.head.value
		dll.head = dll.head.next
		dll.head.prev = nil
		dll.size--
		return element
	case index == dll.size-1:
		element := dll.tail.value
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		dll.size--
		return element
	default:
		delNode := dll.head
		for count := 0; count != index; count++ {
			delNode = delNode.next
		}
		delNode.prev.next = delNode.next
		delNode.next.prev = delNode.prev
		dll.size--
		return delNode.value
	}
	return 0
}

// Операцію видалення елементів зі списку за значенням. Метод видаляє зі списку усі елементи, які за значенням відповідають шуканому.
// -> None
// У випадку передачі елемента, який у списку відсутній, жодні зміни до списку не застосовуються.
func (dll *DoublyLinkedList) DeleteAll(element rune) {
	currentNode := dll.head
	count := 0
	for currentNode != nil {
		nextNode := currentNode.next
		if currentNode.value == element {
			dll.Delete(count)
		} else {
			count++
		}
		currentNode = nextNode
	}
}

// Операцію отримання елементу списку на довільній позиції
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за індекс останнього елементу списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Get(index int) rune {
	switch {
	case index < 0:
		fmt.Println("Error: index < 0")
	case index >= dll.size:
		fmt.Printf("Error: index > кількість елементів у списку (%v)", dll.size)
	default:
		currentNode := dll.head
		for count := 0; count != index; count++ {
			currentNode = currentNode.next
		}
		return currentNode.value
	}
	return 0
}

// Операцію копіювання списку. При виклику повинен створити копію поточного списку та повернути її.
func (dll *DoublyLinkedList) Clone() DoublyLinkedList {
	dll2 := DoublyLinkedList{}
	currentNode := dll.head
	for currentNode != nil {
		dll2.Append(currentNode.value)
		currentNode = currentNode.next
	}
	return dll2
}

// Операцію обернення списку. Метод повинен змінити порядок елементів у поточному списку задом наперед.
// Елемент, що був останнім стане першим, передостаннім — другим, … а перший — останнім.
func (dll *DoublyLinkedList) Reverse() {
	switch {
	case dll.size > 1:
		//tailNode := dll.tail
		currentNode := dll.head
		for currentNode != nil {
			oldNextNode := currentNode.next
			currentNode.next = currentNode.prev
			currentNode.prev = oldNextNode
			currentNode = oldNextNode
		}
		oldHeadNode := dll.head
		dll.head = dll.tail
		dll.tail = oldHeadNode
	case dll.size == 1:
		fmt.Println("only one element, so list won't change its structure")
	default:
		fmt.Println("List is empty!")
	}
}

// Операцію пошуку елемента за значенням з голови списку.
// Метод повинен знайти перший елемент у списку, що дорівнює шуканому та повернути його позицію.
// Нумерація елементів списку починається з 0. У випадку відсутності шуканого елемента у списку, метод повертає -1
func (dll *DoublyLinkedList) FindFirst(element rune) int {
	currentNode := dll.head
	count := 0
	for currentNode != nil {
		if currentNode.value == element {
			return count
		}
		count++
		currentNode = currentNode.next
	}
	return -1
}

// Операцію пошуку елемента за значенням з хвоста списку.
// Метод повинен знайти останній елемент у списку, що дорівнює шуканому та повернути його позицію.
// Нумерація елементів списку починається з 0. У випадку відсутності шуканого елемента у списку, метод повертає -1.
func (dll *DoublyLinkedList) FindLast(element rune) int {
	currentNode := dll.tail
	count := dll.size - 1
	for currentNode != nil {
		if currentNode.value == element {
			return count
		}
		count--
		currentNode = currentNode.prev
	}
	return -1
}

// Операцію очищення списку. Метод видаляє усі елементи списку.
// -> None
func (dll *DoublyLinkedList) Clear() {
	if dll.size > 0 {
		dll.head = nil
		dll.tail = nil
		dll.size = 0
	} else {
		fmt.Println("List is empty!")
	}
}

// Операцію розширення списку. Метод приймає інший список та додає до поточного списку усі елементи останнього.
// При цьому подальші зміни в другий список не повинні впливати на перший.
// -> None
func (dll *DoublyLinkedList) Extend(dll2 DoublyLinkedList) {
	newDll := dll2.Clone()
	fmt.Println("clone:")
	newDll.PrintAll()
	dll.tail.next = newDll.head
	newDll.head.prev = dll.tail
	dll.tail = newDll.tail
}
