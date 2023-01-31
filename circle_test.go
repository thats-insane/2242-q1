package main

import "testing"

func TestCircle(t *testing.T) {
	cTest1 := Circle{
		Radius: 2,
	}

	if cTest1.Perimeter() != 12.56 {
		t.Errorf("%v got, %v expected, %v recieved.", 2, 12.56, cTest1.Perimeter())
	}
	if cTest1.Area() != 12.56 {
		t.Errorf("%v got, %v expected, %v recieved.", 2, 12.56, cTest1.Area())
	}

	cTest2 := Circle{
		Radius: 12,
	}
	if cTest2.Perimeter() != 75.36 {
		t.Errorf("%v got, %v expected, %v recieved.", 12, 75.36, cTest2.Perimeter())
	}
	if cTest2.Area() != 452.16 {
		t.Errorf("%v got, %v expected, %v recieved.", 12, 452.16, cTest2.Area())
	}

	cTest3 := Circle{
		Radius: 45,
	}
	if cTest3.Perimeter() != 282.6 {
		t.Errorf("%v got, %v expected, %v recieved.", 45, 282.6, cTest3.Perimeter())
	}
	if cTest3.Area() != 6358.5 {
		t.Errorf("%v got, %v expected, %v recieved.", 45, 6358.5, cTest3.Area())
	}
}
