package main

import (
	"github.com/adarobin/docker-machine-driver-qemu/qemu"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(qemu.NewDriver("default", "path"))
}
