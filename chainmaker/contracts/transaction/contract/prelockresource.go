// coding:utf-8
// 预锁资源
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 预锁资源
// @param currID 当前事务id
// @param fromapp 源app
// @param toapp 目的app
// @param resource 锁定的资源
// @return []byte 锁定的资源id
// @return error err
func (t *TRANSACTION) preLockResource(msg *common.ProtocolMsg) error {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal msg error:" + err.Error())
	}

	// 参数
	kvs := map[string][]byte{cps_common.KEY_MSG: msgByte}

	// 调用资源合约进行预锁lock资源
	resp := sdk.Instance.CallContract(
		common.LAYER_RESOURCE,
		common.FUNC_RESOURCE_LOCK,
		kvs,
	)
	if resp.Status != sdk.OK {
		// 调用失败
		return fmt.Errorf("call lock resource error:" + resp.Message)
	}

	return nil
}
