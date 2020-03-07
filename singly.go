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


func insertathead (val int, p *mylist) (*mylist){

	cnode := &mylist{val, nil}
	if(p == nil) {
		p = cnode
	} else {
		cnode.next = p
		p = cnode
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

func deletenode (p *mylist, v int) (pv *mylist){
	head := p
	prev := p

	for p != nil {

		if(p.val == v){
			if(p == head){
				head = p.next
			} else {
				prev.next = p.next
			}
			return head
		}

		prev = p
		p = p.next

	}
	fmt.Println("Item ", v, " not found")
	return head
}

func reverselist (p *mylist) (pv *mylist){
	if p == nil {
		return p
	}

	var head *mylist

	for p != nil {
		cnode := &mylist{p.val, nil}
		
		if head == nil {
			head = cnode
		} else {
			cnode.next = head
			head = cnode
		}
		p = p.next
	}
	return head
}

func main(){
	
	var testlist *mylist

	fmt.Println("Enter number of items")
	var items int
	fmt.Scan(&items)

	for i:=0;i<items;i++ {
		var itemval int
		fmt.Scan(&itemval)
		testlist = insertnode(itemval, testlist)
	}

	fmt.Println("Printing list...")
	printlist(testlist)

	/*fmt.Println("Enter item to delete")
	var ditem int
	fmt.Scan(&ditem)

	testlist = deletenode(testlist, ditem)
	
	fmt.Println("Printing list....")
	printlist(testlist)*/

	testlist = reverselist(testlist)
	
	fmt.Println("Reversed list....")
	printlist(testlist)	

}