package frickelcloud

// PricingCPU indicates the price (value) by number of cores (key).
var PricingCPU = map[uint8]float32{
	1:  1.0, // 1 CPU Core = CHF 1.00
	2:  2.0,
	4:  4.0,
	8:  8.0,
	16: 16.0,
}

// PricingRAM indicates the price (value) by megabytes of memory (key).
var PricingRAM = map[uint32]float32{
	512:   1.0, // 512 MB RAM = CHF 1.00
	1024:  2.0,
	2048:  4.0,
	4096:  8.0,
	8192:  16.0,
	16384: 32.0,
	32768: 64.0,
}

// PricingSSD indicates the price (value) by gigabytes of storage (key).
var PricingSSD = map[uint16]float32{
	10:   1.0, // 10 GB SSD = CHF 1.00
	20:   2.0,
	40:   3.0,
	80:   4.0,
	240:  5.0,
	500:  6.0,
	1000: 8.0,
}

// Server is specification for a machine (physical or virtual).
type Server struct {
	CPU uint8
	RAM uint32
	SSD uint16
}

// GuestInventory maps names to virtual machines.
type GuestInventory map[string]Server

// Host is a physical server running multiple named guests.
type Host struct {
	Hardware Server
	Guests   GuestInventory
}

// HostInventory maps names to physical host machines.
type HostInventory map[string]Host

// DataCenter is the initial hardware configuration and VM allocation.
var DataCenter = HostInventory{
	"small-1": Host{
		Hardware: Server{
			CPU: 4,
			RAM: 32768,
			SSD: 4000,
		},
		Guests: GuestInventory{
			"web-server": Server{
				CPU: 2,
				RAM: 1024,
				SSD: 240,
			},
			"file-server": Server{
				CPU: 1,
				RAM: 1024,
				SSD: 80,
			},
		},
	},
	"medium-1": Host{
		Hardware: Server{
			CPU: 8,
			RAM: 65536,
			SSD: 8000,
		},
		Guests: GuestInventory{
			"application-server": Server{
				CPU: 4,
				RAM: 8192,
				SSD: 1000,
			},
			"db-server": Server{
				CPU: 2,
				RAM: 4096,
				SSD: 500,
			},
			"dns-server": Server{
				CPU: 1,
				RAM: 1024,
				SSD: 500,
			},
		},
	},
	"big-1": Host{
		Hardware: Server{
			CPU: 16,
			RAM: 131072,
			SSD: 16000,
		},
		Guests: GuestInventory{
			"backend-server-1": Server{
				CPU: 4,
				RAM: 8192,
				SSD: 20,
			},
			"backend-server-2": Server{
				CPU: 4,
				RAM: 8192,
				SSD: 20,
			},
			"load-balancer": Server{
				CPU: 2,
				RAM: 2048,
				SSD: 10,
			},
			"git-server": Server{
				CPU: 1,
				RAM: 512,
				SSD: 500,
			},
			"web-server": Server{
				CPU: 4,
				RAM: 4096,
				SSD: 40,
			},
		},
	},
}
