package main

import "fmt"

//链表节点
type LinkTableNode struct {
	data interface{}
	next *LinkTableNode
}

//链表
type LinkTable struct {
	head  *LinkTableNode
	tail  *LinkTableNode
	SumOfNode int
}

//创建链表
func CreateLinkTable() *LinkTable{
	linkTable := new(LinkTable)
	linkTable.head = nil
	linkTable.tail = nil
	linkTable.SumOfNode = 0
	return linkTable
}

//获取链表头结点
func getHeadNode(linkTable *LinkTable) *LinkTableNode{
	return linkTable.head
}

//获取链表尾结点
func getTailNode(linkTable *LinkTable) *LinkTableNode{
	return linkTable.tail
}

//获取链表中节点数目
func getSumOfNode(linkTable *LinkTable) int{
	return linkTable.SumOfNode
}

//在链表尾部插入一个节点
func AddLinkTableNode(data interface{},linkTable *LinkTable) {
	node := &LinkTableNode{data,nil}
	if linkTable.head == nil {
		linkTable.head = node
		linkTable.tail = node
	} else {
		linkTable.tail.next = node
		linkTable.tail = node
	}
	linkTable.SumOfNode++
	fmt.Println("插入节点成功，链表如下：")
	showLinkTable(linkTable)
}

//删除指定位置的节点
func DeleteLinkTableNode(index int,linkTable *LinkTable) {
	if index <= 0 || index > linkTable.SumOfNode {
		fmt.Println("输入位置不合法！")
		return
	}
	curNode := linkTable.head
	for i := 0; i < index-1; i++ {
		curNode = curNode.next
	}
	curNode.next = curNode.next.next
	linkTable.SumOfNode--
	fmt.Println("删除节点成功，链表如下：")
	showLinkTable(linkTable)
}

//打印链表
func showLinkTable(linkTable *LinkTable) {
	curNode := linkTable.head
	for curNode != nil {
		fmt.Print(curNode.data)
		fmt.Print(" ")
		curNode = curNode.next
	}
	fmt.Println()
}

//得到下一个结点
func GetNextLinkTableNode(linkTable *LinkTable,linkTableNode *LinkTableNode) *LinkTableNode{
	return linkTableNode.next
}

//删除整个链表
func DeleteLinkTable(linkTable *LinkTable){
	if linkTable != nil{
		linkTable.head = nil
		linkTable.tail = nil
		linkTable.SumOfNode = 0
	}
	fmt.Println("删除链表成功！")
}

// func main(){
// 	linkTable := CreateLinkTable()
// 	AddLinkTableNode(2, linkTable)
// 	AddLinkTableNode(4, linkTable)
// 	AddLinkTableNode(1, linkTable)
// 	AddLinkTableNode(9, linkTable)

// 	head := getHeadNode(linkTable)
// 	fmt.Print("头结点：")
// 	fmt.Println(head.data)

// 	tail := getTailNode(linkTable)
// 	fmt.Print("尾结点：")
// 	fmt.Println(tail.data)

// 	sum := getSumOfNode(linkTable)
// 	fmt.Print("链表中节点个数：")
// 	fmt.Println(sum)

// 	nextNode := GetNextLinkTableNode(linkTable, head)
// 	fmt.Print("头结点的下一个节点：")
// 	fmt.Println(nextNode.data)

// 	DeleteLinkTableNode(3,linkTable)

// 	DeleteLinkTable(linkTable)
// 	showLinkTable(linkTable)
// }