// coding:utf-8
// 验证层
package main

import (
	"CPS/chainmaker/contracts/verify/contract"
	"log"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
)

// main
func main() {
	err := sandbox.Start(new(contract.VERIFY))
	if err != nil {
		log.Fatal(err)
	}
}
