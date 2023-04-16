package dll

import (
	"fmt"
	"unicode"
)

// Двозв'язний список
// Нумерація елементів списку починається з 0.
type DoublyLinkedList struct {
	elements []rune
}

// Виводить усі вузли списку
func (dll *DoublyLinkedList) PrintAll() {
	if len(dll.elements) == 0 {
		fmt.Printf("List is empty\n")
		return
	}
	fmt.Printf("<---head---> \n")
	for i, el := range dll.elements {
		fmt.Printf("index: %v; element: %c\n", i, el)
	}
	fmt.Printf("<---tail---> \n")
}

// Визначає довжину списку
// Якщо список непорожній, то повертає кількість елементів у списку.
// Якщо список порожній, то повертає 0.
func (dll *DoublyLinkedList) Lenght() int {
	return len(dll.elements)
}

// Додає елемент у кінець списку
func (dll *DoublyLinkedList) Append(element rune) error {
	if !unicode.IsLetter(element) {
		return fmt.Errorf("Append error: non-letter argument")
	}
	dll.elements = append(dll.elements, element)
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
	case index > len(dll.elements):
		return fmt.Errorf(" Insert error: index > last index of the element in the list + 1")
	case index == len(dll.elements):
		dll.elements = append(dll.elements, element)
	default:
		dll.elements = append(dll.elements[:index+1], dll.elements[index:]...)
		dll.elements[index] = element
	}
	return nil
}

// Видаляє елемент зі списку на вказаній позиції.
// Повертає значення того елементу, який видаляється.
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за індекс останнього елементу списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Delete(index int) (rune, error) {
	switch {
	case len(dll.elements) == 0:
		return 0, fmt.Errorf("Delete error: list is empty")
	case index < 0:
		return 0, fmt.Errorf("Delete error: index < 0")
	case index >= len(dll.elements):
		return 0, fmt.Errorf("Delete error: index > index of the last element in the list")
	default:
		element := dll.elements[index]
		dll.elements = append(dll.elements[:index], dll.elements[index+1:]...)
		return element, nil
	}
}

// Видаляєя зі списку усі елементи, які за значенням відповідають переданому.
// У випадку передачі елемента, який у списку відсутній, жодні зміни до списку не застосовуються.
func (dll *DoublyLinkedList) DeleteAll(element rune) error {
	if !unicode.IsLetter(element) {
		return fmt.Errorf("DeleteAll error: non-letter argument")
	}
	index := 0
	len := len(dll.elements)
	for j := 0; j < len; j++ {
		if dll.elements[index] == element {
			_, err := dll.Delete(index)
			if err != nil {
				return fmt.Errorf("DeleteAll error: %v", err)
			}
		} else {
			index++
		}
	}
	return nil
}

// Отримує елемент списку з довільної позиції
// У випадку передачі некоректного значення позиції
// (наприклад, від’ємне число, або число, більше за індекс останнього елементу списку) метод повинен генерувати виключну ситуацію
func (dll *DoublyLinkedList) Get(index int) (rune, error) {
	switch {
	case len(dll.elements) == 0:
		return 0, fmt.Errorf("Get error: list is empty")
	case index < 0:
		return 0, fmt.Errorf("Get error: index < 0")
	case index >= len(dll.elements):
		return 0, fmt.Errorf("Get error: index >= size of the list")
	default:
		return dll.elements[index], nil
	}
}

// Копіює поточний список та повертає його копію.
func (dll *DoublyLinkedList) Clone() DoublyLinkedList {
	dll2 := DoublyLinkedList{}
	cloneSlice := make([]rune, len(dll.elements))
	copy(cloneSlice, dll.elements)
	dll2.elements = cloneSlice
	return dll2
}

// Змінює порядок елементів у поточному списку задом наперед.
// Елемент, що був останнім стане першим, передостаннім — другим, … а перший — останнім.
func (dll *DoublyLinkedList) Reverse() {
	if len(dll.elements) > 1 {
		for i, j := 0, len(dll.elements)-1; i < j; i, j = i+1, j-1 {
			dll.elements[i], dll.elements[j] = dll.elements[j], dll.elements[i]
		}
	}
}

// Шукає переданий елемент з голови списку. Повертає перший знайдений, що дорівнює шуканому, та повертає його позицію.
// У випадку відсутності шуканого елемента у списку, метод повертає -1
func (dll *DoublyLinkedList) FindFirst(element rune) (int, error) {
	if !unicode.IsLetter(element) {
		return -1, fmt.Errorf("FindFirst error: non-letter argument")
	}
	for i := 0; i < len(dll.elements); i++ {
		if dll.elements[i] == element {
			return i, nil
		}
	}
	return -1, nil
}

// Шукає переданий елемент з хвоста списку. Повертає перший знайдений, що дорівнює шуканому, та повертає його позицію.
// У випадку відсутності шуканого елемента у списку, метод повертає -1
func (dll *DoublyLinkedList) FindLast(element rune) (int, error) {
	if !unicode.IsLetter(element) {
		return -1, fmt.Errorf("FindLast error: non-letter argument")
	}
	for i := len(dll.elements) - 1; i >= 0; i-- {
		if dll.elements[i] == element {
			return i, nil
		}
	}
	return -1, nil
}

// Видаляє усі елементи списку.
func (dll *DoublyLinkedList) Clear() {
	if len(dll.elements) > 0 {
		dll.elements = []rune{}
	}
}

// Розширює список.
// Приймає інший список та додає до поточного списку усі елементи останнього.
// При цьому подальші зміни в другий список не впливають на перший.
func (dll *DoublyLinkedList) Extend(dll2 DoublyLinkedList) {
	if len(dll2.elements) != 0 {
		dll.elements = append(dll.elements, dll2.elements...)
	}
}
