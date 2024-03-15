// coding:utf-8
// 调用资源进行回滚
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 回滚锁定的资源
// @param resource 需要回滚的资源
// @return error err
func (t *TRANSACTION) rollbackResource(msg *common.ProtocolMsg) error {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("roolback marshal msg error:" + err.Error())
	}
	kvs := map[string][]byte{cps_common.KEY_MSG: msgByte}

	// 调用回滚
	resp := sdk.Instance.CallContract(
		common.LAYER_RESOURCE,
		common.FUNC_RESOURCE_ROLLBACK,
		kvs,
	)

	if resp.Status != sdk.OK {
		// 回滚失败
		return fmt.Errorf("call resource rollback error:" + resp.Message)
	}

	return nil
}
