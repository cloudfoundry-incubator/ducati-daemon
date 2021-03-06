package models

type Container struct {
	ID          string `json:"id"`
	IP          string `json:"ip"`
	MAC         string `json:"mac"`
	HostIP      string `json:"host_ip" db:"host_ip"`
	NetworkID   string `json:"network_id" db:"network_id"`
	SandboxName string `json:"sandbox_name" db:"sandbox_name"`
	App         string `json:"app" db:"app"`
}
