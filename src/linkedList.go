package main

import (
	"encoding/json"
	"fmt"
)

/*
数组与链表的区别：
1.数组需要一块连续的内存空间来存储；链表不需要一块连续的内存空间，它通过指针将一组零散的内存块串联起来使用
2.链表的插入和删除操作时间复杂度为O(1)，随机访问为O(n);数组的插入和删除操作需要做大量的数据搬移，时间复杂度为O(n)，随机访问为O(1)
3.数组的缺点是大小固定；链表本身没有大小限制，天然支持动态扩容
4.链表需要额外的空间进行存储，所以内存消耗会翻倍；如果你的代码对内存的使用非常苛刻，那么数组将更合适
*/

/*
最常见的三种链表结构：单链表、双向链表、循环链表

1.单链表：
	----> data|next ----> data|next ----> data|next ----> data|next ----> null

链表通过指针将一组零散的内存串联在一起，我们把内存块成为链表的“结点”。为了将所有结点穿起来，每个链表的结点除了存储数据之外，
还需要记录链上的下一个结点地址。我们把这个记录下一个节点地址的指针叫做后继指针next。
我们习惯性的把第一个结点叫做头结点，把最后一个结点叫做尾结点。其中头结点用来记录链表的基地址。而尾结点特殊的地方是，
指针不是指向下一个结点，而是指向一个空地址NULL，表示这是链表上的最后一个结点。

2.循环链表：
循环链表是一种特殊的单链表，它跟链表的唯一区别是在尾结点。循环链表的尾结点指向链表的头结点，犹如一个首尾相连的环。
和单链表相比，循环链表的有点事从链表尾到链表头比较方便。当要处理的数据具有环形结构特点时，就特别适合采用循环链表。

3.双向链表：
	----> prew|data|next <----> prew|data|next <----> prew|data|next <----> prew|data|next

双向链表需要额外使用两个空间进行存储，但是在特定的情况下会更高效：

（1）在某个特定节点前进行插入、删除操作。当我们找到这个特定节点后，单链表需要再次遍历链表，找出满足 目标节点-》next=特定节点 的目标节点
然后进行操作，此时删除操作的时间复杂度为O(n)+O(1)（查找+删除）。而双向链表已经存储了前驱结点，因此总时间复杂度为O(1)。

（2）对于一个有序链表，双向链表的按值查询要比单链表高效。因为我们可以记录上次查找的位置p，每次查询时，根据要查找的值与p的大小关系，决定
往前还是往后查找，所以平均只需要查找一般的数据。

双向链表尽管比较费内存，但在实际的软件开发中应用更广泛。比如java中的LinkedList、LinkedHashMap容器，就用到了双向链表这种数据结构。
这是一种用空间换时间的思想方法。

4.双向循环链表：即把双向链表和循环链表结合的一种数据结构。
*/

// 单链表结点结构
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// 单链表数据结构
type LinkedList struct {
	Head   *ListNode
	Length int
}

// 在某个值后面插入值
func (list *LinkedList) InsertAfter(targetValue, value interface{}) (bool, error) {
	if nil == targetValue {
		return false, fmt.Errorf("target value cannot be nil")
	}
	node := list.Head
	for node != nil {
		if node.Value != targetValue {
			node = node.Next
		} else {
			newNode := ListNode{}
			newNode.Value = value
			newNode.Next = node.Next
			node.Next = &newNode
			return true, nil
		}
	}
	return false, fmt.Errorf("target value is not exist in LinkedList")
}

// 在链表头部插入结点
func (list *LinkedList) InsertFirst(value interface{}) error {
	if value == nil {
		return fmt.Errorf("value is empty")
	}
	firstNode := new(ListNode)
	firstNode.Next = list.Head
	firstNode.Value = value
	list.Head = firstNode
	list.Length = list.Length + 1
	return nil
}

// 在链表尾部插入结点
func (list *LinkedList) InsertLast(value interface{}) error {
	if value == nil {
		return fmt.Errorf("value is empty")
	}
	lastNode := new(ListNode)
	lastNode.Next = nil
	lastNode.Value = value
	ln := list.Head
	for ln.Next != nil {
		ln = ln.Next
	}
	list.Length = list.Length + 1
	ln.Next = lastNode
	return nil
}

// 通过索引查找结点
func (list *LinkedList) GetNode(index int) (interface{}, error) {
	if list.Length < index+1 || index < 0 {
		return nil, fmt.Errorf("invalid value of index")
	}
	node := list.Head
	for ; index-1 >= 0; index-- {
		node = node.Next
	}
	return node.Value, nil
}

// 删除传入的值
func (list *LinkedList) DeleteNode(targetValue interface{}) error {
	if targetValue == nil {
		return fmt.Errorf("value is empty")
	}
	ln := list.Head
	for ln.Next != nil {
		if ln.Next.Value == targetValue {
			ln.Next = ln.Next.Next
			return nil
		}
		ln = ln.Next
	}
	return fmt.Errorf("target value is not exist in LinkedList")
}

//打印链表
func (list *LinkedList) PrintLinkedList() error {
	result := "["
	ln := list.Head
	for ln != nil {
		bytes, err := json.Marshal(ln.Value)
		if err != nil {
			return nil
		}
		result += string(bytes)
		if ln.Next != nil {
			result += ", "
		} else {
			break
		}
		ln = ln.Next
	}
	result += "]"
	fmt.Println(result)
	return nil
}

func main() {
	linkedList := new(LinkedList)
	firstNode := new(ListNode)
	firstNode.Value = "8"
	linkedList.Head = firstNode
	linkedList.Length = 1

	linkedList.InsertFirst("6")
	linkedList.InsertLast(7)
	linkedList.InsertLast("6")
	linkedList.InsertAfter(7, "8")
	linkedList.DeleteNode(7)
	val, _ := linkedList.GetNode(3)
	fmt.Println("index 3 :", val)
	fmt.Println("the list length is", linkedList.Length)
	linkedList.PrintLinkedList()

	ok, err := judgeStr(linkedList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok)
}

/* 基于链表实现LRU（Least Recently Used）缓存淘汰算法
思路：我们维护一个有序单链表，越靠近尾部的结点就是越早之前访问的。当新的数据被访问，我们从链表头开始顺序遍历链表
如果此数据已存在，则删除其原位置，然后插入到链表头部
如果不存在，当此缓存未满，则直接插入链表头部；如果已满，删除尾结点，插入头部

基于这个实现思路，不管缓存满不满，未满都需要遍历一边链表，因此缓存访问的时间复杂度为O(n)
*/

/*
	小练习：如果字符串是通过单链表来存储的，如何判断是否是一个回文字符串？
	思路：使用快慢两个指针找到链表的重点，慢指针每次前进一步，快指针每次前进两步。
	在慢指针前进的过程中，同时修改其next指针，使得链表前半部分反序，最后比较中点两侧的链表是否相等
*/
func judgeStr(list *LinkedList) (bool, error) {
	if list == nil {
		return false, fmt.Errorf("invalid parameter : the given list is empty")
	}
	if list.Head.Next == nil {
		return true, nil
	}

	var prev *ListNode
	fast := list.Head
	slow := list.Head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Value != prev.Value {
			return false, nil
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true, nil
}
