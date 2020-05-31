package main

import "fmt"
type node struct {
  key interface{}
  next *node
}
type list struct {
  len int
  tail *node
}

func main() {
  fmt.Println("Hello World")
  l:=new(list)
  l.appendfront(10)
  l.appendfront(11)
  l.appendfront(12)
  l.appendfront(14)
  l.appendfront(16)
  l.appendmiddle(4, 5)
  l.traverse()
  l.remove(4)

  fmt.Println(l)
  l.traverse()
}

func (l *list) appendfront(value int){
  n:= &node{
    key: value,
  }
  if l.tail == nil {
    l.tail = n
  } else {
    n.next = l.tail
    l.tail = n
  }
  l.len++
}

func (l *list) appendmiddle(value int, pos int) {
  n:=&node{
    key: value,
  }
  if pos>l.len {
    panic("too far, list postion below length of list")
  }
  if l.tail ==nil {
    l.tail = n
  }else {
    currentpos:=0
    currentnode:=l.tail
    for ;currentpos < pos -1 ;currentpos++{
      currentnode=currentnode.next
    }
    n.next=currentnode.next
    currentnode.next = n
  }
  l.len++
}

func (l *list) remove(pos int){
  if pos>=l.len {
    panic("no postion available")
  } else{
    currentpos:=0
    currentnode:=l.tail
    for ;currentpos < pos -1 ;currentpos++{
      currentnode=currentnode.next
    }
    currentnode.key=currentnode.next.key
    currentnode.next=currentnode.next.next
    l.len--
  }
}

func (l *list) traverse(){
  current:=l.tail
  for current!=nil {
    fmt.Println("value:",current.key)
    current=current.next
  }
}
