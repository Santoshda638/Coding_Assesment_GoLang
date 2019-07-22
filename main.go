package main

import "fmt"

// === Problem 1 ===
func TraingleType(ab , bc ,ca  float64) string{
	if !(ab + bc <= ca || ab + ca <= bc || bc + ca <= ab) && ab>0 && bc>0 && ca>0  {
	if ab == bc && bc == ca {
		return "Equilateral"
	} else if ab == bc || bc == ca || ca == ab {
		return "Isoceles"
	} else {
		return "Scalene"
	}
	}else{
		return "Not a traingle"
	}
}
// === Problem 1 ===


// === Problem 2 ===
type LinkList struct {
	value     int
	next      *LinkList
}
func ListCreation(arr []int) *LinkList{
	var last *LinkList
	var head *LinkList
	for i,v:=range arr{
		node:=&LinkList{}
		node.value = v
		if i==0{
			head = node
			last = node
		}

		last.next = node

		last = node

	}
	return head
}

func GetNthFromEnd(head *LinkList, n int) *LinkList{
	fifth := head
	begin := head
	nCounter := 0
	if head == nil {
		// head nil empty list
		return head
	}else{
		// move forward by n
		for ; nCounter < n; {
			if begin == nil {
				return nil;
			}
			nCounter = nCounter +1;
			begin = begin.next;
		}
		//move forward until end of list
		for ;begin != nil; {
			fifth = fifth.next;
			begin  = begin.next;
		}
		return fifth
	}

}

// === Problem 2 ===


// === Problem 3 ===
func TwoListCheck(l1,l2 []string) bool{
	flag:=true
	//create sets
	set1:=make(map[string]bool)
	set2:=make(map[string]bool)
	for _,v:=range l1{
		set1[v]=true
	}
	for _,v:=range l2{
		set2[v]=true
	}
	// check if l1 in l2
	for key,_ := range set2{
		if !set1[key]{
			flag=false
			break
		}
	}
	if flag{
		return flag
	}
	flag = true
	// check if l2 in l1
	for key,_ := range set1{
		if !set2[key]{
			flag=false
			break
		}
	}
	return flag
}

// === Problem 3 ===

func main(){
	// problem1
	op := TraingleType(1, 1, 1)
	fmt.Println("Triangle is ", op)
	// problem2
	a:=[]int{1,2,3,4,5,6}
	head:= ListCreation(a)
	n:= GetNthFromEnd(head,5)
	fmt.Println("Fifth element - " ,n.value)
	// problem3
	l1:= []string{"tom","dick","harry"}
	l2:= []string{"tom","tom","dick","dick"}
	fmt.Println( "List1 contain List2 ? ",TwoListCheck(l1,l2))

}