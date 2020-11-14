package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

func GetStatus() string {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "list", "runningvms")
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	return bout.String()
}

func getVMInfo(VMName string, VMStat string) (string, string, bool) {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "showvminfo", VMName)
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		log.Print(berr.String())
	}
	res := bout.String()
	RAM := strings.Replace(strings.Split(res, "Memory size:")[1], " ", "", 20)[:4]
	CPU := strings.Replace(strings.Split(res, "Number of CPUs:")[1], " ", "", 5)[:1]
	on := strings.Contains(VMStat, VMName)
	return RAM, CPU, on
}

type vmsInfo struct {
	Name string `json:"name"`
	RAM  string `json:"ram"`
	CPU  string `json:"cpu"`
	ON   bool   `json:"on"`
}

func getAllVMList() string {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "list", "vms")
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return "error"
	}
	vms := strings.Split(bout.String(), "\n")
	vms = vms[1 : len(vms)-1]
	vmsinfo := []vmsInfo{}
	VMStat := GetStatus()
	for _, vm := range vms {
		name := strings.Replace(strings.Split(vm, " ")[0], "\"", "", 2)
		ram, cpu, on := getVMInfo(name, VMStat)
		vmsinfo = append(vmsinfo, vmsInfo{Name: name, CPU: cpu, RAM: ram, ON: on})
	}
	js, _ := json.Marshal(vmsinfo)
	return string(js)
}

//func main() {
//	log.Print(getAllVMList())
//}
