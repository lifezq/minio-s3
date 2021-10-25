package storage

import (
	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/config"
)

func NewLocalClient(c *config.Config) client.Client {
	return &LocalStorage{}
}
