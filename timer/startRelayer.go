// coding:utf-8
// 启动网关转发逻辑
package relayer

import (
	chainmakercommon "CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"CPS/utils"
	"fmt"
	"strings"
)

// 启动网关转发逻辑
// @param chainname 链名
// @param client 客户端
// @param ch 全部错误信息通过这个通道返回
func (r *Relayer) startRelayer(
	chainname string,
	client cps_common.Client,
	ch chan error,
) {
	// 订阅消息
	contractname := ""
	eventsend := ""
	// 订阅转发层消息
	switch r.mpChainType[chainname] {
	case cps_common.ChainmakerChainType:
		// 长安链
		contractname = chainmakercommon.LAYER_TRANSFER
		eventsend = chainmakercommon.EVENT_TRANSFER_SEND
	default:
		err := fmt.Errorf(
			"invalid chainname:%s, type:%s",
			chainname, r.mpChainType[chainname],
		)
		utils.InfoError(err)
		ch <- err
		return
	}

	// 订阅发送
	chMsg := make(chan *cps_common.SubscribeMsg, 100)
	utils.InfoTips(client.GetChainname(), "订阅合约", contractname, eventsend, "事件")
	go client.SubscribeEvent(contractname, eventsend, cps_common.MSG_SEND, chMsg)

	// 处理每一个消息
	for {
		// 获取消息
		msg, ok := <-chMsg
		if !ok {
			info := fmt.Sprintf("收到不ok的消息:%+v\n", msg)
			utils.InfoWarning(info)
			continue
		}

		// 处理这个监听到的消息
		utils.InfoTips(chainname, "监听到", eventsend, "事件,并开始处理")
		go r.dealRelayerMsg(chainname, client, msg, ch)
	}
}

// 转发监听到的消息
// @param chainname 监听到这个消息的客户端名
// @param client 监听到这个消息的客户端
// @param msg 监听到的消息
// @param ch 错误信息通道
func (r *Relayer) dealRelayerMsg(
	chainname string, client cps_common.Client,
	msg *cps_common.SubscribeMsg, ch chan error,
) {
	// 获得源链和目的链以及对应的链数据包
	// 解析完成后msg对象里面的参数会被改变
	if err := client.ParseEvent(msg); err != nil {
		info := utils.InfoError(err)
		ch <- fmt.Errorf(info)
		return
	}
	fromchain := string(msg.FromChain)
	tochain := string(msg.ToChain)
	utils.Info(chainname, "处理监听消息", msg.ChainName, "fromchain", fromchain, "tochain", tochain)

	// 找到需要转发的目的客户端
	destClient := r.getTargetClient(chainname, fromchain, tochain)
	if destClient == nil {
		utils.InfoWarning(chainname, "出现了不应该转发的消息")
		return
	}
	utils.InfoTips(chainname, "监听到消息,转发到", destClient.GetChainname())

	// 解析成目的链调用合约的参数
	data, err := destClient.SubscribeMsg2ReceiveArgs(msg)
	if err != nil {
		info := utils.InfoError(err)
		ch <- fmt.Errorf(info)
		return
	}

	// 发送给目的链
	_, err = destClient.InvokeContract(
		cps_common.CONTRACT_NAME_PROTOCOL_AGGREGATOR,
		cps_common.FUNC_RECEIVE,
		data,
	)
	if err != nil {
		utils.InfoWarning("chainname", destClient.GetChainname())
		info := utils.InfoError(err)
		ch <- fmt.Errorf(info)
		return
	}
	info := cps_common.GetMsgInfo(msg)
	utils.InfoTips(chainname, "监听到消息", info, "发送给", destClient.GetChainname(), "成功")
}

// 根据源链和目的链判断需要发送给哪一个客户端
// @param eventChainname 事件从哪个客户端监听到的
// @param fromchain 消息源链
// @param tochain 消息目的链
// @return client 转发给这个客户端
func (r *Relayer) getTargetClient(eventChainname, fromchain, tochain string) cps_common.Client {
	// 如果目的端存在，且从中继链监听到的就直接转发到目的链
	if r.mpChainName[tochain] && r.mpIsRelayer[eventChainname] {
		return r.mpClient[tochain]
	}
	// 如果从端链收到消息,就转发给中继链
	if (r.mpChainName[eventChainname] && !r.mpIsRelayer[eventChainname]) && strings.EqualFold(eventChainname, fromchain) {
		// 转发到中继链
		for key := range r.mpClient {
			if r.mpIsRelayer[key] {
				return r.mpClient[key]
			}
		}
	}

	// 异常情况
	utils.InfoWarning(eventChainname, "消息没有找到目标客户端, from:", fromchain, "to:", tochain)
	return nil
}
