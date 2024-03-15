// coding:utf-8
// 订阅事件
package client

import (
	cps_common "CPS/common"
	"CPS/utils"
	"context"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 监听事件
// @param cc 发起监听的客户端
// @param contractName 合约名
// @param eventName 事件名
// @param msg_type 消息类型,send or receive or info
// @param resultChan 保存事件的通道
func (c *Client) SubscribeEvent(
	contractName, eventName string,
	msg_type cps_common.MSG_TYPE,
	resultChan chan<- *cps_common.SubscribeMsg,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 订阅事件
	ec, err := c.ChainmakerSDK.SubscribeContractEvent(ctx, -1, -1, contractName, eventName)
	if err != nil {
		// 返回错误消息
		info := utils.InfoError(err)
		msg := utils.NewSucscribeMsg(
			cps_common.MSG_ERROR, c.Chainname, "",
			contractName, eventName, []byte(info),
		)
		resultChan <- msg
		return
	}

	// 遍历全部事件
	for {
		select {
		case event, ok := <-ec:
			// 获取通道数据失败
			if !ok || event == nil {
				info := utils.InfoError(fmt.Errorf("event不ok:%+v", event))
				msg := utils.NewSucscribeMsg(
					cps_common.MSG_WARNING, c.Chainname, "",
					contractName, eventName, []byte(info),
				)
				resultChan <- msg
				return
			}

			// 类型断言失败
			contractEvent, ok := event.(*common.ContractEventInfo)
			if !ok {
				info := utils.Info("event type convert error")
				msg := utils.NewSucscribeMsg(
					cps_common.MSG_ERROR, c.Chainname, "",
					contractName, eventName, []byte(info),
				)
				resultChan <- msg
				continue
			}

			// fmt.Printf("%s监听到事件:%+v\n", c.Chainname, contractEvent)

			// 转换为特定消息格式
			msg := event2Subevent(contractEvent, c.Chainname, eventName, contractName, msg_type)
			// 无论是错误消息还是跨链消息都发回去,由上层决定如何处理
			resultChan <- msg
		case <-ctx.Done():
			utils.InfoTips("上下文关闭,结束监听")
			close(resultChan)
			return
		}
	}
}
