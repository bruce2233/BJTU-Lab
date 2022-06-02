package main

const (
	DISK_SIZE = 1440*1024
	SECTOR_SIZE = 512
)
type DiskInitAndExport interface{
	Init()
	Export()
}
type Disk struct{
	DiskData []byte
	MBR []byte
	FAT1 []byte 
	FAT2 []byte 
	RootDir []byte 
	FileData []byte
}
func (disk *Disk)Init(){
	disk.DiskData = make([]byte, DISK_SIZE, DISK_SIZE) //磁盘大小1474560字节
	disk.MBR = disk.DiskData[0:SECTOR_SIZE]//引导扇区
	disk.FAT1 = disk.DiskData[SECTOR_SIZE: 10*SECTOR_SIZE] //FAT1
	disk.FAT2 = disk.DiskData[10* SECTOR_SIZE: 19*SECTOR_SIZE]//FAT2
	disk.RootDir = disk.DiskData[20* SECTOR_SIZE,33* SECTOR_SIZE] //根目录
	disk.FileData = disk.DiskData[33* SECTOR_SIZE,2879* SECTOR_SIZE] //数据区
	fmt.Println()
	
}

type iNode struct{

}

func main(){

}

