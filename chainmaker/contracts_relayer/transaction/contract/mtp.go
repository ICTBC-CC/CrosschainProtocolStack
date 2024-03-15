// coding:utf-8
// mtp保存和匹配相关资源
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 执行MTP协议
// @param msg 传入的msg
// @return []byte 匹配完成返回的各个链的资源
// @return error err
func (t *TRANSACTION) saveMarchMtpResource(msg *common.ProtocolMsg) ([]byte, error) {
	t.emit_info([]string{"执行mtp协议"})

	// todo:匹配hashlock
	// ! 目前这里是对hashlock进行计数,是3的倍数就表示完成
	hashlock := fmt.Sprintf("%x", msg.AtomicSwap.HashLock)
	hashcountByte, _ := sdk.Instance.GetStateFromKeyByte(hashlock)
	var hashcount int
	if len(hashcountByte) == 0 {
		// 表示不存在
		hashcount = 0
	} else {
		// 恢复当前数字
		hashcount = int(hashcountByte[0])
	}

	// 增加一个
	hashcount += 1

	// 匹配mtp资源
	flag, mpSwap, err := t.mtpMarchState(hashlock, msg)
	if err != nil {
		t.emit_warning([]string{"march state error:", err.Error()})
		return nil, fmt.Errorf("march state error:" + err.Error())
	}

	// 判断是否匹配完成,这里以达到三个数表示匹配完成, 应该是用flag
	flag = false
	if hashcount%3 == 0 {
		// 表示已经有三个了就返回true
		flag = true
		hashcount = 0
	}

	// 重新保存hashcount
	sdk.Instance.PutStateFromKeyByte(hashlock, []byte{byte(hashcount)})

	if flag {
		// 匹配完成, 返回每条链的swap
		swapByte, err := json.Marshal(mpSwap)
		if err != nil {
			t.emit_warning([]string{"marshal mpswap error", err.Error()})
			return nil, fmt.Errorf("marshal mpswap error:" + err.Error())
		}

		// 作为payload返回
		return swapByte, nil
	}
	// 匹配不完成,不管
	return nil, nil
}

// 获取资源匹配状态
// @param hashlock 哈希锁
// @param swap 当前传入的资源
// @return bool 是否匹配完成
// @return swap 匹配完成时,每条链的资源
// @return error err
func (t *TRANSACTION) mtpMarchState(
	hashlock string,
	msg *common.ProtocolMsg,
) (bool, map[string]*cps_common.AtomicSwap, error) {
	// 获取已有的资源
	swapBytes, err := sdk.Instance.GetStateByte(hashlock, common.MARCH_FIELD_NAME)
	if err != nil {
		t.emit_warning([]string{"get march state error:", err.Error()})
		return false, nil, fmt.Errorf("get march state error:" + err.Error())
	}

	// unmarshal
	mpSwap := make(map[string]*cps_common.AtomicSwap)
	if len(swapBytes) != 0 {
		// 表示存在,就unmarshal
		if err := json.Unmarshal(swapBytes, &mpSwap); err != nil {
			t.emit_warning([]string{"unmarshal swap byte error:", err.Error()})
			return false, nil, fmt.Errorf("unmarshal swap byte error:" + err.Error())
		}
	}

	// 保存当前资源到状态里面
	mpSwap[string(msg.FromChain)] = msg.AtomicSwap

	// 保存到数据库
	stateByte, err := json.Marshal(mpSwap)
	if err != nil {
		t.emit_warning([]string{"marshal state error", err.Error()})
		return false, nil, fmt.Errorf("marshal state error:" + err.Error())
	}
	if err := sdk.Instance.PutStateByte(hashlock, common.MARCH_FIELD_NAME, stateByte); err != nil {
		t.emit_warning([]string{"put state state error", err.Error()})
		return false, nil, fmt.Errorf("put state state error:" + err.Error())
	}

	// todo:判断是否匹配成功

	return false, mpSwap, nil
}
