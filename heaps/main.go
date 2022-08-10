package main

import "fmt"

type maxheap struct{
	array []int
}

func(h *maxheap)insert(key int){
	h.array=append(h.array, key)
	h.maxheapifyup(len(h.array)-1)

}
func (h *maxheap) maxheapifyup(index int){
	for h.array[parent(index)]<h.array[index]{
		h.swap(parent(index),index)
		index=parent(index)
	}
}

func parent(i int)int{
	return(i-1)/2
}
func left(i int)int{
	return 2*i+1
}
func right(i int)int{
	return 2*i+2 
}
func (h *maxheap)swap(i1,i2 int){
	h.array[i1],h.array[i2]=h.array[i2],h.array[i1]
}
func (h *maxheap)extract()int{
	extracted:= h.array[0]
	if len(h.array)==0{
		fmt.Println("array length is 0")
		return -1
	}
	h.array[0]=h.array[len(h.array)-1]
	h.array=h.array[:len(h.array)-1]
	h.maxheapifydown(0)

	return extracted
}
func (h *maxheap)maxheapifydown(index int){
	lastIntex:=len(h.array)-1
	l,r:=left(index),right(index)
	childToCompare:=0
	for l<=lastIntex{
		if l==lastIntex {
			childToCompare=l
		}else if h.array[l]>h.array[r]{
			childToCompare=l
		}else{
			childToCompare=r
		}
		if h.array[index]<h.array[childToCompare]{
			h.swap(index,childToCompare)
			index=childToCompare
			l,r=left(index),right(index)
		}else{
			return
		}
	}


}

func main(){
	m:=&maxheap{}
	fmt.Println(m)
	buldheap:=[]int {20,10,3,6,1,80,56,43,567,8,34,9}
	for _,v:=range buldheap{
		m.insert(v)
		fmt.Println(m)
	}
	for i:=0;i<5;i++{
		fmt.Println(m.extract())
		fmt.Println(m)
	}
	
}