package main

import (
	"fmt"
)

type DoubleListNode struct {
	Val map[int]int
	Next *DoubleListNode
	Prev *DoubleListNode
}

type LRUCache struct {
	ValHead *DoubleListNode
	ValTail *DoubleListNode
	Map map[int]*DoubleListNode
	Len int
	Cap int
}


func Constructor(capacity int) LRUCache {
	head := new(DoubleListNode)
	return LRUCache {
		ValHead: head,
		ValTail: head,
		Map: make(map[int]*DoubleListNode),
		Len: 0,
		Cap: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if val, ok := this.Map[key]; ok {
		if val != this.ValHead.Next {
			if val == this.ValTail {
				this.ValTail = this.ValTail.Prev
			}
			val.Prev.Next = val.Next
			if val.Next != nil {
				val.Next.Prev = val.Prev
			}
			val.Next = this.ValHead.Next
			val.Prev = this.ValHead
			this.ValHead.Next = val
		}
		return val.Val[key]
	}else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if val, ok := this.Map[key]; ok {
		if val != this.ValHead.Next {
			if val == this.ValTail {
				this.ValTail = this.ValTail.Prev
			}
			val.Prev.Next = val.Next
			if val.Next != nil {
				val.Next.Prev = val.Prev
			}
			val.Next = this.ValHead.Next
			val.Prev = this.ValHead
			this.ValHead.Next = val
		}
		return
	}
	p := new(DoubleListNode)
	p.Val = map[int]int{key: value}
	this.Map[key] = p
	if this.ValHead.Next == nil {
		p.Prev = this.ValHead
		this.ValHead.Next = p
		this.ValTail = p
		this.Len++
		return
	}
	p.Next = this.ValHead.Next
	this.ValHead.Prev = p
	p.Prev = this.ValHead
	this.ValHead.Next = p
	this.Len++
	if this.Len > this.Cap {
		delete(this.Map, this.ValTail.Val[])
		this.ValTail = this.ValTail.Prev
		this.ValTail.Next = nil
	}
}

func main() {
	//config.Init()
	//r := routes.NewRouter()
	//_ = r.Run(config.HttpPort)
	obj := Constructor(10)
	obj.Put(2,2)
	param_1 := obj.Get(2)
	fmt.Println(param_1)
}
