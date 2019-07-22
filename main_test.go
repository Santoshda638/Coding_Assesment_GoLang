package main

import (
	"testing"
)

func TestTraingleType(t *testing.T) {
	op := TraingleType(1, 1, 1)
	if op != "Equilateral" {
		t.Error("Test failed")
	}
	op = TraingleType(1, 2, 2)
	if op != "Isoceles" {
		t.Error("Test failed")
	}
	op = TraingleType(3, 4, 2)
	if op != "Scalene" {
		t.Error("Test failed")
	}
}

func TestGetNthFromEnd(t *testing.T) {
	a:=[]int{1,2,3,4,5,6}
	head:= ListCreation(a)
	n:= GetNthFromEnd(head,5)
	if n.value!=2{
		t.Error("Test failed")
	}
	b:=[]int{1,2,3,4}
	head2:= ListCreation(b)
	n= GetNthFromEnd(head2,5)
	if n!=nil{
		t.Error("Test failed")
	}
}

func TestTwoListCheck(t *testing.T) {
	l1:= []string{"tom","dick","harry"}
	l2:= []string{"tom","tom","dick","dick"}
	if !TwoListCheck(l1,l2){
		t.Error("Test failed")
	}
	l2= []string{"tom","dick","harry"}
	l1= []string{"alice"}
	if TwoListCheck(l1,l2){
		t.Error("Test failed")
	}
	l2= []string{"tom","dick"}
	l1= []string{"harry","tom","dick","harry"}
	if !TwoListCheck(l1,l2){
		t.Error("Test failed")
	}
}
