package main

import (
	"fmt"
	"frickelcloud"
)

var vms = map[string]frickelcloud.Server{
	"vm1": {CPU: 1, RAM: 512, SSD: 10},
	"vm2": {CPU: 2, RAM: 2048, SSD: 20},
	"vm3": {CPU: 4, RAM: 4096, SSD: 40},
	"vm4": {CPU: 8, RAM: 8192, SSD: 80},
}

func main() {
	for vmName, vm := range vms {
		revenue := vm.CalculateRevenue()
		fmt.Printf("The vm %s generates CHF %.2f revenue.\n",
			vmName, revenue)
	}

	totalRevenue := frickelcloud.DataCenter.CalculateTotalRevenue()
	fmt.Printf("FrickelCloud generates CHF %.2f revenue in total.\n",
		totalRevenue)
}
