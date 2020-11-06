package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os/exec"
)

var suitable_answers = []string{"bogus password", "021021"}
var pwIdx = 0

func RunCommand(){
	config := &ssh.ClientConfig{
		User: "aicam",
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(Challenge),
			ssh.Password("021021"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "localhost:2222", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("echo \"verbos\""); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	_ = session.Close()
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

func getPass() (string, error) {
	return "021021", nil
}

func CreateVM(){
	var b, berr bytes.Buffer
	cmd := &exec.Cmd{Path:"/home/ali/vm_servers/vm_scripts/create_vm.sh",
		Args:[]string{"", "vm1", "Ubuntu_64", "1024", "10000", "3393", "/home/ali/Downloads/ubuntu/u20.iso", "2222"},
	Stdout: &b,
	Stderr: &berr}

	if err := cmd.Run(); err != nil {
		log.Fatal("Error :", berr.String())
	}
	log.Print(b.String())
}

func GetStatus() string {
	var bout bytes.Buffer
	cmd := exec.Command("vboxmanage", "list", "runningvms")
	cmd.Stdout = &bout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return bout.String()
}

func StartVM(vmname string) string {
	var bout bytes.Buffer
	cmd := exec.Command("vboxmanage", "startvm", vmname, "-type", "headless")
	cmd.Stdout = &bout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return bout.String()
}

func StopVM(vmname string) string {
	var bout bytes.Buffer
	cmd := exec.Command("vboxmanage", "startvm", vmname, "-type", "headless")
	cmd.Stdout = &bout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return bout.String()
}


func main() {

}