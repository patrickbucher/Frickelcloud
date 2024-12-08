package frickelcloud

// FindHost returns the name of the host on which the VM with the given vmName
// is running, or an error if the VM cannot be found in any of the hosts guest
// inventories.
func (i HostInventory) FindHost(vmName string) (string, error) {
	// TODO: iterate over the HostIntenvory
	// TODO: for each host, check if it contains the VM with vmName
	// TODO: return the host's name, if the VM was found on it
	// TODO: return an empty string and an according error, if the VM wasn't found anywhere
	return "", nil
}

// Remove removes the VM with the given name from the inventory. An error is
// returned, if the operation failed.
func (i HostInventory) Remove(vmName string) error {
	// TODO: use the method FindHost to check if the VM with the given vmName actually exists in the inventory
	// TODO: return an error if the VM cannot be found
	// TODO: use the found host name to access the host on the inventory
	// TODO: delete the VM from the host's guests inventory
	return nil
}
