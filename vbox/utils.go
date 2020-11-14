package vbox

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var suitable_answers = []string{"bogus password", ""}
var pwIdx = 0

type vmsInfo struct {
	Name string `json:"name"`
	RAM  string `json:"ram"`
	CPU  string `json:"cpu"`
	ON   bool   `json:"on"`
}

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

func Challenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	for n, q := range questions {
		fmt.Printf("Got question: %s\n", q)
		answers[n] = suitable_answers[pwIdx]
	}
	pwIdx++

	return answers, nil
}

func getForwardedPort(VMName string) string {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "showvminfo", VMName)
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	nic_1 := strings.Split(bout.String(), "NIC 1 Rule(0)")
	port := strings.Split(nic_1[1], "host port = ")[1][:4]
	return port
}

func getVMInfo(VMName string, VMStat string) (string, string, bool) {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "showvminfo", VMName)
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return "err", "err", false
	}
	res := bout.String()
	RAM := strings.Replace(strings.Split(res, "Memory size:")[1], " ", "", 20)[:4]
	CPU := strings.Replace(strings.Split(res, "Number of CPUs:")[1], " ", "", 5)[:1]
	on := strings.Contains(VMStat, "\""+VMName+"\"")
	return RAM, CPU, on
}
