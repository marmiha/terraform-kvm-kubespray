package modelconfig

type LoadBalancer struct {
	Default *struct {
		CPU          *CpuSize
		MainDiskSize *MB
		RAM          *MB
	}
	ForwardPorts    []ForwardPort
	Instances       []LoadBalancerInstance
	VIP             *IP
	VirtualRouterId *LoadBalancerId
}
