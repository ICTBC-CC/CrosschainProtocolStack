// coding:utf-8
// 协议栈接收到一个跨链消息
package contract

import (
	"CPS/chainmaker/contracts/common"
	"encoding/json"
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 从网关收到一个消息
// @return pb.response response
func (t *ProtocolAggregator) receive() pb.Response {
	// todo:确保输入参数的有效性
	args := sdk.Instance.GetArgs()
	// 协议数据
	msg := &common.ProtocolMsg{}
	// 恢复数据
	if err := json.Unmarshal(args["data"], msg); err != nil {
		return sdk.Error("receive unmarshal data error:" + err.Error())
	}
	// 查看发送上来的数据
	t.emit_info([]string{"protocol aggregator receive msg"})

	// 1. 执行转发层
	if err := t.receiveTransfer(msg.FromChain, msg.ToChain, msg.TransferState); err != nil {
		return sdk.Error(fmt.Sprintf("call receive transfer error:%s", err))
	}
	t.emit_info([]string{"call receive transfer success"})

	// 2. 执行验证层
	if err := t.receiveVerify(msg.VerifyType, msg.VerifyData); err != nil {
		return sdk.Error(fmt.Sprintf("call receive verify error:%s", err))
	}
	t.emit_info([]string{
		"call receive verify success",
		string(msg.VerifyType),
	})

	// 3. 执行事务层
	// ! 好好考虑事务层失败后是否回滚, 怎么回滚, 回滚哪些, 还是说等待链下计时器触发
	if err := t.receiveTransaction(msg); err != nil {
		t.emit_warning([]string{
			"protocol aggregator receive transaction error",
			err.Error(),
			string(msg.FromApp), string(msg.AppFunc),
		})
		return sdk.Error("call receive transaction error:" + err.Error())
	}
	// 修改事务状态为提交
	msg.TransactionState = common.TRANSACTION_COMMIT
	t.emit_info([]string{"call recieve transaction success"})

	// 4. 执行应用层
	t.emit_info([]string{"receive app data", string(msg.AppData)})
	appData, err := t.receiveApp(msg)
	if err != nil {
		t.emit_warning([]string{
			"protocol aggregator receive app error",
			err.Error(),
		})
		return sdk.Error("call receive app error:" + err.Error())
	}
	t.emit_info([]string{"receive app response data", string(appData)})
	msg.AppData = appData

	// 5. 调用转发层给peer响应
	if err := t.responseTransfer(msg); err != nil {
		return sdk.Error(fmt.Sprintf("call response transfer error:" + err.Error()))
	}

	return sdk.Success(nil)
}
