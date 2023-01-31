package main

import "testing"

func TestTriangle(t *testing.T) {
	cTest1 := Triangle{
		Base:   3,
		Height: 4,
	}

	if cTest1.Perimeter() != 12 {
		t.Errorf("%v got, %v expected, %v recieved.", cTest1, 12, cTest1.Perimeter())
	}
	if cTest1.Area() != 6 {
		t.Errorf("%v got, %v expected, %v recieved.", cTest1, 6, cTest1.Area())
	}

	cTest2 := Triangle{
		Base:   20,
		Height: 15,
	}
	if cTest2.Perimeter() != 60 {
		t.Errorf("%v got, %v expected, %v recieved.", cTest2, 60, cTest2.Perimeter())
	}
	if cTest2.Area() != 150 {
		t.Errorf("%v got, %v expected, %v recieved.", cTest2, 150, cTest2.Area())
	}

	cTest3 := Triangle{
		Base:   198.01,
		Height: 99.11,
	}
	if cTest3.Perimeter() != 518.548887456 {
		t.Errorf("%v got, %v expected, %v recieved.", cTest3, 518.548887456, cTest3.Perimeter())
	}
	if cTest3.Area() != 9812.38555 {
		t.Errorf("%v got, %v expected, %v recieved.", cTest3, 9812.38555, cTest3.Area())
	}
}
