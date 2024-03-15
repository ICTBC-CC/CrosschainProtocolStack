// coding:utf-8
// 接收函数
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

	t.emit_info([]string{"relayer 收到 msg"})

	// 1. 执行转发层
	if err := t.receiveTransfer(msg.FromChain, msg.ToChain); err != nil {
		return sdk.Error(fmt.Sprintf("call receive transfer error:%s", err))
	}
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"call receive transfer success"})

	// 2. 执行验证层
	if err := t.receiveVerify(msg.VerifyType, msg.VerifyData); err != nil {
		return sdk.Error(fmt.Sprintf("call receive verify error:%s", err))
	}
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"call receive verify success"})

	// 3. 执行事务层
	// ! 好好考虑事务层失败后是否回滚, 怎么回滚, 回滚哪些, 还是说等待链下计时器触发
	if err := t.receiveTransaction(msg); err != nil {
		return sdk.Error(fmt.Sprintf("call receive transaction error:%s", err))
	}
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"call receive transaction success"})

	return sdk.Success(nil)
}
