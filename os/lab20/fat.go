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

var disk = new(Disk)

type INoder interface {
	create()
	rename(newName string)
	cd()
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
func (iNode *INode) cd() {
	fmt.Println("进入目录: ", "/")
}
func (iNode *INode) show() {
	fmt.Println("显示目录内文件: ", "张云鹏.txt")
}
func (iNode *INode) delete(address int) {
	disk.FAT1[address] = 0
	fmt.Println("     删除文件: ", "file.txt")
}
func (iNode *INode) rename() {
	fmt.Println("    重命名文件: ", "张云鹏.txt", "to file.txt")
}

//获取空闲的磁盘块地址
func getAvailableFatAddress() int {
	for item := range disk.FAT1 {
		if item == 0 {
			return item
		}
	}
	panic("No Available Address")
	return -1
}

//将字节流复制到硬盘
func writeBytesToDisk(sourceBytes []byte) int {
	// var n = 0
	blockNum := len(sourceBytes)/64 + 1
	if len(sourceBytes)%64 == 0 {
		blockNum--
	}
	for i := 0; i < blockNum; i++ {
		availableFatAddress := getAvailableFatAddress()
		diskBlock := disk.DiskData[availableFatAddress : availableFatAddress+64]
		diskBlockBuffer := bytes.NewBuffer(diskBlock)
		limit := Min(len(sourceBytes), (i+1)*64)
		n, err := diskBlockBuffer.Write(sourceBytes[i*64 : limit])
		if err != nil {
			fmt.Println("写入硬盘失败")
		}
		fmt.Println("写入", n, "字节")
	}

	return 0
}

func Redirect(sourceBytes []byte, fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("file open err!")
	}
	file.Write(sourceBytes)
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func sliceSourceBytes(sourceBytes []byte) [][]byte {
	var result [][]byte
	line := len(sourceBytes)/64 + 1
	fmt.Println(line)
	return result
}

type Byter interface {
	toBytes()
}

type fat struct {
	next int
}

func main() {

}
