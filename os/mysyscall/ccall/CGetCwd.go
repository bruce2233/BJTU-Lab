package ccall

/*
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
char *buf;
void cGetCwd(){
	char * a;
	buf = (char*)malloc(100 * sizeof(char));
    a = getcwd(buf,100);
}
*/
import "C"
import (
	"unsafe"
)

func CGetCwd() string {
	C.cGetCwd()
	goStr := C.GoString(C.buf)
	C.free(unsafe.Pointer(C.buf))
	return goStr
}
