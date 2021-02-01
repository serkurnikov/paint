package gocv

/*
#include <stdlib.h>
#include "palette.h"
*/
import "C"
import (
	"unsafe"
)

func BuildPalette(imageFilename, imageOutFilename string) {
	cIn := C.CString(imageFilename)
	cOut := C.CString(imageOutFilename)
	defer C.free(unsafe.Pointer(cIn))
	defer C.free(unsafe.Pointer(cOut))
	
	C.BuildPalette(cIn, cOut)
}