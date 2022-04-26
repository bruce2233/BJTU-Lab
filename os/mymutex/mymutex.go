package main

import (
	"fmt"
	"time"
)

var glb, glb2 int = 0, 0
var work, work2 int = 0, 0
var count, count2 int = 0, 0
var a, b, c int = 0, 0, 0
var theadNum = 2
var useTime = 100 * time.Millisecond

func mySayA(threadNum int) {
	mytime := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		// fmt.Println(i,"mySayA")
		if threadNum == 0 {
			a = 1
			c = 0
		} else if threadNum == 1 {
			b = 1
			c = 1
		}
		if threadNum == 0 {
			for c == threadNum && b == 1 {
			}
		}
		if threadNum == 1 {
			for c == threadNum && b == 1 {
			}
		}
		glb++ //占用开始
		work++
		time.Sleep(useTime)
		if glb%2 != 1 { //若共享期间修改数据则count++
			count++
		}
		glb-- //占用结束
		if threadNum == 0 {
			a =0
		}else if threadNum == 1{
			b = 0
		}
	}
	mytime = time.Now().UnixNano() - mytime
	fmt.Println(threadNum,"耗时: ", mytime)
}

func originsay() {
	// fmt.Println("new originsay thread")
	mytime := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		glb2++ //占用开始
		work2++
		time.Sleep(1 * time.Nanosecond)
		if glb2% 2 != 1 {
			count2++
		}
		glb2--
	}
	mytime = time.Now().UnixNano() - mytime
	fmt.Println("origin耗时: ", mytime)

}
func main2() {
	go mySayA(0)
	go mySayA(1)
	// time.Sleep(1 * time.Second)
	go originsay()
	// time.Sleep(10 * time.Millisecond)
	go originsay()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("work: ", work)
	fmt.Println("work2: ", work2)
	fmt.Println("count: ", count)
	fmt.Println("count2: ", count2)
}
