// coding:utf-8
// 多链事务协同
package example

import (
	chainmaker_common "CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 测试资源相关程序
// @param client 客户端1
// @return error err
func TestResource(client cps_common.Client) error {
	// 1. 制作预锁资源
	swap := makeLockResource(client.GetChainname())
	// 2. 调用函数进行预锁
	utils.InfoTips("调用资源合约进行预锁")
	if err := callResource(client, swap, chainmaker_common.FUNC_RESOURCE_LOCK); err != nil {
		panic(err)
	}
	// 3. 获取锁定余额
	if err := getLockCount(client, cps_common.CONTRACT_NAME_TESTAPP, "a"); err != nil {
		panic(err)
	}
	// 4. 获取可用余额
	if err := getValidCount(client, cps_common.CONTRACT_NAME_TESTAPP, "a"); err != nil {
		panic(err)
	}

	// 5. 获取可用余额
	if err := getValidCount(client, cps_common.CONTRACT_NAME_TESTAPP, "c"); err != nil {
		panic(err)
	}
	// 6. 调用函数进行提交资源
	utils.InfoTips("调用资源合约进行提交")
	if err := callResource(client, swap, chainmaker_common.FUNC_RESOURCE_COMMIT); err != nil {
		panic(err)
	}
	// 7. 获取可用余额
	if err := getValidCount(client, cps_common.CONTRACT_NAME_TESTAPP, "c"); err != nil {
		panic(err)
	}

	return nil
}

// 获取Lockresource
// @param chainname 链名
// @return atomicswap swap
func makeLockResource(chainname string) *cps_common.AtomicSwap {
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
		ChainName:    cps_common.CHAIN_ID_CHAIN2,
		ResourceName: "c",
		Count:        1,
	}
	swap := cps_common.AtomicSwap{
		LockResource: []cps_common.Item{
			itema, itemb,
		},
		WantResource: []cps_common.Item{
			itemc,
		},
	}
	return &swap
}

// 调用resource合约
// @param client 客户端
// @param swap 涉及到的资源
// @param funcName 调用的功能
// @return error err
func callResource(client cps_common.Client, swap *cps_common.AtomicSwap, funcName string) error {

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
		{Key: "method", Value: []byte(funcName)},
		{Key: cps_common.KEY_MSG, Value: msgByte},
	}
	utils.InfoTips(client.GetChainname(), "调用资源合约", funcName)
	if _, err := client.InvokeContract(
		cps_common.CONTRACT_NAME_RESOURCE,
		"", kvs,
	); err != nil {
		panic(err)
	}
	utils.InfoTips(client.GetChainname(), "调用完成")
	return nil
}

// 获取可用余额
// @param clien 客户端
// @param app app
// @param name 资源名
// @return error err
func getValidCount(client cps_common.Client, app, name string) error {
	// ! 以事件的形式抛出来
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte(chainmaker_common.FUNC_RESOURCE_GET_VALID_COUNT)},
		{Key: chainmaker_common.KEY_FROM_APP, Value: []byte(app)},
		{Key: chainmaker_common.KEY_RESOURCE_USE_FROM, Value: []byte(name)},
	}
	utils.InfoTips(client.GetChainname(), "调用资源合约获取可用数量", app, name)
	if _, err := client.InvokeContract(
		cps_common.CONTRACT_NAME_RESOURCE,
		"", kvs,
	); err != nil {
		panic(err)
	}
	utils.InfoTips(client.GetChainname(), app, name, "获取可用数量调用完成")
	return nil
}

// 获取锁定余额数量
// @param clien 客户端
// @param app app
// @param name 资源名
// @return error err
func getLockCount(client cps_common.Client, app, name string) error {
	// ! 以事件的形式抛出来
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte(chainmaker_common.FUNC_RESOURCE_GET_LOCKCOUNT)},
		{Key: chainmaker_common.KEY_FROM_APP, Value: []byte(app)},
		{Key: chainmaker_common.KEY_RESOURCE_USE_FROM, Value: []byte(name)},
	}
	utils.InfoTips(client.GetChainname(), "调用资源合约获取锁定数量", app, name)
	if _, err := client.InvokeContract(
		cps_common.CONTRACT_NAME_RESOURCE,
		"", kvs,
	); err != nil {
		panic(err)
	}
	utils.InfoTips(client.GetChainname(), app, name, "获取锁定数量调用完成")
	return nil
}
