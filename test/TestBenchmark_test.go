package test

import (
	"testing"
	"fmt"
)

func Benchmark_Aaa(b *testing.B){
	var n int
	for i:=0;i<b.N;i++{
		n++
		fmt.Sprintf("%d",n)
	}
}
