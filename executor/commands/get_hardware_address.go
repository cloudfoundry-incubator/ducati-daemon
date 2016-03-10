package commands

import (
	"fmt"
	"net"
)

type GetHardwareAddress struct {
	LinkName string
	Result   net.HardwareAddr
}

func (cmd *GetHardwareAddress) Execute(context Context) error {
	hwAddr, err := context.LinkFactory().HardwareAddress(cmd.LinkName)
	if err != nil {
		return fmt.Errorf("get hardware address: %s", err)
	}

	cmd.Result = hwAddr

	return nil
}

func (cmd *GetHardwareAddress) String() string {
	return fmt.Sprintf("ip link show %s", cmd.LinkName)
}
