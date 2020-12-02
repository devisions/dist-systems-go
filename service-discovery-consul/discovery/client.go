package discovery

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type Client interface {
	Register(tags []string) error
	GetServiceHealth(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error)
}

// client includes a Consul client and implements the `Client` spec.
type client struct {
	client  *api.Client
	address string
	name    string
	port    int
}

// NewClient creates a new client, used for registering and discovering services.
func NewClient(cfg *api.Config, address, name string, port int) (Client, error) {

	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println(">>> Started a new Consul client.")
	return &client{
		client:  c,
		name:    name,
		address: address,
		port:    port,
	}, nil
}

// Exec performs
func Exec(c Client, queryServiceName string) error {

	if err := c.Register([]string{"MyTag", "Consul"}); err != nil {
		return err
	}

	entries, _, err := c.GetServiceHealth(queryServiceName, "Consul")
	if err != nil {
		return err
	}

	if len(entries) > 0 {
		fmt.Println(">>> Discovered service(s):")
		for _, e := range entries {
			fmt.Printf(" - %+v\n", *e.Service)
		}
	} else {
		fmt.Println(">>> No service has been discovered.")
	}
	return nil
}

// Register is registering the service that this client represents.
// Provided `tags` are associated with the service registration.
func (c *client) Register(tags []string) error {

	reg := &api.AgentServiceRegistration{
		ID:      c.name,
		Name:    c.name,
		Port:    c.port,
		Address: c.address,
		Tags:    tags,
	}
	return c.client.Agent().ServiceRegister(reg)
}

// GetServiceHealth is querying the health information of a service based on provided `svcName`.
// The provided `tag` can be used for additional server-side filtering.
func (c *client) GetServiceHealth(serviceName, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {

	return c.client.Health().Service(serviceName, tag, false, nil)
}
