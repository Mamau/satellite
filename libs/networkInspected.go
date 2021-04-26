package libs

import "time"

// NetworkInspected
// You can get it by  docker network inspect bridge
// Use it for determine your Gateway host
type NetworkInspected struct {
	Name    string    `json:"Name"`
	Created time.Time `json:"Created"`
	Scope   string    `json:"Scope"`
	Driver  string    `json:"Driver"`
	IPAM    struct {
		Config []struct {
			Subnet  string `json:"Subnet"`
			Gateway string `json:"Gateway"`
		} `json:"Config"`
	} `json:"IPAM"`
}

func (n *NetworkInspected) GetGateway() string {
	config := n.IPAM.Config[0]
	if &config != nil {
		return config.Gateway
	}
	return ""
}
