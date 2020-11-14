### VirtualBox web service with GoLang

This project is a web service to control VirtualBox with the following features. This project uses vbox package to communicate with VBoxManage CLI commands. All tasks are done by CLI api.

### Features
Full clone, Delete, Change resources(RAM, CPU), SSH into VM, Turn on/off

### API
`GET /get_vms_info` Returns all VMs status with resource information and off/on status <br>
`GET /stop_vm/<VM name>` Turns off the VM <br>
`GET /remove_vm/<VM name>` Remove VM <br>
`GET /config_vm/<VM name>/<RAM>/<CPU cores>` Changes VM resources by RAM in MB and CPU by number of cores <br>
`GET /start_vm/<VM name>` Starts VM <br>
`POST /run_command` Run commands in a selected VM <br>
`{
	"vm_name": <VM name>
	"command": <Command>
}`

### How commands are ran inside a VM

Each VM is created with NAT network and has a forwarded port from host. In this regard, each time our server ssh to the existing VM and runs the command and returns the result. If you clone a VM, it automatically assignes a new port in host for ssh, the port is 2221 + number of VMs.
