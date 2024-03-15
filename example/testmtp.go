// coding:utf-8
// 多链事务协同
package example

import (
	chainmaker_common "CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 多链事务协同
// @param client1 客户端1
// @param client2 客户端2
// @param client3 客户端3
// @return error err
func TestMtp(
	client1, client2, client3 cps_common.Client,
) error {
	// ! 要启动relayer执行转发程序

	// 多链哈希锁
	hashlock := []byte(fmt.Sprintf("%d", utils.GetTimestamp()))
	// 多链时间锁, 10秒后超时
	timelock := utils.GetTimestamp() + 10
	utils.InfoTips("哈希锁:", string(hashlock), " 时间锁:", fmt.Sprintf("%d", timelock))

	// 链1
	utils.InfoTips(cps_common.CHAIN_ID_CHAIN1, "发起MTP交易", "hash锁:", string(hashlock), "时间锁:", timelock)
	if err := startMtp(client1, hashlock, timelock); err != nil {
		panic(err)
	}

	// 链2
	utils.InfoTips(cps_common.CHAIN_ID_CHAIN2, "发起MTP交易", "hash锁:", string(hashlock), "时间锁:", timelock)
	if err := startMtp(client2, hashlock, timelock); err != nil {
		panic(err)
	}

	// 链3
	utils.InfoTips(cps_common.CHAIN_ID_CHAIN3, "发起MTP交易", "hash锁:", string(hashlock), "时间锁:", timelock)
	if err := startMtp(client3, hashlock, timelock); err != nil {
		panic(err)
	}

	return nil
}

// 在自己的链上发起一个多链交换交易
func startMtp(
	client cps_common.Client,
	hashlock []byte,
	timelock int64,
) error {
	// 1. 制作自己需要转移的资源以及参数
	swap := makeAtomicSwap(client.GetChainname(), hashlock, timelock)

	msg := chainmaker_common.ProtocolMsg{}
	msg.AppData = []byte("this is a mtp test data")
	msg.FromApp = []byte(cps_common.CONTRACT_NAME_TESTAPP)
	msg.AppFunc = []byte(cps_common.FUNC_MTP)
	msg.FromChain = []byte(client.GetChainname())
	msg.ToChain = []byte(cps_common.CHAIN_ID_RELAYER)
	msg.VerifyType = cps_common.VERIFY_TYPE_HAPPY
	msg.TransactionProtocol = cps_common.TRANSACTION_PROTOCOL_TYPE_MTP
	msg.AtomicSwap = swap

	msgByte, err := json.Marshal(msg)
	if err != nil {
		info := utils.InfoError(err)
		panic(info)
	}

	// 2. 调用聚合器
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte(cps_common.FUNC_SEND)},
		{Key: cps_common.KEY_MSG, Value: msgByte},
	}
	utils.InfoTips(client.GetChainname(), "调用聚合器")
	if _, err := client.InvokeContract(
		cps_common.CONTRACT_NAME_PROTOCOL_AGGREGATOR,
		"", kvs,
	); err != nil {
		panic(err)
	}
	utils.InfoTips(client.GetChainname(), "调用完成")
	return nil
}

// 根据链名返回需要交换的资源
// @param chainname 链名
func makeAtomicSwap(
	chainname string,
	hashlock []byte,
	timelock int64,
) *cps_common.AtomicSwap {
	// 链1转移a资源10个,b资源10个, 换取c资源1个
	// 链2转移c资源1个,b资源10个, 换取d资源1个
	// 链3转移,d资源1个, 换取a资源10个,b资源20个
	switch chainname {
	case cps_common.CHAIN_ID_CHAIN1:
		{
			utils.Info(cps_common.CHAIN_ID_CHAIN1, "使用a资源10个", "b资源10个", "换取c资源1个")
			// 链1
			itema := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "a",
				Count:        10,
			}
			itemb := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "b",
				Count:        10,
			}
			itemc := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "c",
				Count:        1,
			}
			return &cps_common.AtomicSwap{
				HashLock: hashlock,
				TimeLock: timelock,
				LockResource: []cps_common.Item{
					itema, itemb,
				},
				WantResource: []cps_common.Item{
					itemc,
				},
			}
		}
	case cps_common.CHAIN_ID_CHAIN2:
		{
			// 链2
			utils.Info(cps_common.CHAIN_ID_CHAIN2, "使用b资源10个", "c资源1个", "换取d资源1个")
			itemb := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "b",
				Count:        10,
			}
			itemc := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "c",
				Count:        1,
			}
			itemd := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "d",
				Count:        1,
			}
			return &cps_common.AtomicSwap{
				HashLock: hashlock,
				TimeLock: timelock,
				LockResource: []cps_common.Item{
					itemb, itemc,
				},
				WantResource: []cps_common.Item{
					itemd,
				},
			}
		}
	case cps_common.CHAIN_ID_CHAIN3:
		{
			utils.Info(cps_common.CHAIN_ID_CHAIN3, "使用d资源1个", "换取a资源10个", "b资源20个")
			// 链3
			itemd := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "d",
				Count:        1,
			}
			itema := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "a",
				Count:        10,
			}
			itemb := cps_common.Item{
				ChainName:    chainname,
				ResourceName: "b",
				Count:        20,
			}
			return &cps_common.AtomicSwap{
				HashLock: hashlock,
				TimeLock: timelock,
				LockResource: []cps_common.Item{
					itemd,
				},
				WantResource: []cps_common.Item{
					itema, itemb,
				},
			}
		}
	default:
		{
			utils.InfoWarning("无效的链名")
			return nil
		}
	}
}
