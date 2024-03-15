// 公证人验证
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 公证人验证
// @param verify_data 验证数据
// @return error err
func (t *VERIFY) notaryValid(verify_data []byte) error {
	// notary验证交易原始发起者的地址要是指定的之一
	sender, err := sdk.Instance.Origin()
	if err != nil {
		return fmt.Errorf("get origin address error:%s", err.Error())
	}

	// 判断这个公证人是否注册
	if !t.validNotary(sender) {
		return fmt.Errorf("valid %s failed", sender)
	}

	return nil
}

// 添加一个公证人
// @param address 公证人地址
// @param error err
func (t *VERIFY) addNotary() pb.Response {
	// 获取发起者
	address, err := sdk.Instance.Origin()
	if err != nil {
		return sdk.Error("get origin error:" + err.Error())
	}

	// 是公证人就保存他的value为特定字符
	if err := sdk.Instance.PutState(
		address, common.KEY_FIELD_NOTARY, common.VALUE_VERIFY_NOTARY,
	); err != nil {
		return sdk.Error("save notary " + address + " error:" + err.Error())
	}

	return sdk.Success(nil)
}

// 验证一个公证人
// @param address 待验证的地址
// @return bool 是否是注册了的公证人
func (t *VERIFY) validNotary(address string) bool {
	// 在数据库中查找公证人
	str, err := sdk.Instance.GetState(address, common.KEY_FIELD_NOTARY)
	if err != nil {
		sdk.Instance.EmitEvent(common.EVENT_WARNING, []string{
			"search notary", address, "error:", err.Error(),
		})
		return false
	}

	return str == common.VALUE_VERIFY_NOTARY
}
