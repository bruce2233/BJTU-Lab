package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFat(t *testing.T) {
	disk.Init()
	disk.Export()
	iNode := INode{
		name:       "19281030",
		createTime: time.Now().Unix(),
		accessTime: time.Now().Unix(),
		modifyTime: time.Now().Unix(),
		size:       0,
	}
	file2 := INode{
		name:       "张云鹏.txt",
		createTime: time.Now().Unix(),
		accessTime: time.Now().Unix(),
		modifyTime: time.Now().Unix(),
		size:       18,
		fileBytes:  []byte("This's a txt file."),
	}
	iNode.toBytes()
	writeBytesToDisk(file2.fileBytes)
	iNode.show()
	iNode.rename()
	iNode.delete()
}

func TestSlice(t *testing.T) {
	a := []int{9}
	t.Log(a)
	changeSlice(a)
	t.Log(a)
}
func changeSlice(slice []int) {
	slice[0] = 1
	fmt.Println(slice)
}

func TestWrite(t *testing.T) {
	disk.Init()
	sourceBytes := []byte{'a'}
	// t.Log(disk)
	// writeBytesToDisk(sourceBytes)
	sliceSourceBytes(sourceBytes)
	// t.Log(sourceBytes[:100])
}
