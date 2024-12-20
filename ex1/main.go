package main

import (
	"fmt"
	"github.com/patrickbucher/frickelcloud"
)

var newVMs = map[string]frickelcloud.Server{
	"vm1": {CPU: 1, RAM: 512, SSD: 10},
	"vm2": {CPU: 2, RAM: 4096, SSD: 80},
	"vm3": {CPU: 8, RAM: 8192, SSD: 1000},
}

func main() {
	for hostName, host := range frickelcloud.DataCenter {
		free := host.AvailableRessources()
		fmt.Printf("host %s has %d CPU cores, %d MB memory, and %d GB storage available\n",
			hostName, free.CPU, free.RAM, free.SSD)

		for vmName, vm := range newVMs {
			fits := host.CanFitIn(vm)
			fmt.Printf("can %s fit in the VM %s (%v)? %v\n",
				hostName, vmName, vm, fits)
		}
	}
}
