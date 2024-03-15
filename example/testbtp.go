// coding:utf-8client1
// 基本事务协议
// 仅进行消息流通
package example

import (
	chainmaker_common "CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 基本事务协议,仅进行信息交互
// @return error err
func TestBtp(client cps_common.Client) error {
	// ! 要启动relayer执行转发程序

	// 1. 参数
	msg := chainmaker_common.ProtocolMsg{}
	msg.AppData = []byte("this is a btp test data")
	msg.FromApp = []byte(cps_common.CONTRACT_NAME_TESTAPP)
	msg.AppFunc = []byte(cps_common.FUNC_BTP)
	msg.FromChain = []byte(client.GetChainname())
	msg.VerifyType = cps_common.VERIFY_TYPE_HAPPY
	msg.TransactionProtocol = cps_common.TRANSACTION_PROTOCOL_TYPE_BTP
	msg.AtomicSwap = makeBtpSwap(cps_common.CHAIN_ID_CHAIN1, cps_common.CHAIN_ID_CHAIN2, cps_common.CHAIN_ID_CHAIN3)

	utils.InfoTips("定义基本参数", fmt.Sprintf("%+v", msg))

	msgByte, err := json.Marshal(msg)
	if err != nil {
		info := utils.InfoError(err)
		panic(info)
	}

	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte(cps_common.FUNC_SEND)},
		{Key: cps_common.KEY_MSG, Value: msgByte},
	}

	// 2. 调用协议栈
	utils.InfoTips(client.GetChainname(), "发起BTP交易")
	if _, err := client.InvokeContract(
		cps_common.CONTRACT_NAME_PROTOCOL_AGGREGATOR,
		"", kvs,
	); err != nil {
		panic(err)
	}

	utils.InfoTips(client.GetChainname(), "交易上链成功")
	return nil
}

// 制作BTP交换的资源数据
// @param oriChain 源链
// @param dest1Chain 目的链1
// @param dest2Chain 目的链2
// @return atomicswap 每条链锁定的资源
func makeBtpSwap(oriChain, dest1Chain, dest2Chain string) *cps_common.AtomicSwap {
	utils.Info(oriChain, "使用a资源10个", "换取", dest1Chain, "的b资源10个", "和", dest2Chain, "的c资源1个")
	// 链1
	itema := cps_common.Item{
		ChainName:    oriChain,
		ResourceName: "a",
		Count:        10,
	}
	itemb := cps_common.Item{
		ChainName:    dest1Chain,
		ResourceName: "b",
		Count:        10,
	}
	itemc := cps_common.Item{
		ChainName:    dest2Chain,
		ResourceName: "c",
		Count:        1,
	}
	return &cps_common.AtomicSwap{
		LockResource: []cps_common.Item{
			itema,
		},
		WantResource: []cps_common.Item{
			itemc, itemb,
		},
	}
}
