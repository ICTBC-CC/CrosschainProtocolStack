// coding:utf-8
// 从订阅的消息里面解析出目的链
package client

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"
	"fmt"
)

// 从订阅的MSG中解析出链上协议栈消息
// @param msg 需要被解析的数据格式,通过指针改变对象
// @return error 错误消息
func (c *Client) ParseEvent(msg *cps_common.SubscribeMsg) error {
	// 对data进行解码,得到抛出的数据
	ls_data := []string{}
	if err := json.Unmarshal(msg.ChainData, &ls_data); err != nil {
		// 解析失败
		info := utils.InfoError(err)
		// return "", "", nil, fmt.Errorf(info)
		return fmt.Errorf(info)
	}

	// fmt.Println("长安链解析:", ls_data)
	// ! 解析出来应该是三个数据,分别是 源链,目的链,整个跨链数据包
	if len(ls_data) != 3 {
		// return "", "", nil, fmt.Errorf("length of chainmaker msg must be equal to 2")
		utils.InfoWarning(ls_data)
		return fmt.Errorf("length of chainmaker msg must be equal to 3")
	}
	// return ls_data[0], ls_data[1], []byte(ls_data[2]), nil

	// 对整个跨链数据包进行解码
	pack := &common.ProtocolMsg{}
	if err := json.Unmarshal([]byte(ls_data[2]), pack); err != nil {
		// 解码失败
		info := utils.InfoError(err)
		return fmt.Errorf(info)
	}

	// 赋值
	// ***** 资源 *****
	msg.AtomicSwap = pack.AtomicSwap
	// ***** 应用层 *****
	msg.AppFrom = pack.FromApp
	msg.AppFunc = pack.AppFunc
	msg.AppData = pack.AppData
	// ***** 事务层 *****
	msg.TransactionID = pack.TransactionID
	msg.TransactionData = pack.TransactionData
	msg.TransactionState = pack.TransactionState
	msg.TransactionProtocol = pack.TransactionProtocol
	// ***** 验证层 *****
	msg.VerifyType = pack.VerifyType
	msg.VerifyData = pack.VerifyData
	// ***** 转发层 *****
	msg.FromChain = pack.FromChain
	msg.ToChain = pack.ToChain
	msg.TransferState = pack.TransferState

	return nil
}
