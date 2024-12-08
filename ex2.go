package frickelcloud

// FindFittingHosts returns a map containing the inventories hosts that are
// capable of fitting in the given vm.
func (i HostInventory) FindFittingHosts(vm Server) map[string]Host {
	// TODO: iterate over the HostInventory
	// TODO: for each host, check if it can fit in the given vm
	// TODO: add all hosts with enough free ressources to a map, and return it
	return map[string]Host{}
}

// FitIn adds the given vm to the host, and returns an error if the operation
// would exceed the host's available ressources.
func (h Host) FitIn(vm Server, name string) error {
	// TODO: make sure that the host has enough available ressources to fit in vm
	// TODO: return an error if the host has too few available ressources

	// TODO: make sure that the host not yet has a vm with the given name
	// TODO: return an error if the name was already taken

	// TODO: add the vm to the host's inventory
	return nil
}
