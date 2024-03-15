// coding:utf-8
// 部署合约
package main

import (
	"CPS/chainmaker/deploy"
	cps_common "CPS/common"
	"CPS/utils"
	"fmt"
	"sync"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

// 新建长安链sdk客户端
// @param configPath 配置文件
// @return chainclient 长安链sdk客户端
// @return error err
func chainmakerClient(configPath string) (*sdk.ChainClient, error) {
	cc, err := sdk.NewChainClient(
		sdk.WithConfPath(configPath),
		sdk.WithEnableTxResultDispatcher(true),
	)
	if err != nil {
		info := utils.InfoError(err)
		return nil, fmt.Errorf(info)
	}
	if cc.GetAuthType() == sdk.PermissionedWithCert {
		if err := cc.EnableCertHash(); err != nil {
			info := utils.InfoError(err)
			return nil, fmt.Errorf(info)
		}
	}
	return cc, nil
}

// 封装一下合约部署的函数
// @param c 链客户端
// @param contractName 合约名
// @param contractPath 合约路径
// @param isrelayer 中继合约标志位
// param usernames 背书节点
func deployContract(
	c *sdk.ChainClient,
	contractName string,
	contractPath string,
	isrelayer bool,
	usernames ...string,
) error {
	// 中继链合约目录
	contractFolder := "contracts"
	if isrelayer {
		contractFolder = "contracts_relayer"
	}
	utils.Info("start to deploy", contractName)
	// 调用部署脚本
	aProposalArgs := []*common.KeyValuePair{}
	_, err := deploy.DeployContractWithArgs(
		c,
		contractName, "0",
		fmt.Sprintf("./chainmaker/%s/%s/%s.7z", contractFolder, contractPath, contractName),
		aProposalArgs,
		10,
		usernames...)
	if err != nil {
		info := utils.InfoError(err)
		return fmt.Errorf(info)
	}
	utils.Info("deploy", contractName, "contract success")
	return nil
}

// 部署合约
// @param configPath 配置文件路径
// @param isrelayer 中继链标志位
// @param usernames endorse
func deployChainmakerContract(configPath string, isrelayer bool, usernames ...string) error {
	// 新建客户端
	cc, err := chainmakerClient(configPath)
	if err != nil {
		info := utils.InfoError(err)
		panic(info)
	}

	// ******** 部署合约 ********
	var wg sync.WaitGroup

	// // 部署 transfer 合约
	// wg.Add(1)
	// go func() {
	if err := deployContract(cc, "transfer", "transfer", isrelayer, usernames...); err != nil {
		utils.InfoError(err)
		panic(err)
	}
	// 	wg.Done()
	// }()

	// // 部署 verify 合约
	// wg.Add(1)
	// go func() {
	if err := deployContract(cc, "verify", "verify", isrelayer, usernames...); err != nil {
		utils.InfoError(err)
		panic(err)
	}
	// 	wg.Done()
	// }()

	// // 部署 transaction 合约
	// wg.Add(1)
	// go func() {
	if err := deployContract(cc, "transaction", "transaction", isrelayer, usernames...); err != nil {
		utils.InfoError(err)
		panic(err)
	}
	// 	wg.Done()
	// }()

	// // 部署 protocolaggregator 合约
	// wg.Add(1)
	// go func() {
	if err := deployContract(cc, "protocolaggregator", "protocolaggregator", isrelayer, usernames...); err != nil {
		utils.InfoError(err)
		panic(err)
	}
	// 	wg.Done()
	// }()

	// 端链才有其他合约
	if isrelayer {
		wg.Wait()
		return nil
	}

	// // 部署 resource 合约
	// wg.Add(1)
	// go func() {
	if err := deployContract(cc, "resource", "resource", isrelayer, usernames...); err != nil {
		utils.InfoError(err)
		panic(err)
	}
	// 	wg.Done()
	// }()

	// // 部署 testapp 合约
	// wg.Add(1)
	// go func() {
	if err := deployContract(cc, "testapp", "app/testapp", isrelayer, usernames...); err != nil {
		utils.InfoError(err)
		panic(err)
	}
	// 	wg.Done()
	// }()

	wg.Wait()
	return nil
}

func main() {
	var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	utils.Info("开始部署chain1合约")
	// 配置文件
	chain1UserConfigPath := cps_common.MpConfig[cps_common.CHAIN_ID_CHAIN1]
	// 部署合约
	deployChainmakerContract(chain1UserConfigPath, false, cps_common.Chain1Admins...)
	// 	wg.Done()
	// }()

	// wg.Add(1)
	// go func() {
	utils.Info("开始部署chain2合约")
	// 配置文件
	chain2UserConfigPath := cps_common.MpConfig[cps_common.CHAIN_ID_CHAIN2]
	// 部署合约
	deployChainmakerContract(chain2UserConfigPath, false, cps_common.Chain2Admins...)
	// 	wg.Done()
	// }()

	// wg.Add(1)
	// go func() {
	utils.Info("开始部署chain3合约")
	// 链3是中继链
	chain3UserConfigPath := cps_common.MpConfig[cps_common.CHAIN_ID_CHAIN3]
	// 部署合约
	deployChainmakerContract(chain3UserConfigPath, false, cps_common.Chain3Admins...)
	// 	wg.Done()
	// }()

	// wg.Add(1)
	// go func() {
	utils.Info("开始部署relayer合约")
	// 链3是中继链
	chain4UserConfigPath := cps_common.MpConfig[cps_common.CHAIN_ID_RELAYER]
	// 部署合约
	deployChainmakerContract(chain4UserConfigPath, true, cps_common.ChainRelayerAdmins...)
	// 	wg.Done()
	// }()

	wg.Wait()
}
