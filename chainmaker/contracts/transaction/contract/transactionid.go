// coding:utf-8
// 交易id相关东西
package contract

import (
	"CPS/chainmaker/contracts/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 获取当前最新的id
// @return TypeTransactionID 最新id
// @return error err
func (t *TRANSACTION) getNewestID() (common.TypeTransactionID, error) {
	// 读取数据库
	idBytes, err := sdk.Instance.GetStateByte(
		common.KEY_TRANSACTION_NEWEST_ID,
		common.KEY_FIELD_TRANSACTION_ID,
	)
	if err != nil {
		return common.NIL_TRANSACTION_ID, fmt.Errorf("get newest transaction id error:" + err.Error())
	}

	// 解码
	newestID := common.NIL_TRANSACTION_ID
	if err := json.Unmarshal(idBytes, &newestID); err != nil {
		return common.NIL_TRANSACTION_ID, fmt.Errorf("unmarshal newest transaction id error:" + string(idBytes) + err.Error())
	}

	return newestID, nil
}

// 增加当前最新ID
// @return error err
func (t *TRANSACTION) updateNesestID(newestID common.TypeTransactionID) error {
	// 获取最新ID
	if newestID == common.NIL_TRANSACTION_ID {
		// 没有输入当前id就从数据库寻找
		currID, err := t.getNewestID()
		if err != nil {
			return fmt.Errorf("get newest id error:" + err.Error())
		}
		newestID = currID
	}

	// 加一
	// ! 这里一定注意ID的类型一定需要有+=操作
	newestID += 1

	// 字节化
	idBytes, err := json.Marshal(newestID)
	if err != nil {
		return fmt.Errorf("marshal transaction id error:" + err.Error())
	}

	// 保存到数据库
	if err := sdk.Instance.PutStateByte(
		common.KEY_TRANSACTION_NEWEST_ID,
		common.KEY_FIELD_TRANSACTION_ID,
		idBytes,
	); err != nil {
		return fmt.Errorf("put newest transaction id state error:" + err.Error())
	}

	return nil
}

// 获取并增加当前最新交易id
// @return transactionid 事务id
// @return error err
func (t *TRANSACTION) getUpdateNewestID() (common.TypeTransactionID, error) {
	// 获取当前最新id
	currID, err := t.getNewestID()
	if err != nil {
		return common.NIL_TRANSACTION_ID, fmt.Errorf("get newest transaction id error:" + err.Error())
	}

	// 更新最新id
	if err := t.updateNesestID(currID); err != nil {
		return common.NIL_TRANSACTION_ID, fmt.Errorf("update newest transaction id error:" + err.Error())
	}

	return currID, nil
}
