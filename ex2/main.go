package main

import (
	"fmt"
	"frickelcloud"
)

var newVMs = map[string]frickelcloud.Server{
	"vm1": {CPU: 1, RAM: 512, SSD: 10},
	"vm2": {CPU: 2, RAM: 4096, SSD: 80},
	"vm3": {CPU: 8, RAM: 8192, SSD: 1000},
}

func main() {
	for vmName, vm := range newVMs {
		fittingHosts := frickelcloud.DataCenter.FindFittingHosts(vm)
		if len(fittingHosts) == 0 {
			fmt.Printf("Sorry, cannot fit VM %s (%v) into any host.\n",
				vmName, vm)
		}

		for hostName, host := range fittingHosts {
			err := host.FitIn(vm, vmName)
			if err != nil {
				fmt.Printf("Sorry, cannot fit VM %s (%v) into any host.\n",
					vmName, vm)
			} else {
				fmt.Printf("The VM %s (%v) was fitted in to host %s.",
					vmName, vm, hostName)
			}
			break // we only need to fit in the VM once!
		}
	}
}
