package main

import (
	"fmt"
	"golab/structure/linelist"
	queue2 "golab/structure/queue"
	stack2 "golab/structure/stack"
	"golab/structure/tree"
)

func main() {

	//----------------------------------------------二叉树的操作--------------------------------------------//
	root := tree.NewBinaryTree(1)
	root.Left = tree.NewBinaryTree(2)
	root.Right = tree.NewBinaryTree(3)
	root.Left.Left = tree.NewBinaryTree(4)
	root.Left.Right = tree.NewBinaryTree(5)
	root.Right.Left = tree.NewBinaryTree(6)
	root.Right.Right = tree.NewBinaryTree(7)

	fmt.Println("前序遍历")
	fmt.Println(root.PreOrder())
	fmt.Println("后序遍历")
	fmt.Println(root.AfterOrder())
	fmt.Println("中序遍历")
	fmt.Println(root.MiddleOrder())
	fmt.Println("层序遍历")
	fmt.Println(root.LayersOrder())

	//----------------------------------------------线性表的操作------------------------------------------------------//
	list := linelist.InitLineList()

	fmt.Println(list.Empty())
	fmt.Println(list.Append(1))
	fmt.Println(list.Append(2))
	fmt.Println(list.Append(3))
	fmt.Println(list.Append(4))
	fmt.Println(list.Insert(2, 10))
	fmt.Println(list.LineListContent)
	fmt.Println(list.MaxLength)
	fmt.Println(list.Length)
	fmt.Println(list.DelElem(3))
	fmt.Println(list.LineListContent)
	fmt.Println(list.Pop())
	fmt.Println(list.LineListContent)

	//---------------------------------------------单链表操作-------------------------------------------------------------------//
	single := linelist.NewSingleList()
	fmt.Println("添加结点")
	fmt.Println(single.Add(1))
	fmt.Println(single.Add(2))
	fmt.Println(single.Add(3))
	fmt.Println("获取第一个结点")
	fmt.Println(single.GetFirst())
	fmt.Println("获取最后一个结点")
	fmt.Println(single.GetLast())
	fmt.Println("获取单链表的长度")
	fmt.Println(single.Length())
	fmt.Println("获取单链表的元素")
	fmt.Println(single.GetAll())
	fmt.Println(single.GetElem(1))
	fmt.Println("删除结点")
	fmt.Println(single.Delete(2))
	fmt.Println("获取单链表的元素")
	fmt.Println(single.GetAll())

	//----------------------------------------------双向链表操作-------------------------------------------------------------------//
	double := linelist.InitDoubleList()

	double.AddHeadNode(linelist.InitDoubleNode(1))
	double.AddHeadNode(linelist.InitDoubleNode(2))
	double.AddHeadNode(linelist.InitDoubleNode(3))
	double.AddHeadNode(linelist.InitDoubleNode(4))
	fmt.Println("获取双向链表的所有元素")
	fmt.Println(double.GetAll())
	fmt.Println(double.GetOrder(2))
	fmt.Println(double.GetOrder(3))
	fmt.Println("插入双向链表")
	fmt.Println(double.InsertNextInt(1, linelist.InitDoubleNode(5)))
	fmt.Println(double.GetAll())
	fmt.Println(double.GetReverse(1))
	fmt.Println(double.GetAll())
	fmt.Println(double.Remove(1))
	fmt.Println(double.GetAll())

	//----------------------------------------循环链表操作--------------------------------------------------------------------------------------//
	circle := linelist.InitCircleList()

	fmt.Println(circle.Append(linelist.InitCircleNode(1)))
	fmt.Println(circle.Append(linelist.InitCircleNode(2)))
	fmt.Println(circle.Append(linelist.InitCircleNode(3)))
	fmt.Println(circle.Append(linelist.InitCircleNode(4)))

	fmt.Println(circle.Insert(2, linelist.InitCircleNode(5)))
	fmt.Println(circle.Insert(2, linelist.InitCircleNode(6)))

	fmt.Println(circle.GetAll())
	fmt.Println(circle.Get(3))
	fmt.Println(circle.RemoveInt(2))
	fmt.Println(circle.GetAll())

	//-----------------------------------------栈--------------------------------//
	stack := stack2.InitStackEntry()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Size())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Clear())
	fmt.Println(stack)

	//-----------------------------------------队列--------------------------------//
	queue := queue2.InitQueueEntry()
	fmt.Println(queue.IsEmpty())
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Len())
}
