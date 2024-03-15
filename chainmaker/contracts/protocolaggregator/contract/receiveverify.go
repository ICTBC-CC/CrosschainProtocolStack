// coding:utf-8
// 执行验证层receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用验证层receive
// @param verify_type 验证层类型
// @param verify_data 验证层数据
// @param error err
func (t *ProtocolAggregator) receiveVerify(verify_type string, verify_data []byte) error {
	kvs := map[string][]byte{
		common.KEY_VERIFY_TYPE: []byte(verify_type),
		common.KEY_VERIFY_DATA: verify_data,
	}
	// 调用转发层
	resp := sdk.Instance.CallContract(
		common.LAYER_VERIFY,
		common.FUNC_RECEIVE,
		kvs,
	)
	// 转发层调用失败就返回错误
	if resp.Status != sdk.OK {
		return fmt.Errorf("call verify receive error:%s", resp.String())
	}
	return nil
}
