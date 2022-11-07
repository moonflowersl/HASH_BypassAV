package parser

import (
	"HASH_BypassAV/log"
	"fmt"
	"io/ioutil"
	"strings"
)

func OriginShellCode(path string) []byte {
	log.Info("parse shellcode")
	data, _ := ioutil.ReadFile(path)
	return data
}

func ParseShellCode(path string) string {
	log.Info("parse shellcode")
	data, _ := ioutil.ReadFile(path)
	data_string := string(data)
	pos := strings.Index(data_string, "\"")
	data_string = data_string[pos:]
	data_string = strings.ReplaceAll(data_string, "\n", "")
	data_string = strings.ReplaceAll(data_string, "\"", "")
	data_string = strings.ReplaceAll(data_string, ";", "")
	data_string = strings.ReplaceAll(data_string, "\\x", "")
	return data_string
}

func GetFinalCode(module string, shellcode string) string {
	log.Info("use module: %s", module)
	log.Info("core/%s/%s.go", module, module)
	template, _ := ioutil.ReadFile(fmt.Sprintf("core/%s/%s.go", module, module))
	return strings.ReplaceAll(string(template), "__SHELLCODE__", shellcode)
}
