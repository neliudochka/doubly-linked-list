package dll

import (
	"fmt"
	"testing"
)

type testCase struct {
	index int
	val   rune
}

func TestLenght(t *testing.T) {
	t.Run("Expect size of en empty list to be equal 0", func(t *testing.T) {
		dll := DoublyLinkedList{}
		if dll.Lenght() != 0 {
			t.Errorf("Expected 0, got %v", dll.Lenght())
		}
	})

	t.Run("Expect to return the number of elements in the list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		count := 5
		for i := 0; i < count; i++ {
			dll.Append('b')
		}
		if dll.Lenght() != count {
			t.Errorf("Expected dll size is %v, got %v", count, dll.Lenght())
		}
	})
}

func TestAppend(t *testing.T) {
	t.Run("Expect size of the non-empty list to increase by 1 after adding the element", func(t *testing.T) {
		dll := DoublyLinkedList{}
		err := dll.Append('B')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dll.Lenght() != 1 {
			t.Errorf("Expected size: 1, got %v", dll.Lenght())
		}

		err = dll.Append('s')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dll.Lenght() != 2 {
			t.Errorf("Expected size: 2, got %v", dll.Lenght())
		}
	})

	t.Run("Expect dll to contain added elements", func(t *testing.T) {
		dll := DoublyLinkedList{}
		tests := []testCase{
			{0, 'ф'},
			{1, 'j'},
		}

		for _, test := range tests {
			err := dll.Append(test.val)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			char, _ := dll.Get(test.index)
			if char != test.val {
				t.Errorf("Expected '%c' on the %v place, got '%c'", test.val, test.index, char)
			}
		}
	})

	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		n := '2'
		err := dll.Append(n)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("Expect method to insert element on the beginning of an empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		l := 'ґ'
		err := dll.Insert(l, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		char, _ := dll.Get(0)
		if char != l {
			t.Errorf("Expected %c, got %v", l, char)
		}
	})

	t.Run("Expect the method to insert elements in different positions in the non-empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Append('a')
		dll.Append('b')
		dll.Append('c')
		tests := []testCase{
			{3, 'ь'},
			{1, 'я'},
			{0, 'ю'},
		}

		for _, test := range tests {
			err := dll.Insert(test.val, test.index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			char, _ := dll.Get(test.index)
			if char != test.val {
				t.Errorf("Expected %c, got %c", test.val, char)
			}
		}
	})

	t.Run("Expect method to insert element at the end of a non-empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		l := 'ґ'
		l2 := 'м'
		dll.Insert(l, 0)
		dll.Insert(l, 0)
		err := dll.Insert(l2, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		char, _ := dll.Get(2)
		if char != l2 {
			t.Errorf("Expected %c, got %v", l, char)
		}
	})

	t.Run("Expect method to insert element in the middle of a non-empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		l := 'ґ'
		l2 := 'м'
		dll.Insert(l, 0)
		dll.Insert(l, 0)
		err := dll.Insert(l2, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		char, _ := dll.Get(1)
		if char != l2 {
			t.Errorf("Expected %c, got %v", l, char)
		}
	})

	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		n := '9'
		err := dll.Insert(n, 0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	//size check
	t.Run("Expect size of the non-empty list to increase by 1 after inserting the element", func(t *testing.T) {
		dll := DoublyLinkedList{}
		err := dll.Insert('B', 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dll.Lenght() != 1 {
			t.Errorf("the first one\nExpected size: 1, got %v", dll.Lenght())
		}

		err = dll.Insert('c', 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dll.Lenght() != 2 {
			t.Errorf("at the end\nExpected size: 2, got %v", dll.Lenght())
		}

		err = dll.Insert('a', 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dll.Lenght() != 3 {
			t.Errorf("at the begining\nExpected size: 3, got %v", dll.Lenght())
		}

		err = dll.Insert('d', 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dll.Lenght() != 4 {
			t.Errorf("in the middle\nExpected size: 4, got %v", dll.Lenght())
		}
	})

	//index check
	t.Run("Expect method to return an error for wrong index argument: index < 0", func(t *testing.T) {
		dll := DoublyLinkedList{}
		index := -1
		l := 'E'
		err := dll.Insert(l, index)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to return an error for wrong index argument: index > size of the list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		l := 'f'

		err := dll.Insert(l, 1)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}

		dll.Insert(l, 0)
		dll.Insert(l, 1)
		index := 3
		err = dll.Insert(l, index)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})
}

// preparing DoublyLinkedList with lowercase letters in alphabetical order
// reruns created DoublyLinkedList and the first element
func prepareList(count int) (*DoublyLinkedList, rune, error) {
	if count > 26 {
		return nil, 0, fmt.Errorf("prepareList: there are only 26 letters in the English alphabet")
	}
	dll := DoublyLinkedList{}
	l1 := 'a'
	for i := 0; i < count; i++ {
		newL := rune(int(l1) + i)
		dll.Append(newL)
	}
	return &dll, l1, nil
}
func TestDelete(t *testing.T) {
	//check del element
	t.Run("Expect method to delete the first and the only element in the list", func(t *testing.T) {
		dll, _, _ := prepareList(1)

		_, err := dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = dll.Get(0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to delete the first element in the list", func(t *testing.T) {
		dll, _, _ := prepareList(3)
		i := 0
		l2, _ := dll.Get(i + 1)

		_, err := dll.Delete(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		char, _ := dll.Get(i)
		if char != l2 {
			t.Errorf("Expected %c, got %c", l2, char)
		}
	})

	t.Run("Expect method to delete the last element in the list", func(t *testing.T) {
		dll, _, _ := prepareList(3)
		l2, _ := dll.Get(dll.Lenght() - 2)

		_, err := dll.Delete(dll.Lenght() - 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		char, _ := dll.Get(dll.Lenght() - 1)
		if char != l2 {
			t.Errorf("Expected %c, got %c", l2, char)
		}
	})

	t.Run("Expect method to delete middle element in the list", func(t *testing.T) {
		dll, _, _ := prepareList(3)
		i := 1
		l3, _ := dll.Get(i + 1)

		_, err := dll.Delete(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		char, _ := dll.Get(i)
		if char != l3 {
			t.Errorf("Expected %c, got %c", l3, char)
		}
	})

	t.Run("Expect method to return an error for an empty list", func(t *testing.T) {
		dll, _, _ := prepareList(0)
		_, err := dll.Delete(0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	//return check
	t.Run("Expect method to return deleted element", func(t *testing.T) {
		dll, _, _ := prepareList(4)
		l1, _ := dll.Get(0)
		l2, _ := dll.Get(1)
		l3, _ := dll.Get(2)
		l4, _ := dll.Get(3)

		rL2, err := dll.Delete(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if rL2 != l2 {
			t.Errorf("in the middle \nExpected %c, got %c", l2, rL2)
		}

		rL1, err := dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if rL1 != l1 {
			t.Errorf("at the begining \nExpected %c, got %c", l1, rL1)
		}

		rL4, err := dll.Delete(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if rL4 != l4 {
			t.Errorf("at the end \nExpected %c, got %c", l4, rL4)
		}

		rL3, err := dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if rL3 != l3 {
			t.Errorf("the only one \nExpected %c, got %c", l3, rL3)
		}
	})

	//index check
	t.Run("Expect method to return an error for wrong index argument: index < 0", func(t *testing.T) {
		dll, _, _ := prepareList(1)
		_, err := dll.Delete(-1)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to return an error for wrong index argument: index >= size of the list", func(t *testing.T) {
		dll, _, _ := prepareList(1)
		_, err := dll.Delete(2)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}

		_, err = dll.Delete(1)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	//size check
	t.Run("Expect size of the non-empty list to decrease by 1 after deleting the element", func(t *testing.T) {
		size := 4
		dll, _, _ := prepareList(size)

		_, err := dll.Delete(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		size--
		if dll.Lenght() != size {
			t.Errorf("in the middle\nExpected size: %v, got %v", size, dll.Lenght())
		}

		_, err = dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		size--
		if dll.Lenght() != size {
			t.Errorf("at the begining\nExpected size: %v, got %v", size, dll.Lenght())
		}

		_, err = dll.Delete(size - 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		size--
		if dll.Lenght() != size {
			t.Errorf("at the end\nExpected size: %v, got %v", size, dll.Lenght())
		}

		_, err = dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		size--
		if dll.Lenght() != size {
			t.Errorf("the last one\nExpected size: %v, got %v", size, dll.Lenght())
		}
	})

}

func prepareListForDeleteAllM() DoublyLinkedList {
	dll := DoublyLinkedList{}
	//4 'a'
	dll.Append('a')
	dll.Append('d')
	dll.Append('a')
	dll.Append('c')
	dll.Append('d')
	dll.Append('a')
	dll.Append('p')
	dll.Append('a')
	return dll
}
func TestDeleteAll(t *testing.T) {
	// del element check
	t.Run("Expect method to delete all elements from the list that are equal to the argument", func(t *testing.T) {
		dll := prepareListForDeleteAllM()
		l := 'a'
		err := dll.DeleteAll(l)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		count := dll.Lenght()
		for i := 0; i < count; i++ {
			char, _ := dll.Get(i)
			if char == l {
				t.Fail()
			}
		}
	})

	// no element in list
	t.Run("Expect method to do nothing if there no elements in the list that are equal to the argument", func(t *testing.T) {
		dll1 := prepareListForDeleteAllM()
		dll2 := prepareListForDeleteAllM()
		l := 'z'
		err := dll1.DeleteAll(l)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if dll1.Lenght() != dll2.Lenght() {
			t.Error()
		}

		count := dll2.Lenght()
		for i := 0; i < count; i++ {
			char1, _ := dll1.Get(i)
			char2, _ := dll2.Get(i)
			if char1 != char2 {
				t.Fail()
			}
		}
	})
	// input check
	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		err := dll.DeleteAll('1')
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	// size check
	t.Run("Expect the size of the list to decrease by the number of the deleted elements", func(t *testing.T) {
		dll := prepareListForDeleteAllM()
		// number of the l elements = 4
		oldLen := dll.Lenght()
		l := 'a'
		err := dll.DeleteAll(l)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if dll.Lenght() != oldLen-4 {
			t.Errorf("Expected len %v, got %v", oldLen-4, dll.Lenght())
		}
	})
}

func prepareListForGetM() DoublyLinkedList {
	dll := DoublyLinkedList{}
	//0
	dll.Append('a')
	dll.Append('w')
	//2
	dll.Append('q')
	dll.Append('c')
	dll.Append('d')
	dll.Append('a')
	dll.Append('p')
	//7
	dll.Append('l')
	return dll
}
func TestGet(t *testing.T) {
	// check returned element
	t.Run("Expect method to return element of the list at a particular index", func(t *testing.T) {
		dll := prepareListForGetM()
		tests := []struct {
			index int
			val   rune
		}{
			{0, 'a'},
			{2, 'q'},
			{7, 'l'},
		}

		for _, test := range tests {
			char, err := dll.Get(test.index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if char != test.val {
				t.Errorf("Expected %c, got %c", test.val, char)
			}
		}
	})

	//index check
	t.Run("Expect method to return an error for an empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		_, err := dll.Get(0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to return an error for wrong index argument: index < 0", func(t *testing.T) {
		dll := prepareListForGetM()
		_, err := dll.Get(-1)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to return an error for wrong index argument: index >= size of the list", func(t *testing.T) {
		dll := prepareListForGetM()
		_, err := dll.Get(dll.Lenght())
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})
}

func prepareList3() DoublyLinkedList {
	dll := DoublyLinkedList{}
	dll.Append('a')
	dll.Append('b')
	dll.Append('c')
	dll.Append('d')
	return dll
}
func TestClone(t *testing.T) {
	//check if they equal
	t.Run("Expect method to return copy of the  list", func(t *testing.T) {
		dll := prepareList3()
		cloneDll := dll.Clone()

		if dll.Lenght() != cloneDll.Lenght() {
			t.Error()
		}

		count := cloneDll.Lenght()
		for i := 0; i < count; i++ {
			char1, _ := dll.Get(i)
			char2, _ := cloneDll.Get(i)
			if char1 != char2 {
				t.Fail()
			}
		}
	})

	//check if they are not the same
	t.Run("Expect method to return an independant list ", func(t *testing.T) {
		dll := DoublyLinkedList{}
		cloneDll := dll.Clone()

		dll.Append('ї')
		_, err := cloneDll.Get(0)
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}

		cloneDll.Append('o')
		l1, _ := dll.Get(0)
		l2, _ := cloneDll.Get(0)
		if l1 == l2 {
			t.Error()
		}

		_, err1 := dll.Get(1)
		_, err2 := cloneDll.Get(1)
		if err1 == nil || err2 == nil {
			t.Errorf("Expected  error, got %v and %v", err1, err2)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("Expect method to reverse list ", func(t *testing.T) {
		dll := prepareList3()
		sameDll := prepareList3()

		dll.Reverse()
		if dll.Lenght() != sameDll.Lenght() {
			t.Error()
		}
		count := sameDll.Lenght()
		for i := 0; i < count; i++ {
			char1, _ := dll.Get(i)
			char2, _ := sameDll.Get(count - 1 - i)
			if char1 != char2 {
				t.Fail()
			}
		}
	})

	t.Run("Expect method to return the same list if there is only one element", func(t *testing.T) {
		dll := DoublyLinkedList{}
		l := 'x'
		dll.Append(l)

		dll.Reverse()
		char, _ := dll.Get(0)
		if char != l {
			t.Errorf("Expected %c, got %c", l, char)
		}
		if dll.Lenght() != 1 {
			t.Errorf("Expected 1, got %v", dll.Lenght())
		}
	})

	t.Run("Expect method to return the same list if the list is empty", func(t *testing.T) {
		dll := DoublyLinkedList{}

		dll.Reverse()
		if dll.Lenght() != 0 {
			t.Errorf("Expected 0, got %v", dll.Lenght())
		}
		_, err := dll.Get(0)
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})
}

func prepareList4() DoublyLinkedList {
	dll := DoublyLinkedList{}
	dll.Append('a')
	dll.Append('b')
	dll.Append('q')
	dll.Append('a')
	dll.Append('x')
	dll.Append('j')
	dll.Append('x')
	dll.Append('q')
	return dll
}
func TestFindFirst(t *testing.T) {
	//find element
	t.Run("Expect method to return index of the first found (start - head) element that matches argument", func(t *testing.T) {
		dll := prepareList4()
		tests := []struct {
			val   rune
			index int
		}{
			{'a', 0},
			{'b', 1},
			{'q', 2},
			{'x', 4},
			{'j', 5},
		}

		for _, test := range tests {
			i, err := dll.FindFirst(test.val)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if i != test.index {
				t.Errorf("Expected %v, got %v", test.index, i)
			}
		}
	})
	//not find element
	t.Run("Expect method to return -1 if there no such element in the list", func(t *testing.T) {
		dll := prepareList4()

		i, err := dll.FindFirst('є')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if i != -1 {
			t.Errorf("Expected -1, got %v", i)
		}
	})
	//empty list
	t.Run("Expect method to return -1 if the list is empty", func(t *testing.T) {
		dll := DoublyLinkedList{}

		i, err := dll.FindFirst('є')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if i != -1 {
			t.Errorf("Expected -1, got %v", i)
		}
	})
	//wrong input
	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}

		_, err := dll.FindFirst('-')
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

}

func TestFindLast(t *testing.T) {
	// find element
	t.Run("Expect method to return index of the first found (start - tail) element that matches argument", func(t *testing.T) {
		dll := prepareList4()
		tests := []struct {
			val   rune
			index int
		}{
			{'a', 3},
			{'b', 1},
			{'q', 7},
			{'x', 6},
			{'j', 5},
		}

		for _, test := range tests {
			i, err := dll.FindLast(test.val)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if i != test.index {
				t.Errorf("Expected %v, got %v", test.index, i)
			}
		}
	})
	// not find element
	t.Run("Expect method to return -1 if there no such element in the list", func(t *testing.T) {
		dll := prepareList4()

		i, err := dll.FindLast('є')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if i != -1 {
			t.Errorf("Expected -1, got %v", i)
		}
	})
	// empty list
	t.Run("Expect method to return -1 if the list is empty", func(t *testing.T) {
		dll := DoublyLinkedList{}

		i, err := dll.FindLast('є')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if i != -1 {
			t.Errorf("Expected -1, got %v", i)
		}
	})
	// wrong input
	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}

		_, err := dll.FindLast('+')
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})
}

func TestClear(t *testing.T) {
	t.Run("Expect method to delete all elements from the list", func(t *testing.T) {
		dll := prepareList4()
		dll.Clear()
		_, err := dll.Get(0)
		if err == nil {
			t.Errorf("Expected error: list is empty, got %v", err)
		}
	})

	t.Run("Expect method to decrease size of the list to 0", func(t *testing.T) {
		dll := prepareList4()
		dll.Clear()
		if dll.Lenght() != 0 {
			t.Errorf("Expected 0, got %v", dll.Lenght())
		}
	})

	t.Run("Expect method to do nothing with an empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Clear()
		if dll.Lenght() != 0 {
			t.Errorf("Expected 0, got %v", dll.Lenght())
		}
	})
}

func prepareListsForExtendM() (DoublyLinkedList, DoublyLinkedList) {
	dll1 := DoublyLinkedList{}
	dll1.Append('a')
	dll1.Append('b')
	dll1.Append('c')

	dll2 := DoublyLinkedList{}
	dll2.Append('я')
	dll2.Append('ю')
	dll2.Append('є')
	dll2.Append('ї')

	return dll1, dll2
}

func TestExtend(t *testing.T) {
	t.Run("Expect method to add all elements from the given list to the current list", func(t *testing.T) {
		dll1, dll2 := prepareListsForExtendM()
		dll1.Extend(dll2)

		tests := []struct {
			index int
			val   rune
		}{
			{3, 'я'},
			{4, 'ю'},
			{5, 'є'},
			{6, 'ї'},
		}

		for _, test := range tests {
			char, err := dll1.Get(test.index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if char != test.val {
				t.Errorf("Expected %c on %v index, got %c", test.val, test.index, char)
			}
		}

	})

	t.Run("Expect method to do nothing if the given list is empty", func(t *testing.T) {
		dll1, _ := prepareListsForExtendM()
		oldLen := dll1.Lenght()
		emptyDll := DoublyLinkedList{}
		dll1.Extend(emptyDll)
		newLen := dll1.Lenght()

		if oldLen != newLen {
			t.Errorf("Expected newLen(%v) be equal to oldLen(%v)", oldLen, newLen)
		}
		_, err := dll1.Get(oldLen)
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})

	t.Run("Expect method to not modify the given list", func(t *testing.T) {
		dll1, dll2 := prepareListsForExtendM()
		oldLen := dll2.Lenght()
		dll1.Extend(dll2)
		newLen := dll2.Lenght()
		dll1.Append('f')

		if oldLen != newLen {
			t.Errorf("Expected newLen(%v) be equal to oldLen(%v)", oldLen, newLen)
		}
		_, err := dll2.Get(oldLen)
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})

	t.Run("Expect method to incraese size of the current list according to the number of added elements", func(t *testing.T) {
		dll1, dll2 := prepareListsForExtendM()
		expectedLen := dll1.Lenght() + dll2.Lenght()
		dll1.Extend(dll2)
		newLen := dll1.Lenght()

		if expectedLen != newLen {
			t.Errorf("Expected newLen(%v) be equal to oldLen(%v)", expectedLen, newLen)
		}
	})
}
