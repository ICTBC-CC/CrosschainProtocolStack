// coding:utf-8
// 验证层
package main

import (
	"CPS/chainmaker/contracts/common"
	"log"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type VERIFY struct{}

func (t *VERIFY) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *VERIFY) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *VERIFY) InvokeContract(method string) pb.Response {
	// according method segment to select contract functions
	switch method {
	case common.FUNC_SEND:
		// 发送
		return t.send()
	case common.FUNC_RECEIVE:
		// 收到消息
		return t.receive()
	default:
		return sdk.Error("invalid function")
	}
}

// 转发层发送
func (t *VERIFY) send() pb.Response {
	// 抛出数据给链下网关
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"verify send"})
	return sdk.Success(nil)
}

// 转发层接收到消息
func (t *VERIFY) receive() pb.Response {
	// 抛出数据给
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"verify receive"})
	return sdk.Success(nil)
}

// main
func main() {
	err := sandbox.Start(new(VERIFY))
	if err != nil {
		log.Fatal(err)
	}
}
