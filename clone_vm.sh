#!/bin/bash

VBoxManage clonevm $3 --name=$1 --mode=all --register
VBoxManage modifyvm $1 --natpf1 delete "guestssh"
VBoxManage modifyvm $1 --natpf1 "guestssh,tcp,,$2,,22"
