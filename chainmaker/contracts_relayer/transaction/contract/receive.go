// coding:utf-8
// 中继链事务层receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层接收到消息
func (t *TRANSACTION) receive() pb.Response {
	msg, err := t.getArgs()
	if err != nil {
		t.emit_warning([]string{"receive get args error:", err.Error()})
		return sdk.Success(nil)
	}

	// 对swap内的数据进行存储和比对,实现交换
	payload, err := t.saveMarchResource(msg)
	if err != nil {
		// todo:return sdk.error
		// return sdk.Error("save march resource error:" + err.Error())
		t.emit_warning([]string{"save march resource error", err.Error()})
		return sdk.Success(nil)
	}
	return sdk.Success(payload)
}

// 保存和匹配相关资源
// @param msg 传入的消息
// @return []byte 表示MTP事件是否已经匹配成功,处理完成
// @return error err
func (t *TRANSACTION) saveMarchResource(msg *common.ProtocolMsg) ([]byte, error) {
	switch msg.TransactionProtocol {
	case cps_common.TRANSACTION_PROTOCOL_TYPE_BTP:
		return t.saveMarchBtpResource(msg)
	case cps_common.TRANSACTION_PROTOCOL_TYPE_MTP:
		return t.saveMarchMtpResource(msg)
	}
	return nil, fmt.Errorf("invalid transaction protocol type:" + msg.TransactionProtocol)
}
