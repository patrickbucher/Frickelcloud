package main

import (
	"fmt"
	"frickelcloud"
)

var vmsToBeRemoved = []string{
	"file-server",
	"dns-server",
	"windows-server",
}

func main() {
	for _, vmName := range vmsToBeRemoved {
		err := frickelcloud.DataCenter.Remove(vmName)
		if err != nil {
			fmt.Printf("the VM %s wasn't removed from the inventory: %v\n",
				vmName, err)
		} else {
			fmt.Printf("successfully removed VM %s from the inventory\n",
				vmName)
		}
	}
}
