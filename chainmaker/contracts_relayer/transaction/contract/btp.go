// coding:utf-8
// btp保存和匹配相关资源
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"
)

// *************************** btp ***************************
// 执行BTP协议
func (t *TRANSACTION) saveMarchBtpResource(msg *common.ProtocolMsg) ([]byte, error) {
	t.emit_info([]string{"执行btp协议"})

	// 1. 获得want里面,对应链的全部swap
	// 每条链需要锁定的东西
	mpSwap := make(map[string]*cps_common.AtomicSwap)
	for _, item := range msg.AtomicSwap.WantResource {
		// 2. 制作为map
		chainname := item.ChainName
		// 如果链不存在就新建跨链资源
		if mpSwap[chainname] == nil {
			mpSwap[chainname] = &cps_common.AtomicSwap{
				WantResource: []cps_common.Item{},
			}
		}
		// 加入这条链的资源
		mpSwap[chainname].WantResource = append(
			mpSwap[chainname].WantResource,
			item,
		)
	}

	// 3. 字节化
	swapByte, err := json.Marshal(mpSwap)
	if err != nil {
		t.emit_warning([]string{"marshal btp swap state error", err.Error()})
		return nil, fmt.Errorf("marshal btp swap state error:" + err.Error())
	}
	return swapByte, nil
}
