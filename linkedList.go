package main 

import "fmt"

func main (){

	var l = LinkedList{}
	var p Person= Person{"ddd",3} 

	l.insert(&p)
	l.insert(&Person{"John", 19})
	l.insert(&Person{"John", 19})
	l.insert(&Person{"John", 19})
	l.insert(&Person{"John", 19})

	l.remove()
	
	l.display()

	
}


type Person struct {
	name string
	age int
}

var p = Person{"",3}

type Node struct {
	data *Person
	next *Node
}

type LinkedList struct {
	head *Node
	current *Node
}

func (l *LinkedList) isEmpty() bool{
	return l.head == nil
} 

func (l *LinkedList) findNext(){
	l.current = l.current.next
}

func (l *LinkedList) insert(val *Person){
	if l.head == nil {
		l.current = &Node{val,nil}	
		l.head = l.current
	} else {
		tmp:= l.current.next
		
		n:= Node{val, tmp}
		
		l.current.next = &n

		l.current = l.current.next
		
	}

}

func (l *LinkedList) remove() {
	if l.current == l.head {
		l.head = l.head.next
	}else {
        	temp:= l.head
        	for temp.next != l.current {
                	temp = temp.next
        	}
		temp.next = l.current.next
	}
	if l.current.next == nil { // single node list
		l.current = l.head 
	}else{
		l.current = l.current.next
	}

	
}

func (l *LinkedList) display(){
	p:= l.head
	for p != nil {
		fmt.Printf(" %v -> ", p.data)
		
		p = p.next
	}
	fmt.Println()
}


