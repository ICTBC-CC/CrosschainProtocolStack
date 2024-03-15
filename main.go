// coding:utf-8
// 主函数
package main

import (
	chainmaker_client "CPS/chainmaker/client"
	cps_common "CPS/common"
	"CPS/example"
	"CPS/relayer"
	"CPS/utils"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"sync"
)

// 链以及对应的证书路径
var mpConfig = cps_common.MpConfig

// 动作信号
var sign string

// 目标链
var tochain string

func args_parse() {
	flag.StringVar(&sign, "t", "sendmsg2peer", "sendmsg2peer:发送消息到对端链,1:执行测试,relayer:启动网关")
	flag.Parse()
}

func main() {
	// 参数解析
	args_parse()
	var wg sync.WaitGroup
	switch sign {
	case "tx":
		{
			// 获取交易
			tx_id := "17b6cdf036f24c68ca1bbd0d5fef429727ab8602954d4e3598aa712601f95f36"
			client, err := chainmaker_client.NewChainmakerClient(
				cps_common.CHAIN_ID_CHAIN1,
				mpConfig[cps_common.CHAIN_ID_CHAIN1],
			)
			if err != nil {
				panic(err)
			}
			tx, err := client.ChainmakerSDK.GetTxByTxId(tx_id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v", tx)
		}
	case "1":
		{
			// 测试长安链
			// test.Run("chain1", chain1UserConfigPath)
		}
	case "btp":
		{
			// 执行发送消息到对端链
			utils.InfoTips("新建chain1用户客户端")
			client, err := newClient(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN1,
				mpConfig[cps_common.CHAIN_ID_CHAIN1],
			)
			if err != nil {
				panic(err)
			}
			utils.InfoTips("客户端新建完成, 执行BTP测试", "由", client.GetChainname(), "与", cps_common.CHAIN_ID_CHAIN2, "通信")
			if err := example.TestBtp(client); err != nil {
				panic(err)
			}
		}
	case "testapp":
		{
			// 执行发送消息到对端链
			client, err := newClient(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN1,
				mpConfig[cps_common.CHAIN_ID_CHAIN1],
			)
			if err != nil {
				panic(err)
			}
			if err := example.Testapp(client, "send a msg", tochain); err != nil {
				panic(err)
			}
		}
	case "mtp":
		{
			utils.InfoTips("开始执行mtp测试")
			testMTP()
		}
	case "relayer":
		{
			utils.Info("启动relayer")
			wg.Add(1)
			go run_relayer()
		}
	case "testresource":
		{
			utils.Info("开始测试资源合约")
			// 执行发送消息到对端链
			client, err := newClient(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN1,
				mpConfig[cps_common.CHAIN_ID_CHAIN1],
			)
			if err != nil {
				panic(err)
			}
			if err := example.TestResource(client); err != nil {
				panic(err)
			}
		}
	case "testresourcecount":
		{
			utils.Info("开始测试资源合约数量")
			// 执行发送消息到对端链
			client1, err := newClient(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN1,
				mpConfig[cps_common.CHAIN_ID_CHAIN1],
			)
			if err != nil {
				panic(err)
			}
			client2, err := newClient(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN2,
				mpConfig[cps_common.CHAIN_ID_CHAIN2],
			)
			if err != nil {
				panic(err)
			}
			client3, err := newClient(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN3,
				mpConfig[cps_common.CHAIN_ID_CHAIN3],
			)
			if err != nil {
				panic(err)
			}
			if err := example.TestResourceCount(
				[]cps_common.Client{client1, client2, client3},
			); err != nil {
				panic(err)
			}
		}
	case "info":
		{
			// 监听info事件
			wg.Add(1)
			utils.Info("启动事件监听器")
			// 链1事件
			utils.InfoTips("监听 chain1 info事件")
			go subscribeInfo(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN1,
				mpConfig[cps_common.CHAIN_ID_CHAIN1],
			)
			// 链2事件
			utils.InfoTips("监听 chain2 info事件")
			go subscribeInfo(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN2,
				mpConfig[cps_common.CHAIN_ID_CHAIN2],
			)
			// 链3事件
			utils.InfoTips("监听 chain3 info事件")
			go subscribeInfo(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_CHAIN3,
				mpConfig[cps_common.CHAIN_ID_CHAIN3],
			)
			// relayer链事件
			utils.InfoTips("监听 relayer chain info事件")
			go subscribeInfo(
				cps_common.ChainmakerChainType,
				cps_common.CHAIN_ID_RELAYER,
				mpConfig[cps_common.CHAIN_ID_RELAYER],
			)
		}
	default:
		{
			// 异常情况
			panic("invalid parameter")
		}
	}
	wg.Wait()
}

// 测试执行MTP事务
func testMTP() error {
	utils.InfoTips("新建", cps_common.CHAIN_ID_CHAIN1, "客户端")
	client1, err := newClient(
		cps_common.ChainmakerChainType,
		cps_common.CHAIN_ID_CHAIN1,
		mpConfig[cps_common.CHAIN_ID_CHAIN1],
	)
	if err != nil {
		panic(err)
	}

	utils.InfoTips("新建", cps_common.CHAIN_ID_CHAIN2, "客户端")
	client2, err := newClient(
		cps_common.ChainmakerChainType,
		cps_common.CHAIN_ID_CHAIN2,
		mpConfig[cps_common.CHAIN_ID_CHAIN2],
	)
	if err != nil {
		panic(err)
	}

	utils.InfoTips("新建", cps_common.CHAIN_ID_CHAIN3, "客户端")
	client3, err := newClient(
		cps_common.ChainmakerChainType,
		cps_common.CHAIN_ID_CHAIN3,
		mpConfig[cps_common.CHAIN_ID_CHAIN3],
	)
	if err != nil {
		panic(err)
	}

	// 执行测试代码
	if err := example.TestMtp(
		client1, client2, client3,
	); err != nil {
		panic(err)
	}
	return nil
}

// 监听info事件
func subscribeInfo(chainType, chainName, configPath string) {
	msg_chan := make(chan *cps_common.SubscribeMsg, 100)

	// client
	client, err := newClient(chainType, chainName, configPath)
	if err != nil {
		panic(err)
	}
	utils.InfoTips("开始监听", chainName, "事件")
	go client.SubscribeEvent("protocolaggregator", cps_common.EVENT_INFO, cps_common.MSG_INFO, msg_chan)
	go client.SubscribeEvent("protocolaggregator", cps_common.EVENT_WARNING, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("protocolaggregator", cps_common.EVENT_ERROR, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("transaction", cps_common.EVENT_INFO, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("transaction", cps_common.EVENT_WARNING, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("transaction", cps_common.EVENT_ERROR, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("resource", cps_common.EVENT_INFO, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("resource", cps_common.EVENT_WARNING, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent("resource", cps_common.EVENT_ERROR, cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent(
		cps_common.CONTRACT_NAME_TESTAPP,
		cps_common.EVENT_WARNING,
		cps_common.MSG_WARNING, msg_chan)
	go client.SubscribeEvent(
		cps_common.CONTRACT_NAME_TESTAPP,
		cps_common.EVENT_INFO,
		cps_common.MSG_INFO, msg_chan)

	for {
		msg, ok := <-msg_chan
		if !ok {
			utils.InfoWarning("监听不ok")
			continue
		}
		// 解码
		ls_data := &[]string{}
		if err := json.Unmarshal(msg.ChainData, ls_data); err != nil {
			utils.InfoWarning("解码监听事件错误:", err)
			continue
		}
		// 输出
		info := fmt.Sprintf(msg.ChainName, msg.ContractName, msg.EventName, ls_data)
		if strings.Contains(info, "warning") {
			utils.InfoWarning(info)
		} else if strings.Contains(info, "error") {
			utils.InfoError(fmt.Errorf(info))
		} else {
			utils.Info("msg", info)
		}
	}
}

// 新建客户端
// @param chaintype 链类型
// @param config_path 配置文件
// @param chain_name 链名
// @return client 客户端
// @return error err
func newClient(chaintype, chain_name, config_path string) (cps_common.Client, error) {
	switch chaintype {
	case cps_common.ChainmakerChainType:
		// 长安链
		client, err := chainmaker_client.NewChainmakerClient(chain_name, config_path)
		if err != nil {
			info := utils.InfoError(err)
			return nil, fmt.Errorf(info)
		}
		return client, nil
	}
	return nil, fmt.Errorf("invalid chain name")
}

// 执行中继
// @param chain1type 链1类型
// @param config1Path 链1配置文件路径
// @param chain1Name 链1名字
// @param isrelayer1 链1是否是中继链
// @param chain2type 链2类型
// @param config2Path 链2配置文件路径
// @param chain2Name 链2名字
// @param isrelayer2 链2是否是中继链
// @param wg waitgroup
func run_relayer() {
	lsConfig := []string{
		cps_common.MpConfig[cps_common.CHAIN_ID_CHAIN1],
		cps_common.MpConfig[cps_common.CHAIN_ID_CHAIN2],
		cps_common.MpConfig[cps_common.CHAIN_ID_CHAIN3],
		cps_common.MpConfig[cps_common.CHAIN_ID_RELAYER],
	}
	lsChainName := []string{
		cps_common.CHAIN_ID_CHAIN1,
		cps_common.CHAIN_ID_CHAIN2,
		cps_common.CHAIN_ID_CHAIN3,
		cps_common.CHAIN_ID_RELAYER,
	}
	lsChainType := []string{
		cps_common.ChainmakerChainType,
		cps_common.ChainmakerChainType,
		cps_common.ChainmakerChainType,
		cps_common.ChainmakerChainType,
	}
	lsIsrelayer := []bool{false, false, false, true}
	// 启动relayer
	relayer, err := relayer.NewRelayer(
		lsConfig, lsChainName, lsChainType, lsIsrelayer,
	)
	if err != nil {
		panic("start relayer error:" + err.Error())
	}
	chErr := make(chan error, 100)
	go relayer.Run(chErr)
	// 等待完成
	for {
		err, ok := <-chErr
		if !ok {
			utils.InfoWarning("收到relayer通道消息,发生错误")
			continue
		}

		// 输出err信息
		utils.InfoWarning("relayer收到一个err")
		utils.InfoError(err)

		if strings.Contains(err.Error(), "panic") {
			break
		}
	}
}
