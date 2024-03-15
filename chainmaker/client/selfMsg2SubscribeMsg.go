// coding:utf-8
// 客户端内部消息转为订阅格式的消息
package client

import "CPS/common"

// 客户端内部消息转化为订阅MSG格式
// @param msg 客户端内部消息
// @return subscribemsg 订阅消息
// @return error 报错
func (c *Client) SelfMsg2SubscribeMsg(msg any) (*common.SubscribeMsg, error) {
	// todo:
	return nil, nil
}
