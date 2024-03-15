// coding:utf-8
// 验证层接收到receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层接收到消息
func (t *VERIFY) receive() pb.Response {
	// 获取参数
	args := sdk.Instance.GetArgs()
	verify_type := args[common.KEY_VERIFY_TYPE]
	verify_data := args[common.KEY_VERIFY_DATA]

	// 根据类型做对应的验证
	if err := t.verify(verify_type, verify_data); err != nil {
		return sdk.Error(err.Error())
	}

	// 抛出数据
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"verify receive success"})
	return sdk.Success(nil)
}

// 根据不同类型做验证
// @param verify_type 验证类型
// @param verify_data 验证用到的数据
func (t *VERIFY) verify(verify_type, verify_data []byte) error {
	verify_str := string(verify_type)
	// 判断
	switch verify_str {
	case string(common.VERIFY_TYPE_HAPPY):
		{
			return t.happyValid(verify_data)
		}
	case string(common.VERIFY_TYPE_NOTARY):
		{
			return t.notaryValid(verify_data)
		}
	case string(common.VERIFY_TYPE_SPV):
		{
			return t.spvValid(verify_data)
		}
	}
	return fmt.Errorf("invalid verify type:%s", verify_str)
}
