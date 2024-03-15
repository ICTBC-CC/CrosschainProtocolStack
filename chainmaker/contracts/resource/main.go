// coding:utf-8
// 资源和余额
package main

import (
	"CPS/chainmaker/contracts/resource/contract"
	"log"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
)

// main
func main() {
	err := sandbox.Start(new(contract.Resource))
	if err != nil {
		log.Fatal(err)
	}
}
