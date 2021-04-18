package myintf

import (
	"fmt"
)
//虚回调的实现方法
type BaseIntf interface {
	F1()
	Init(app interface{})
}
type Base1 struct {
	App interface{}
}
func (b *Base1)F1(){
	v,ok:=((interface{})(b.App)).(SonIntf)
	fmt.Println(v,ok)
	v.F2()

}
func (b *Base1)Init(app interface{}){
	b.App=app
}
//============================================
type SonIntf interface {
	F2()
}


type Son struct {
	Base1
}
func (s *Son)F2(){
	fmt.Println("F2")
}

func Myintf3(){
	//方法１
	//var b interface{}=new(Son)
	//b.(BaseIntf).Init(b)
	//b.(BaseIntf).F1()

	//方法２
	//var b Son
	//b.Init(&b)
	//b.F1()

	//以下方式不可以
	//var b Son
	//b.Init(b)
	//b.F1()

	fmt.Println("=====================================")

//类转换成接口的方法，必须是指针，否则无法转换
	var b Son
	v,ok:=((interface{})(b)).(SonIntf)
	if ok{
		fmt.Println("xxxx",v)
	}
	v1,ok1:=((interface{})(&b)).(SonIntf)
	if ok1{
		fmt.Println("yyyy",v1)
	}

	v2,ok2:=((interface{})(b)).(BaseIntf)
	if ok2{
		fmt.Println("zzzz",v2)
	}
	v3,ok3:=((interface{})(&b)).(BaseIntf)
	if ok3{
		fmt.Println("tttt",v3)
	}
}
