## Service Discovery using Consul

This is a simple example of doing service discovery using Consul.

It covers the followings:
1. It uses its own `Client` for service registration (of itself) and discovery.
1. It registers itself as a service.
1. It queries for other services based on a given _key_ and _tag_.

Basically, after registration it is querying Consul for health data of a service with the same name.
The result should at least include the entry representing the service that just has been registered. 

<br/>

### Setup

1. Install Consul locally
    - see [Install Consul](https://learn.hashicorp.com/tutorials/consul/get-started-install) page for details
2. Start the Consul server agent
    - using `consul agent -dev -node=localhost`
    - or use the included `run_consul_agent.sh` script

<br/>
Consul agent running in server mode is listening on port `8500` for HTTP communications with the clients.<br/>
Then you can see the membership of that one-node cluster that is instantiated by this agent using `consul members`:

```shell
$ consul members 
Node       Address         Status  Type    Build  Protocol  DC   Segment
localhost  127.0.0.1:8301  alive   server  1.9.0  2         dc1  <all>
$ 
```

<br/>

## Run

Run `go run main.go` (or use the included `run_main.sh` script).

Example:
```shell
$ ./run_main.sh
>>> Started a new Consul client.
>>> Discovered service(s):
 - {Kind: ID:service1 Service:service1 Tags:[MyTag Consul] Meta:map[] Port:8700 Address:localhost TaggedAddresses:map[] Weights:{Passing:1 Warning:1} EnableTagOverride:false CreateIndex:108 ModifyIndex:167 ContentHash: Proxy:0xc0002006c0 Connect:0xc000238290 Namespace: Datacenter:}
```
