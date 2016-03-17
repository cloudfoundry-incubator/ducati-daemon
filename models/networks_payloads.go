package models

import "github.com/appc/cni/pkg/types"

type NetworksSetupContainerPayload struct {
	Args               string       `json:"args"`
	ContainerNamespace string       `json:"container_namespace"`
	InterfaceName      string       `json:"interface_name"`
	IPAM               types.Result `json:"ipam"`
	NetworkID          string       `json:"network_id"`
	ContainerID        string       `json:"container_id"`
}

type NetworksDeleteContainerPayload struct {
	InterfaceName      string `json:"interface_name"`
	ContainerNamespace string `json:"container_namespace"`
	NetworkID          string `json:"network_id"`
	ContainerID        string `json:"container_id"`
}
