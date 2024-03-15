// coding:utf-8
// 客户端内部消息转为订阅格式的消息
package client

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"
	"fmt"
)

// 订阅的MSG消息格式转为长安链的kvs,配合invoke函数
// @param msg 被转化的订阅MSG消息
// @return any 客户端内部消息
// @return error 错误消息
func (c *Client) SubscribeMsg2ReceiveArgs(msg *cps_common.SubscribeMsg) (any, error) {
	// 长安链部分链上直接使用unmarshal就可以了,所以不用麻烦的进行编码,也不需要做成kv键值对
	// 直接转为链上数据格式,再marshal为字节码以data传上去

	pack := common.ProtocolMsg{}
	// ***** 资源 *****
	pack.AtomicSwap = msg.AtomicSwap
	// ***** 应用层 *****
	pack.FromApp = msg.AppFrom
	pack.AppFunc = msg.AppFunc
	pack.AppData = msg.AppData
	// ***** 事务层 *****
	pack.TransactionData = msg.TransactionData
	pack.TransactionState = msg.TransactionState
	pack.TransactionID = msg.TransactionID
	pack.TransactionProtocol = msg.TransactionProtocol
	// ***** 验证层 *****
	pack.VerifyType = msg.VerifyType
	pack.VerifyData = msg.VerifyData
	// ***** 转发层 *****
	pack.FromChain = msg.FromChain
	pack.ToChain = msg.ToChain
	pack.TransferState = msg.TransferState

	// 转为字节码
	data, err := json.Marshal(pack)
	if err != nil {
		info := utils.InfoError(err)
		return nil, fmt.Errorf(info)
	}

	return data, nil
}
