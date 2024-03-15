// coding:utf-8
// 协议聚合器发送
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 将消息发送出去
// @return pb.response response
func (t *ProtocolAggregator) send() pb.Response {
	protoMsg, err := t.getArgs()
	if err != nil {
		return sdk.Error("get args error:" + err.Error())
	}
	t.emit_info([]string{"初始app data", string(protoMsg.AppData)})

	// 1. 执行应用层,返回应用层跨链数据
	appData, err := t.sendApp(protoMsg)
	if err != nil {
		return sdk.Error("call send app error:" + err.Error())
	}
	// 重设跨链层数据
	protoMsg.AppData = appData
	t.emit_info([]string{"call send app success", "app data", string(appData)})

	// 2. 执行事务层
	// send, 则初始化事务状态为发起状态
	protoMsg.TransactionState = common.TRANSACTION_PRE_PREPARE
	transactionID, err := t.sendTransaction(protoMsg)
	if err != nil {
		return sdk.Error("call send transaction error:" + err.Error())
	}
	// 修改事务状态
	protoMsg.TransactionID = transactionID
	protoMsg.TransactionState = common.TRANSACTION_PREPARE
	t.emit_info([]string{"call send transaction success", "transaction id:", string(transactionID)})

	// 3. 执行验证层或回滚事务层
	if err := t.sendVerify(); err != nil {
		return sdk.Error(fmt.Sprintf("call send verify error:%s", err))
	}
	t.emit_info([]string{"call send verify success"})

	// 4. 执行转发层或回滚事务验证层
	if err := t.sendTransfer(protoMsg); err != nil {
		return sdk.Error(fmt.Sprintf("call send transfer error:%s", err))
	}
	t.emit_info([]string{"call send transfer success"})

	return sdk.Success(nil)
}
