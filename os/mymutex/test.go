package main

import (
	"fmt"
	"time"
)

var glb, glb2 int = 0, 0
var work, work2 int = 0, 0
var count, count2 int = 0, 0
var astate, bstate int = 3, 6
var theadNum = 2
var useTime = 100 * time.Millisecond

func mySayA() {
	mytime := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		// fmt.Println(i,"mySayA")
		astate = 1
		for bstate!=6{}
		astate = 2


	}
	mytime = time.Now().UnixNano() - mytime
	fmt.Println("mySayA耗时: ", mytime)
}

func mySayB() {
	mytime := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		if astate == 3 && bstate == 6 {
			bstate = 4
			for astate != 3 {
			}
			bstate = 5
			glb++ //占用开始
			work++
			time.Sleep(useTime)
			if glb%2 != 1 {
				count++
			}
			glb-- //占用结束
			bstate = 6
		}
	}
	mytime = time.Now().UnixNano() - mytime
	fmt.Println("mySayB耗时:  ", mytime)

}

func originsay() {
	// fmt.Println("new originsay thread")
	mytime := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		glb2++ //占用开始
		work2++
		time.Sleep(1 * time.Nanosecond)
		if glb2%2 != 1 {
			count2++
		}
		glb2--
	}
	mytime = time.Now().UnixNano() - mytime
	fmt.Println("origin耗时: ", mytime)

}

func unisay(process int){

}
func main() {
	go mySayA()
	// time.Sleep(1 * time.Second)
	go mySayB()
	go originsay()
	// time.Sleep(10 * time.Millisecond)
	go originsay()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("work: ",work)
	fmt.Println("work2: ",work2)
	fmt.Println("count: ",count)
	fmt.Println("count2: ",count2)
}
