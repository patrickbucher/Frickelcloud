package frickelcloud

// AvailableRessources computes the Host's available ressources.
func (h Host) AvailableRessources() Server {
	// TODO: sum up allocated ressources of h.Guests
	// TODO: subtract allocted ressources from h.Hardware
	// TODO: return Server struct with the ressources
	return Server{0, 0, 0}
}

// CanFitIn returns true if the host has enough available ressources for the
// given vm, and false otherwise.
func (h Host) CanFitIn(vm Server) bool {
	// TODO: call AvailableRessources on the host
	// TODO: check that all the host's available ressources are larger than the given vm
	return false
}
