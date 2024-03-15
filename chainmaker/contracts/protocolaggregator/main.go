// coding:utf-8
// 协议聚合器
package main

import (
	"CPS/chainmaker/contracts/protocolaggregator/contract"
	"log"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
)

// main
func main() {
	err := sandbox.Start(new(contract.ProtocolAggregator))
	if err != nil {
		log.Fatal(err)
	}
}
