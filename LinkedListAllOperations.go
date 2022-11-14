package main
import (
	"fmt"
	"strings"
)
type node struct {
	data int
	next *node
}
//
type linkidList struct {
	head   *node
	length int
}
//main
func main() {
l:=linkidList{}
l.add(1)
l.add(2)
l.add(3)
l.add(4)
l.add(5)
l.add(6)
l.add(7)
l.add(8)
l.add(9)
l.add(10)
fmt.Println(l)
fmt.Println(*(l.get(5)))
l.remove(4)
l.remove(5)
fmt.Println(l)
}
//add values
func (l *linkidList) add(value int) {
	tail := new(node)
	tail.data = value
	if l.head == nil {
		l.head = tail
	} else {
		in := l.head
		for ; in.next != nil; in = in.next {}
		in.next=tail
	}}
	//string operations
func (l linkidList) String() string {
	sb := strings.Builder{}
	for in:=l.head;in!=nil;in=in.next{
		sb.WriteString(fmt.Sprintf("%d ",in.data))
	}
	return sb.String()
}
//to get values
func (l linkidList) get(value int) *node{
	for in:=l.head;in!=nil;in=in.next{
		if in.data==value{
			return in
		}
	}
	return nil
}
func (n node) String()string{
	return fmt.Sprintf("%d",n.data)
}
//to revome
func (l *linkidList) remove(value int) {
	var previous *node
		for current:=l.head;current!=nil;current=current.next{
		if current.data==value {
			previous.next=current.next
			return
		}
		previous=current
	}
}