package skiplist

import (
	"fmt"
	"testing"
)

type IntElement int

func (i IntElement) Compare(other interface{}) int {
	return int(i)  - int(other.(IntElement))
}


func TestIsEmpty(t *testing.T) {
	s := Generate(15)
	fmt.Println("len(15)", s.isEmpty())
}

func TestSkipList_Insert(t *testing.T) {
	s := Generate(15)
	s.Insert(IntElement(1))
	s.Insert(IntElement(5))
	s.Insert(IntElement(8))
	fmt.Println(s)
}

func TestSkipList_Find(t *testing.T) {
	s := Generate(15)
	for i := 0; i < 7000; i ++ {
		s.Insert(IntElement(i))
	}
	s.Delete(IntElement(10))

	for i := 0; i < 700; i ++ {
		found := s.Find(IntElement(i))
		if found != nil {
			fmt.Println("found", found.value)
		}
	}
}

func TestLevel(t *testing.T) {
	s := Generate(15)
	for i:=0; i < 1000; i ++ {
		fmt.Println(s.level())
	}
}