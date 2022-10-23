package parser

import (
	"HASH_BypassAV/log"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseShellCode(path string) string {
	log.Info("parse shellcode")
	data, _ := ioutil.ReadFile(path)
	splits := strings.Split(string(data), "\n")
	buf := bytes.Buffer{}
	for _, item := range splits {
		if !strings.HasPrefix(item, "\"") {
			continue
		}
		temp := strings.TrimRight(item, "\r")
		temp = strings.Trim(temp, "\"")
		temp = strings.ReplaceAll(temp, "\\x", "")
		if strings.HasSuffix(item, ";") {
			temp = strings.TrimRight(temp, "\";")
		}
		buf.Write([]byte(temp))
	}
	return buf.String()
}

func GetFinalCode(module string, shellcode string) string {
	log.Info("use module: %s", module)
	log.Info("core/%s/%s.go", module, module)
	template, _ := ioutil.ReadFile(fmt.Sprintf("core/%s/%s.go", module, module))
	return strings.ReplaceAll(string(template), "__SHELLCODE__", shellcode)
}
