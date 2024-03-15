// coding:utf-8
// 事务层
package main

import (
	"CPS/chainmaker/contracts/transaction/contract"
	"log"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
)

// main
func main() {
	err := sandbox.Start(new(contract.TRANSACTION))
	if err != nil {
		log.Fatal(err)
	}
}
