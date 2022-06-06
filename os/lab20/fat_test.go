package main

import (
	"testing"
	"time"
)

func TestFat(t *testing.T) {
	iNode := INode{
		name:       "19281030",
		createTime: time.Now().Unix(),
		accessTime: time.Now().Unix(),
		modifyTime: time.Now().Unix(),
		size:       0,
	}
	iNode.toBytes()

	
}
