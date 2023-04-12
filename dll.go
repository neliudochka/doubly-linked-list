package dll

import (
	"fmt"
	"unicode"
)

type Node struct {
	value rune
	prev  *Node
	next  *Node
}

// Вивести вузол
func (node *Node) Print() {
	fmt.Printf("current: %p, value: %c, prev: %p, next: %p \n", node, node.value, node.prev, node.next)
}

// Двозв'язний список
// Нумерація елементів списку починається з 0.
type DoublyLinkedList struct {
	size int
	head *Node
	tail *Node
}

// Виводить усі вузли списку
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

// Визначає довжину списку
// Якщо список непорожній, то повертає кількість елементів у списку.
// Якщо список порожній, то повертає 0.
func (dll *DoublyLinkedList) Lenght() int {
	return dll.size
}

// Додає елемент у кінець списку
func (dll *DoublyLinkedList) Append(element rune) error {
	if !unicode.IsLetter(element) {
		return fmt.Errorf("Append error: non-letter argument")
	}
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
	return nil
}

// Вставляє елемент на довільну позицію у списку.
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за кількість елементів у списку) генерує виключну ситуацію
func (dll *DoublyLinkedList) Insert(element rune, index int) error {
	if !unicode.IsLetter(element) {
		return fmt.Errorf("Insert error: non-letter argument")
	}
	switch {
	case index < 0:
		return fmt.Errorf("Insert error: index < 0")
	case index > dll.size:
		return fmt.Errorf(" Insert error: index > last index of the element in the list + 1")
	case index == 0:
		newNode := Node{value: element}
		if dll.size == 0 {
			dll.head = &newNode
			dll.tail = &newNode
		} else {
			dll.head.prev = &newNode
			newNode.next = dll.head
			dll.head = &newNode
		}
	case index == dll.size:
		newNode := Node{value: element, prev: dll.tail}
		dll.tail.next = &newNode
		dll.tail = &newNode
	default:
		nextNode := dll.head
		for count := 0; count != index; count++ {
			nextNode = nextNode.next
		}
		prevNode := nextNode.prev
		newNode := Node{value: element, prev: prevNode, next: nextNode}
		prevNode.next = &newNode
		nextNode.prev = &newNode
	}
	dll.size++
	return nil
}

// Видаляє елемент зі списку на вказаній позиції.
// Повертає значення того елементу, який видаляється.
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за індекс останнього елементу списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Delete(index int) (rune, error) {
	var element rune = 0
	switch {
	case dll.size == 0:
		return 0, fmt.Errorf("Delete error: list is empty")
	case index < 0:
		return 0, fmt.Errorf("Delete error: index < 0")
	case index >= dll.size:
		return 0, fmt.Errorf("Delete error: index > index of the last element in the list")
	case dll.size == 1:
		element = dll.head.value
		dll.head = nil
		dll.tail = nil
	case index == 0:
		element = dll.head.value
		dll.head = dll.head.next
		dll.head.prev = nil
	case index == dll.size-1:
		element = dll.tail.value
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	default:
		delNode := dll.head
		for count := 0; count != index; count++ {
			delNode = delNode.next
		}
		delNode.prev.next = delNode.next
		delNode.next.prev = delNode.prev
		element = delNode.value
	}
	dll.size--
	return element, nil
}

// Видаляєя зі списку усі елементи, які за значенням відповідають переданому.
// У випадку передачі елемента, який у списку відсутній, жодні зміни до списку не застосовуються.
func (dll *DoublyLinkedList) DeleteAll(element rune) error {
	if !unicode.IsLetter(element) {
		return fmt.Errorf("DeleteAll error: non-letter argument")
	}
	currentNode := dll.head
	count := 0
	for currentNode != nil {
		nextNode := currentNode.next
		if currentNode.value == element {
			_, err := dll.Delete(count)
			if err != nil {
				return fmt.Errorf("DeleteAll error: %v", err)
			}
		} else {
			count++
		}
		currentNode = nextNode
	}
	return nil
}

// Отримує елемент списку з довільної позиції
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за індекс останнього елементу списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Get(index int) (rune, error) {
	switch {
	case dll.size == 0:
		return 0, fmt.Errorf("Get error: list is empty")
	case index < 0:
		return 0, fmt.Errorf("Get error: index < 0")
	case index >= dll.size:
		return 0, fmt.Errorf("Get error: index >= size of the list")
	default:
		currentNode := dll.head
		for count := 0; count != index; count++ {
			currentNode = currentNode.next
		}
		return currentNode.value, nil
	}
}

// Копіює поточний список та повертає його копію.
func (dll *DoublyLinkedList) Clone() DoublyLinkedList {
	dll2 := DoublyLinkedList{}
	currentNode := dll.head
	for currentNode != nil {
		dll2.Append(currentNode.value)
		currentNode = currentNode.next
	}
	return dll2
}

// Змінює порядок елементів у поточному списку задом наперед.
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
	}
}

// Шукає переданий елемент з голови списку. Повертає перший знайдений, що дорівнює шуканому, та повертає його позицію.
// У випадку відсутності шуканого елемента у списку, метод повертає -1
func (dll *DoublyLinkedList) FindFirst(element rune) (int, error) {
	if !unicode.IsLetter(element) {
		return -1, fmt.Errorf("FindFirst error: non-letter argument")
	}
	currentNode := dll.head
	count := 0
	for currentNode != nil {
		if currentNode.value == element {
			return count, nil
		}
		count++
		currentNode = currentNode.next
	}
	return 0, nil
}

// Шукає переданий елемент з хвоста списку. Повертає перший знайдений, що дорівнює шуканому, та повертає його позицію.
// У випадку відсутності шуканого елемента у списку, метод повертає -1
func (dll *DoublyLinkedList) FindLast(element rune) (int, error) {
	if !unicode.IsLetter(element) {
		return -1, fmt.Errorf("FindLast error: non-letter argument")
	}
	currentNode := dll.tail
	count := dll.size - 1
	for currentNode != nil {
		if currentNode.value == element {
			return count, nil
		}
		count--
		currentNode = currentNode.prev
	}
	return -1, nil
}

// Видаляє усі елементи списку.
func (dll *DoublyLinkedList) Clear() {
	if dll.size > 0 {
		dll.head = nil
		dll.tail = nil
		dll.size = 0
	}
}

// Розширює список.
// Приймає інший список та додає до поточного списку усі елементи останнього.
// При цьому подальші зміни в другий список не впливають на перший.
func (dll *DoublyLinkedList) Extend(dll2 DoublyLinkedList) {
	switch {
	case dll2.size == 0:
	case dll.size == 0:
		newDll := dll2.Clone()
		dll = &newDll
		dll.size += newDll.size
	default:
		newDll := dll2.Clone()
		dll.tail.next = newDll.head
		newDll.head.prev = dll.tail
		dll.tail = newDll.tail
		dll.size += dll2.size
	}
}
