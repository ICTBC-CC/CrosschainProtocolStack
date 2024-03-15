// coding:utf-8
// MTP协议的send
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 执行BTP协议
// @param txID 事务id
// @param msg 消息
func (t *TRANSACTION) sendMtp(
	txID common.TypeTransactionID,
	msg *common.ProtocolMsg,
) error {
	// 预锁资源并保存到数据库
	if err := t.preLockResource(msg); err != nil {
		return fmt.Errorf("pre lock resouce error:" + err.Error())
	}

	// 保存到数据库
	if err := sdk.Instance.PutState(
		fmt.Sprintf("%d", txID),
		common.KEY_FIELD_TRANSACTION_ID,
		common.VALUE_TRANSACTION_EXIST,
	); err == nil {
		// 保存成功
		return nil
	}

	// 保存失败就要回滚资源锁定
	if err := t.rollbackResource(msg); err != nil {
		return fmt.Errorf(
			"rollback error: %s",
			// fmt.Sprintf("%+v", msg.AtomicSwap),
			err.Error(),
		)
	}

	return fmt.Errorf("save transaction resource error")
}
