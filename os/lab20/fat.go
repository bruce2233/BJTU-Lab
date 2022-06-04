package main

import (
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
	FAT1     []byte
	FAT2     []byte
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

type INoder interface {
	create()
	rename(newName string)
	show()
	delete()
}
type INode struct {
	name       string
	createTime string
	size       int
	fatHead    int
}

func (iNode *INode) toBytes() {

	nameBytes := make([]byte, 8, 8) //名称转化为字节
	nameBytes = []byte()
	myTime := time.Time{}
}

type Byter interface {
	toBytes(byter *Byter)
}

type fat struct {
	next int
}

func main() {

}
