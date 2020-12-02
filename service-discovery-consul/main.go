package main

import (
	"devisions.org/xps/service-discovery-consul/discovery"
	consul "github.com/hashicorp/consul/api"
)

func main() {

	cfg := consul.DefaultConfig()
	// By default, the Consul server listens on `localhost:8500` endpoint.
	// Therefore, it's explicitly defined here just for showcase purposes.
	cfg.Address = "localhost:8500"

	serviceName := "service1"

	client, err := discovery.NewClient(cfg, "localhost", serviceName, 8700)
	if err != nil {
		panic(err)
	}

	if err := discovery.Exec(client, serviceName); err != nil {
		panic(err)
	}
}
