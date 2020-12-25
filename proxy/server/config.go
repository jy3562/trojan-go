package server

import (
	"github.com/jy3562/trojan-go/config"
	"github.com/jy3562/trojan-go/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return new(client.Config)
	})
}
