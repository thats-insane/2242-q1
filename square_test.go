package main

import "testing"

func TestSquare(t *testing.T) {
	sTPeri1, sTArea1 := Square(2)
	if sTPeri1 != 8 {
		t.Errorf("%v got, %v expected, %v recieved.", 2, 8, sTPeri1)
	}
	if sTArea1 != 4 {
		t.Errorf("%v got, %v expected, %v recieved.", 2, 4, sTArea1)
	}

	sTPeri2, sTArea2 := Square(4)
	if sTPeri2 != 16 {
		t.Errorf("%v got, %v expected, %v recieved.", 4, 16, sTPeri2)
	}
	if sTArea2 != 16 {
		t.Errorf("%v got, %v expected, %v recieved.", 4, 16, sTArea2)
	}

	sTPeri3, sTArea3 := Square(591.203)
	if sTPeri3 != 2364.812 {
		t.Errorf("%v got, %v expected, %v recieved.", 591.203, 2364.812, sTPeri3)
	}
	if sTArea3 != 349520.987209 {
		t.Errorf("%v got, %v expected, %v recieved.", 591.203, 349520.987209, sTArea3)
	}
}
