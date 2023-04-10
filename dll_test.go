package dll

import (
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

		for i := 0; i < 2; i++ {
			oldLen := dll.Lenght()
			err := dll.Append('B')
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dll.Lenght() != oldLen+1 {
				t.Errorf("Expected size: %v, got %v", oldLen+1, dll.Lenght())
			}
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

	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		n := '9'
		err := dll.Insert(n, 0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect size of list to increase by 1 after inserting the elements", func(t *testing.T) {
		dll := DoublyLinkedList{}
		tests := []testCase{
			{0, 'B'},
			{1, 'c'},
			{0, 'a'},
			{1, 'd'},
		}

		for _, test := range tests {
			oldLen := dll.Lenght()
			err := dll.Insert(test.val, test.index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dll.Lenght() != oldLen+1 {
				t.Errorf("Expected size: %v, got %v", oldLen+1, dll.Lenght())
			}
		}
	})

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

// preparing DoublyLinkedList for TestDelete func
func prepareListForDeleteM() *DoublyLinkedList {
	dll := DoublyLinkedList{}
	l1 := 'a'
	for i := 0; i < 4; i++ {
		newL := rune(int(l1) + i)
		dll.Append(newL)
	}
	return &dll
}
func TestDelete(t *testing.T) {
	t.Run("Expect method to delete the first and the only element in the list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Append('r')

		_, err := dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = dll.Get(0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to delete the elements in the list on the different positions (at the beginning, in the middle)", func(t *testing.T) {
		//what val i want to have on position with according index after deletion
		tests := []testCase{
			{0, 'b'},
			{1, 'c'},
			{2, 'd'},
		}
		for _, test := range tests {
			dll := prepareListForDeleteM()
			_, err := dll.Delete(test.index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			char, _ := dll.Get(test.index)
			if char != test.val {
				t.Errorf("Expected %c, got %c", test.val, char)
			}
		}
	})

	//del?
	t.Run("Expect method to delete the elements in the list on the different positions (at the end)", func(t *testing.T) {
		dll := prepareListForDeleteM()
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

	t.Run("Expect method to return an error for an empty list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		_, err := dll.Delete(0)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to return deleted element (del at the begining, in the middle, at the end)", func(t *testing.T) {
		tests := []testCase{
			{0, 'a'},
			{1, 'b'},
			{3, 'd'},
		}
		for _, test := range tests {
			dll := prepareListForDeleteM()
			l, err := dll.Delete(test.index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if l != test.val {
				t.Errorf("Expected %c, got %c", test.val, l)
			}
		}
	})

	t.Run("Expect method to return deleted element (the only one element in the list)", func(t *testing.T) {
		dll := DoublyLinkedList{}
		l := 'p'
		dll.Append(l)
		char, err := dll.Delete(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if char != l {
			t.Errorf("Expected %c, got %c", l, char)
		}
	})

	t.Run("Expect method to return an error for wrong index argument: index < 0", func(t *testing.T) {
		dll := prepareListForDeleteM()
		_, err := dll.Delete(-1)
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect method to return an error for wrong index argument: index >= size of the list", func(t *testing.T) {
		dll := prepareListForDeleteM()
		_, err := dll.Delete(dll.Lenght())
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect size of the non-empty list to decrease by 1 after deleting the element", func(t *testing.T) {
		dll := prepareListForDeleteM()

		indexes := []int{3, 1, 0, 0}
		for _, index := range indexes {
			size := dll.Lenght()
			_, err := dll.Delete(index)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dll.Lenght() != size-1 {
				t.Errorf("Expected size: %v, got %v", size-1, dll.Lenght())
			}
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
	t.Run("Expect method to delete all elements from the list that are equal to the argument", func(t *testing.T) {
		dll := prepareListForDeleteAllM()
		l := 'a'
		err := dll.DeleteAll(l)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for i := 0; i < dll.Lenght(); i++ {
			char, _ := dll.Get(i)
			if char == l {
				t.Fail()
			}
		}
	})

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

	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		err := dll.DeleteAll('1')
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})

	t.Run("Expect the size of the list to decrease by the number of the deleted elements", func(t *testing.T) {
		dll := prepareListForDeleteAllM()
		// number of the 'a' = 4
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
		tests := []testCase{
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

func TestClone(t *testing.T) {
	t.Run("Expect method to return copy of the  list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Append('a')
		dll.Append('b')
		dll.Append('c')
		dll.Append('c')
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

func prepareListForReverseM() DoublyLinkedList {
	dll := DoublyLinkedList{}
	dll.Append('k')
	dll.Append('t')
	dll.Append('c')
	dll.Append('v')
	return dll
}
func TestReverse(t *testing.T) {
	t.Run("Expect method to reverse list ", func(t *testing.T) {
		dll := prepareListForReverseM()
		sameDll := prepareListForReverseM()

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

func prepareListFindM() DoublyLinkedList {
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
	t.Run("Expect method to return index of the first found (start - head) element that matches argument", func(t *testing.T) {
		dll := prepareListFindM()
		tests := []testCase{
			{0, 'a'},
			{1, 'b'},
			{2, 'q'},
			{4, 'x'},
			{5, 'j'},
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

	t.Run("Expect method to return -1 if there no such element in the list", func(t *testing.T) {
		dll := prepareListFindM()
		i, err := dll.FindFirst('є')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if i != -1 {
			t.Errorf("Expected -1, got %v", i)
		}
	})

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

	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		_, err := dll.FindFirst('-')
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})
}

func TestFindLast(t *testing.T) {
	t.Run("Expect method to return index of the first found (start - tail) element that matches argument", func(t *testing.T) {
		dll := prepareListFindM()
		tests := []testCase{
			{3, 'a'},
			{1, 'b'},
			{7, 'q'},
			{6, 'x'},
			{5, 'j'},
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

	t.Run("Expect method to return -1 if there no such element in the list", func(t *testing.T) {
		dll := prepareListFindM()
		i, err := dll.FindLast('є')
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if i != -1 {
			t.Errorf("Expected -1, got %v", i)
		}
	})

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

	t.Run("Expect method to return an error for non-letter arguments", func(t *testing.T) {
		dll := DoublyLinkedList{}
		_, err := dll.FindLast('+')
		if err == nil {
			t.Errorf("Expected method to return an error, got %v", err)
		}
	})
}

func prepareListForClearM() DoublyLinkedList {
	dll := DoublyLinkedList{}
	dll.Append('f')
	dll.Append('u')
	dll.Append('y')
	dll.Append('k')
	return dll
}

func TestClear(t *testing.T) {
	t.Run("Expect method to delete all elements from the list", func(t *testing.T) {
		dll := prepareListForClearM()
		dll.Clear()
		_, err := dll.Get(0)
		if err == nil {
			t.Errorf("Expected error: list is empty, got %v", err)
		}
	})

	t.Run("Expect method to decrease size of the list to 0", func(t *testing.T) {
		dll := prepareListForClearM()
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

		tests := []testCase{
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
