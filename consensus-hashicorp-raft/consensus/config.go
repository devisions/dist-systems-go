package consensus

import (
	"fmt"

	"github.com/hashicorp/raft"
)

var rafts map[raft.ServerAddress]*raft.Raft

func init() {
	rafts = make(map[raft.ServerAddress]*raft.Raft)
}

// Config creates a number of `num` `Raft` instances in memory.
func Config(num int) {

	conf := raft.DefaultConfig()
	snapshotStore := raft.NewDiscardSnapshotStore()

	addrs := []raft.ServerAddress{}
	transports := []*raft.InmemTransport{}

	for i := 0; i < num; i++ {
		addr, transport := raft.NewInmemTransport("")
		addrs = append(addrs, addr)
		transports = append(transports, transport)
	}

	memStore := raft.NewInmemStore()

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i != j {
				transports[i].Connect(addrs[j], transports[j])
			}
		}

		// LocalID cannot be empty, otherwise Raft will panic.
		conf.LocalID = raft.ServerID(fmt.Sprintf("%d", i))

		r, err := raft.NewRaft(conf, NewFSM(), memStore, memStore, snapshotStore, transports[i])
		if err != nil {
			panic(err)
		}
		rafts[addrs[i]] = r
	}

}
