package main

import (
	"net/http"

	"devisions.org/xps/consensus-hashicorp-raft/consensus"
)

func main() {

	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)

	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
