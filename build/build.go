package build

import (
	"HASH_BypassAV/log"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func Build(code string, module string, debug bool, strip bool, hide bool) {
	log.Info("build...")
	cmd := []string{
		"build",
		"-o",
		"output.exe",
		"-ldflags",
		"",
		"output/main.go",
	}
	if debug {
		cmd[4] = "-s -w"
	}
	if hide {
		cmd[4] = "-s -w -H windowsgui"
	}
	privateBuild(code, cmd, module)
}

func privateBuild(code string, command []string, module string) {
	_ = os.RemoveAll(filepath.Join(".", "output.exe"))
	newPath := filepath.Join(".", "output")
	_ = os.MkdirAll(newPath, os.ModePerm)
	_ = ioutil.WriteFile("output/main.go", []byte(code), 0777)
	cpcommand := "xcopy.exe .\\core\\" + module + " .\\output"
	cmd := exec.Command("cmd", cpcommand)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	cmd = exec.Command("go", command...)
	err = cmd.Run()
	if err == nil {
		log.Info("build success")
		log.Info("file: output.exe")
	} else {
		log.Error("error")
	}
	_ = os.RemoveAll(newPath)
}
