package main

import (
	"fmt"
	"p2p-node-with-static-node-id/p2pnode"
)

func main() {
	_, host, err := p2pnode.EstablishP2PNode()
	if err != nil {
		fmt.Errorf("%s", err.Error())

	} else {
		fmt.Println("Completed node establishment proceedure")
		fmt.Println(host.ID())
	}
	for i := 0; ; i++ {

	}
}
