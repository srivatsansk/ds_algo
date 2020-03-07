package main

import (
	"fmt"
)

type mylist struct {
    val int
    next *mylist
} 

func insertnode (val int, p *mylist) (*mylist){

	tnode := &mylist{val, nil}

	if p == nil {
		p = tnode		
	} else {
		curnode := p
		for curnode.next != nil {
			curnode = curnode.next
		}
		curnode.next = tnode		
	}
	return p
}


func printlist (p *mylist) {

	if p == nil {
		fmt.Println("List is empty")
		return
	}

	for p != nil {
		fmt.Println(p)
		p = p.next
	}
}

func mergelist(p1 *mylist, p2 *mylist, p3 *mylist) (*mylist) {
	
	for p1 != nil && p2 != nil {
		if p1.val < p2.val {
			p3 = insertnode(p1.val, p3)
			p1 = p1.next
		} else if p1.val >= p2.val {
			p3 = insertnode(p2.val, p3)
			p2 = p2.next
		}
	}
	for p1 != nil {
		p3 = insertnode(p1.val, p3)
		p1 = p1.next
	}
	for p2 != nil {
		p3 = insertnode(p2.val, p3)
		p2 = p2.next
	}

	return p3
}


func main(){
	
	var testlist1 *mylist
	var testlist2 *mylist
	var testlist3 *mylist

	fmt.Println("Enter number of items")
	var items int
	fmt.Scan(&items)

	for i:=0;i<items;i++ {
		var itemval int
		fmt.Scan(&itemval)
		testlist1 = insertnode(itemval, testlist1)
	}

	fmt.Println("Enter number of items")
	fmt.Scan(&items)

	for i:=0;i<items;i++ {
		var itemval int
		fmt.Scan(&itemval)
		testlist2 = insertnode(itemval, testlist2)
	}	

	fmt.Println("Printing 1st list...")
	printlist(testlist1)
	fmt.Println("Printing 2nd list...")
	printlist(testlist2)

	testlist3 = mergelist(testlist1, testlist2, testlist3)

	fmt.Println("Printing merged list...")
	printlist(testlist3)


}