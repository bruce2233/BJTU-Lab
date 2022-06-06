package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

const (
	DISK_SIZE   = 1440 * 1024
	SECTOR_SIZE = 512
)

type DiskInitAndExport interface {
	Init()
	Export()
}

type Disk struct {
	DiskData []byte
	MBR      []byte
	FAT1     []int
	FAT2     []int
	RootDir  []byte
	FileData []byte
}

func (disk *Disk) Init() {
	disk.DiskData = make([]byte, DISK_SIZE, DISK_SIZE)               //磁盘大小1474560字节
	disk.MBR = disk.DiskData[0:SECTOR_SIZE]                          //引导扇区
	disk.FAT1 = disk.DiskData[SECTOR_SIZE : 10*SECTOR_SIZE]          //FAT1
	disk.FAT2 = disk.DiskData[10*SECTOR_SIZE : 19*SECTOR_SIZE]       //FAT2
	disk.RootDir = disk.DiskData[20*SECTOR_SIZE : 33*SECTOR_SIZE]    //根目录
	disk.FileData = disk.DiskData[33*SECTOR_SIZE : 2879*SECTOR_SIZE] //数据区
	fmt.Println("磁盘初始化完成.")
}
func (disk *Disk) Export() {
	imgFileName := "img"
	os.Create(imgFileName)
	file, err := os.OpenFile(imgFileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open file error")
	}
	file.Write(disk.DiskData) //写入字节流到文件, 导出镜像文件
	fmt.Println("导出映像文件完成.")
}

var disk = new(Disk)

type INoder interface {
	create()
	rename(newName string)
	show()
	delete()
}

type INode struct {
	name       string
	createTime int64
	accessTime int64
	modifyTime int64
	size       int64
	fatHead    int
	fatTail    int
	fileBytes  []byte
}

func (iNode *INode) toBytes() []byte {
	iNodeBytes := make([]byte, 0, 64)
	nameBytes := []byte(iNode.name)
	nameBytes = nameBytes[:16] //名称保留前16字节
	iNodeBytes = append(iNodeBytes, nameBytes...)
	fmt.Println("文件名称: ", string(nameBytes))

	t := new(bytes.Buffer)
	binary.Write(t, binary.LittleEndian, iNode.createTime) //创建时间8字节, 小端序
	iNodeBytes = append(iNodeBytes, t.Bytes()...)
	fmt.Println("创建时间: ", time.Unix(iNode.createTime, 0))

	binary.Write(t, binary.LittleEndian, iNode.accessTime) //访问时间8字节, 小端序
	iNodeBytes = append(iNodeBytes, t.Bytes()...)
	fmt.Println("最近访问时间: ", time.Unix(iNode.accessTime, 0))

	binary.Write(t, binary.LittleEndian, iNode.modifyTime) //修改时间8字节, 小端序
	iNodeBytes = append(iNodeBytes, t.Bytes()...)
	fmt.Println("最近修改时间: ", time.Unix(iNode.modifyTime, 0))

	binary.Write(t, binary.LittleEndian, iNode.size) //文件大小8字节, 小端序, 定位Byte
	iNodeBytes = append(iNodeBytes, t.Bytes()...)
	fmt.Println("文件大小: ", iNode.size)

	return iNodeBytes
}

func (iNode *INode) create(newNode *INode) {
	availableBlockAddress := getAvailableFatAddress()
	iNode.fatTail = availableBlockAddress
	fmt.Println("创建文件/目录, iNode保存fat地址: ", iNode.fatTail)
	fmt.Println("创建文件/目录, 更新内容: ", iNode.fatTail)
	copy(disk.DiskData[availableBlockAddress:availableBlockAddress+64], newNode.toBytes())
	if iNode.size != 0 {

	}
}

//获取空闲的磁盘块地址
func getAvailableFatAddress() int {
	for item := range disk.FAT1 {
		if item == 0 {
			return item
		}
	}
	return -1
}

type Byter interface {
	toBytes()
}

type fat struct {
	next int
}

func main() {

}
