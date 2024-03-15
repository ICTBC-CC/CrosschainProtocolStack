// coding:utf-8
// 转发层
package main

import (
	"CPS/chainmaker/contracts/transfer/contract"
	"log"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
)

// main
func main() {
	err := sandbox.Start(new(contract.TRANSFER))
	if err != nil {
		log.Fatal(err)
	}
}
