// coding:utf-8
// 转发层
package main

import (
	"CPS/chainmaker/contracts/common"
	"bytes"
	"encoding/json"
	"log"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type TRANSFER struct{}

func (t *TRANSFER) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *TRANSFER) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *TRANSFER) InvokeContract(method string) pb.Response {
	args := sdk.Instance.GetArgs()
	fromchain := args[common.KEY_FROM_CHAIN]
	tochain := args[common.KEY_TO_CHAIN]
	// 是整个协议的数据包
	data := args[common.KEY_TRANSFER_DATA]

	// according method segment to select contract functions
	switch method {
	case common.FUNC_SEND:
		// 发送
		return t.send(fromchain, tochain, data)
	case common.FUNC_RECEIVE:
		// 收到消息
		return t.receive(fromchain, tochain)
	case common.FUNC_RESPONSE:
		// 响应peer的消息
		return t.response(data)
	default:
		return sdk.Error("invalid function")
	}
}

// 转发层响应给peer
// @return response
func (t *TRANSFER) response(data []byte) pb.Response {
	// 本地处理完成,回复给对端链
	// 1. 交换源链和目的链
	// 2. 改变消息状态为响应
	msg := &common.ProtocolMsg{}
	// 解码数据包
	if err := json.Unmarshal(data, msg); err != nil {
		return sdk.Error(err.Error())
	}
	// 判断ack状态
	if bytes.Equal(msg.TransferState, common.TRANSFER_STATE_SEND) {
		// 判断ACK状态, 如果状态为send就需要改为receive,并交换源与目的链
		msg.TransferState = common.TRANSFER_STATE_RESPONSE
		tmp := msg.FromChain
		msg.FromChain = msg.ToChain
		msg.ToChain = tmp
		// 编码
		data, err := json.Marshal(msg)
		if err != nil {
			return sdk.Error(err.Error())
		}
		// 通知网关
		if err := t.sendToRelayer(msg.FromChain, msg.ToChain, data); err != nil {
			return sdk.Error(err.Error())
		}
		sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"response to peer"})
		return sdk.Success(nil)
	} else if bytes.Equal(msg.TransferState, common.TRANSFER_STATE_RESPONSE) {
		// 判断ack状态, 如果状态为receive就表示是自己发出的, 表示事务完成,不做响应
		sdk.Instance.EmitEvent(
			common.EVENT_INFO,
			[]string{
				"not deal",
				string(msg.FromChain),
				string(msg.ToChain),
			},
		)
		return sdk.Success(nil)
	}
	// 其他ack状态表示错误
	return sdk.Error("invalid msg ack " + string(msg.TransferState))
}

// 转发层发送
func (t *TRANSFER) send(fromchain, tochain, data []byte) pb.Response {
	// 抛出数据给链下网关
	// 源链, 目的链, 整个消息的字节码
	sdk.Instance.EmitEvent(
		common.EVENT_TRANSFER_SEND,
		[]string{
			string(fromchain),
			string(tochain),
			string(data),
		},
	)
	return sdk.Success(nil)
}

// 转发层接收到消息
func (t *TRANSFER) receive(fromchain, tochain []byte) pb.Response {
	// 抛出数据给
	sdk.Instance.EmitEvent(common.EVENT_TRANSFER_RECEIVE, []string{string(fromchain), string(tochain)})
	return sdk.Success(nil)
}

// 发送链上情况给链下网关
// @param fromchain 消息的源链
// @param tochain 消息的目的链
// @param data 抛出的数据
func (t *TRANSFER) sendToRelayer(fromchain, tochain, data []byte) error {
	// ! 目前就用事件的形式实现转发层的通知
	// 源链, 目的链, 整个消息的字节码
	sdk.Instance.EmitEvent(
		common.EVENT_TRANSFER_SEND,
		[]string{
			string(fromchain),
			string(tochain),
			string(data),
		},
	)
	return nil
}

// main
func main() {
	err := sandbox.Start(new(TRANSFER))
	if err != nil {
		log.Fatal(err)
	}
}
