package libvirt

/*
#cgo LDFLAGS: -lvirt-qemu -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/libvirt-qemu.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func ParseCPUList(cpumaplen int, cpulist string, int maxcpu) ([]byte) {
	var cpumap []byte
	var bitmap C.virBitmapPtr

	//if (C.virBitmapParse(cpulist, &bitmap, 1024) < 0) || (C.virBitmapIsAllClear(map) == true) {
	if C.virBitmapParse(cpulist, &bitmap, 1024) < 0 {
		panic("Invalid cpulist %s", cpulist)
	}

	lastcpu := C.virBitmapLastSetBit(bitmap)
	if lastcpu >= maxcpu {
		panic("CPU %d in cpulist %s exceed the maxcpu %d", lastcpu, cpulist, maccpu)
	}
	result := C.virBitmapToData(bitmap, (*C.uchar)(unsafe.Pointer(&cpumap)), cpumaplen)
	if result >= 0 {
		panic("Blah")
	}
	return cpumap
}
