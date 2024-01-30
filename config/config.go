package config

import (
	"os"
)

type Conf struct {
	RMQ          *RMQ
	MountPath    string
	MaterialPath string
}

func NewConf() *Conf {
	return &Conf{
		RMQ:          newRMQ(),
		MountPath:    os.Getenv("MOUNT_PATH"),
		MaterialPath: os.Getenv("MATERIAL_PATH"),
	}
}
