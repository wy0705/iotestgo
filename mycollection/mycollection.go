package mycollection

import (
	"fmt"
	"unsafe"
	"sync"
	"container/list"
)

func Mycollection(){
	//Array

	var a [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(a)
	for i,v:=range a{
		fmt.Println(i,v)
		v=888
	}
	fmt.Println(a)

	for i,v:=range a{
		fmt.Println(i,v)
		a[i]=888
	}
	fmt.Println(a)

	fmt.Println("======")

	var b =[...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(b)

	var c=new([10]int)
	for i:=1;i<11;i++{
		c[i-1]=i
	}
	fmt.Println(c,*c)

	fmt.Println("======")

	var ptr *[10]int=c

	fmt.Println(ptr)
	fmt.Println(ptr[1],c[1],(*c)[1],(*ptr)[1])
	//fmt.Println(*(ptr+1)) //这个不允许，go有指针，但是不同的是，go的指针不能做指针运算

	fmt.Println("======")

	var ptrnum =new([10]*int)
	for i:=0;i<10;i++{
		ptrnum[i]=&a[i]
	}
	fmt.Println(ptrnum,*ptrnum)
	fmt.Println(*ptrnum[1],*(*ptrnum)[1])

	fmt.Println("===可以用变态的黑魔法实现指针运算===")
	var aa [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
	//unsafe.Pointer相当于void* uintptr很牛逼，实质是个整数，但是可以表示指针，因此这个东西充当了整数与指针的转换桥梁
	var myvar int64=int64(uintptr(unsafe.Pointer(&aa[0])))
	fmt.Println(myvar,myvar+1)

	var myvarptr uintptr=uintptr(myvar+8) //相当于指针从０位置加１到１位置了
	var myvarptra *int=(*int)(unsafe.Pointer(myvarptr))
	fmt.Println(myvarptra,*myvarptra)


	fmt.Println("========================1=======================")
	//========================================================================
	//Silence
	var a1 []int=make([]int,3,10) //长度3　容量10
	fmt.Println(a1,len(a1),cap(a1))

	fmt.Println("===1===")

	var b1=[10]int{1,2,3,4,5,6,7,8,9,10}
	ab1:=b1[1:5]
	ab2:=b1[2:6]
	ab1[1]=999
	fmt.Println(ab1,ab2,b1)

	fmt.Println("===2===")

	ab1=append(ab1,1,2,3,4,5,6,7,8,9,10)
	ab1[1]=888
	fmt.Println(ab1,ab2,b1)

	fmt.Println("===3===")

	var c1 =new([]int)
	fmt.Println(c1)
	//c1=append(c1,1,2,3,4,5)　//由于c1是指向[]int的指针，然而没有调用make的情况下，系统不可能给他一个数组去指，因此更谈不上追加，因为他是野指针
	c1=&ab1
	var tmp []int
	tmp=append(*c1,1,2,3,4,5)
	c1=&tmp
	fmt.Println(c1)

	fmt.Println("===4===")

	var d1=[]int{1,2,3,4,5}
	var e1=[]int{10,20,30}
	copy(d1,e1)
	fmt.Println(d1,e1)

	var d11=[]int{1,2,3,4,5}
	var e11=[]int{10,20,30}
	copy(e11,d11)
	fmt.Println(d11,e11)

	fmt.Println("========================2=======================")
	//========================================================================
	var a2 map[int]string=make(map[int]string)
	a2[0]="aaa"
	fmt.Println(a2)

	fmt.Println("======")

	var b2=map[int]string{0:"aaa",1:"bbb",2:"ccc"}
	for k,v:=range b2{
		fmt.Println(k,v)
		b2[k]="ok"
	}
	fmt.Println(b2)

	fmt.Println("======")

	var c2 =make(map[int]map[int]string)
	c2[0]=make(map[int]string)
	c2[0][0]="aaa"
	c2[0][1]="bbb"
	c2[1]=make(map[int]string)
	c2[1][0]="ccc"
	c2[1][1]="ddd"
	fmt.Println(c2)

	fmt.Println("======")

	var d2=map[int]string{0:"aaa",1:"bbb",2:"ccc"}
	var e2=new(map[int]string)
	e2=&d2
	fmt.Println(e2)

	fmt.Println("========================3=======================")
	//========================================================================

	var mylist *list.List=list.New()
	mylist.PushBack("aaa")
	mylist.PushFront("bbb")
	mylist.PushFront("ccc")
	ele:=mylist.PushBack("eee")
	mylist.PushBack("fff")
	mylist.Remove(ele)
	for i:=mylist.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}


	fmt.Println("========================4=======================")
	//========================================================================
	var smap sync.Map //相当于java的并发包
	smap.Store("x1","ccc")
	smap.Store("x2","ddd")
	smap.Range(func(k,v interface{})bool{
		fmt.Println(k,v)
		return true
	})


}
