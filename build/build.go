package build

import (
	"HASH_BypassAV/log"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func Build(code string) {
	log.Info("build...")

	cmd := []string{
		"build",
		"-o",
		"output.exe",
		"output/main.go",
	}
	privateBuild(code, cmd)
}

func privateBuild(code string, command []string) {
	_ = os.RemoveAll(filepath.Join(".", "output.exe"))
	newPath := filepath.Join(".", "output")
	_ = os.MkdirAll(newPath, os.ModePerm)
	_ = ioutil.WriteFile("output/main.go", []byte(code), 0777)
	cmd := exec.Command("go", command...)
	err := cmd.Run()
	if err == nil {
		log.Info("build success")
		log.Info("file: output.exe")
	} else {
		log.Error("error")
	}
	//_ = os.RemoveAll(newPath)
}
