// coding:utf-8
// 事务层发送
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 事务层发送
// @return response 交易id
func (t *TRANSACTION) send() pb.Response {
	msg, err := t.getArgs()
	if err != nil {
		return sdk.Error("get send args error:" + err.Error())
	}

	// 判断状态是否正确
	if msg.TransactionState != common.TRANSACTION_PRE_PREPARE {
		return sdk.Error("transaction state should be pre_prepare, get:" + msg.TransactionState)
	}

	// 分配id号
	currID, err := t.getUpdateNewestID()
	if err != nil {
		return sdk.Error("get and update newest transaction id error" + err.Error())
	}

	// 执行对应的协议
	switch msg.TransactionProtocol {
	case cps_common.TRANSACTION_PROTOCOL_TYPE_BTP:
		{
			// 执行btp
			if err := t.sendBtp(currID, msg); err != nil {
				return sdk.Error("send btp error:" + err.Error())
			}
		}
	case cps_common.TRANSACTION_PROTOCOL_TYPE_MTP:
		{
			// 执行mtp
			if err := t.sendMtp(currID, msg); err != nil {
				return sdk.Error("send mtp error:" + err.Error())
			}
		}
	default:
		return sdk.Error("invalid protocol type:" + msg.TransactionProtocol)
	}

	// 获取id的[]byte
	idBytes, err := json.Marshal(currID)
	if err != nil {
		return sdk.Error("marshal currid error:" + err.Error())
	}

	return sdk.Success(idBytes)
}
