package vbox

import (
	"bytes"
	"encoding/json"
	"golang.org/x/crypto/ssh"
	"os/exec"
	"strconv"
	"strings"
)

func RunCommand(command string, VMName string) string {
	config := &ssh.ClientConfig{
		User: "",
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(Challenge),
			ssh.Password(""),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "localhost:"+getForwardedPort(VMName), config)
	if err != nil {
		return err.Error()
	}
	session, err := client.NewSession()
	if err != nil {
		return err.Error()
	}
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		return err.Error()
	}
	_ = session.Close()
	return b.String()
}

func CreateVM(vmname string) string {
	var b, berr bytes.Buffer
	cmd := &exec.Cmd{Path: "/home/ali/vm_servers/vm_scripts/create_vm.sh",
		Args:   []string{"", vmname, "Ubuntu_64", "1024", "10000", "3393", "/home/ali/Downloads/ubuntu/u20.iso", "2222"},
		Stdout: &b,
		Stderr: &berr}

	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	return (b.String())
}

func StartVM(vmname string) string {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "startvm", vmname, "-type", "headless")
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	return bout.String()
}

func StopVM(vmname string) string {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "controlvm", vmname, "poweroff", "soft")
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	return bout.String()
}

func RemoveVM(vmname string) string {
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "unregistervm", vmname, "--delete")
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	return bout.String()
}

func ChangeVMProperties(vmname string, RAM int, CPU int) string {
	StopVM(vmname)
	var bout, berr bytes.Buffer
	cmd := exec.Command("vboxmanage", "modifyvm", vmname, "--cpus", strconv.Itoa(CPU), "--memory", strconv.Itoa(RAM))
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return berr.String()
	}
	StartVM(vmname)
	return bout.String()
}

func GetAllVMList() string {
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

func CloneVM(VMName string) string {
	var b, berr, bout bytes.Buffer
	cmd := exec.Command("vboxmanage", "list", "vms")
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	if err := cmd.Run(); err != nil {
		return "error"
	}
	vms := strings.Split(bout.String(), "\n")
	vms = vms[1 : len(vms)-1]
	cmd2 := &exec.Cmd{Path: "/home/ali/go/src/github.com/aicam/virtualbox-web-service/clone_vm.sh",
		Args:   []string{"", VMName + "_2", strconv.Itoa(len(vms) + 2222), VMName},
		Stdout: &b,
		Stderr: &berr}
	go cmd2.Run()
	return "Started"
}
