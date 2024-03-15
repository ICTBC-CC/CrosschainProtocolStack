// coding:utf-8
// BTP协议的send
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"
)

// 执行BTP协议
// @param txID 事务id
// @param msg 消息
func (t *TRANSACTION) sendBtp(
	txID common.TypeTransactionID,
	msg *common.ProtocolMsg,
) error {
	// ! 目前send部分和MTP一样
	if err := t.sendMtp(txID, msg); err != nil {
		return fmt.Errorf("send btp error:" + err.Error())
	}
	return nil
}
