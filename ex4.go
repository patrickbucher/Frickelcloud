package frickelcloud

// CalculateRevenue calculates the server's revenue based on the pricing
// definitions PricingCPU, PricingRAM, PricingSSD.
func (s Server) CalculateRevenue() float32 {
	// TODO: get CPU cost using PricingCPU
	// TODO: get RAM cost using PricingRAM
	// TODO: get SSD cost using PricingSSD
	// TODO: return sum of the three above information
	return 0.0
}

// CalculateRevenue calculates the revenue of all the VMs running on the
// inventories servers.
func (i HostInventory) CalculateTotalRevenue() float32 {
	// TODO: iterate over the inventories hosts
	// TODO: for each host, iterate over the Guest VMs
	// TODO: calculate the revenue per VM and return the sum
	return 0.0
}
