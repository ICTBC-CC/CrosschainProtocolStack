// coding:utf-8
// 部署合约
package deploy

import (
	cps_common "CPS/common"
	"CPS/utils"
	"errors"
	"fmt"
	"strconv"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
)

// 检查回执
func checkProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		if resp.Message == "" {
			resp.Message = resp.Code.String()
		}
		return fmt.Errorf(resp.Message)
	}
	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}
	// 0是成功代码
	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return fmt.Errorf(resp.ContractResult.Message)
	}
	return nil
}

// 自动根据客户端的连接模式获取背书节点
func getEndorsers(hashType crypto.HashType,
	authType sdk.AuthType,
	payload *common.Payload,
	usernames ...string,
) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	for _, name := range usernames {
		var entry *common.EndorsementEntry
		var err error
		users := cps_common.GetAdmins()
		switch authType {
		case sdk.PermissionedWithCert:
			u, ok := users[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakeEndorserWithPath(u.SignKeyPath, u.SignCrtPath, payload)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("invalid authType")
		}
		endorsers = append(endorsers, entry)
	}

	return endorsers, nil
}

// 输入全部参数进行合约部署
// @param chainmakerSDK 客户端
// @param contractName 合约名称
// @param version 合约版本
// @param byteCodePath 编译好的合约文件路径
// @param kvs: 合约初始化所需要的参数，以key-value形式传入
// @param withSyncResult: 是否需要同步获取交易结果
// @param usernames: 链节点数(需要节点背书同意)
func DeployContractWithArgs(
	chainmakerSDK *sdk.ChainClient,
	contractName string,
	version string,
	byteCodePath string,
	kvs []*common.KeyValuePair,
	createContractTimeout int64,
	usernames ...string,
) (*common.TxResponse, error) {
	// 查询合约是否存在
	contractInfo, err := chainmakerSDK.GetContractInfo(contractName)
	needContractResult := true

	var payload = &common.Payload{}
	if err == nil {
		// 合约存在就升级
		utils.InfoTips(contractName, "合约存在:")
		utils.Info(contractInfo)
		needContractResult = false
		newversion := getNextVersion(contractInfo.Version)
		payload, err = chainmakerSDK.CreateContractUpgradePayload(contractName, newversion, byteCodePath, common.RuntimeType_DOCKER_GO, kvs)
	} else {
		// 不存在就部署
		payload, err = chainmakerSDK.CreateContractCreatePayload(contractName, version, byteCodePath, common.RuntimeType_DOCKER_GO, kvs)
	}
	if err != nil {
		info := utils.InfoError(err)
		return nil, fmt.Errorf(info)
	}

	// 根据不同的模式，获取背书节点
	endorsers, err := getEndorsers(chainmakerSDK.GetHashType(),
		chainmakerSDK.GetAuthType(), payload, usernames...)
	if err != nil {
		return nil, fmt.Errorf("getEndorsers error: " + err.Error())
	}

	// 发送合约管理请求
	resp, err := chainmakerSDK.SendContractManageRequest(payload, endorsers, createContractTimeout, true)
	if err != nil {
		return resp, err
	}

	// 检查回执
	utils.Info("检查回执")
	err = checkProposalRequestResp(resp, needContractResult)
	if err != nil {
		return resp, err
	}

	// log
	utils.Print("deploy/upgrade contract success\n[code]:" + string(resp.Code) + "\n[message]:" + resp.Message + "\n[contractResult]:" + resp.ContractResult.String() + "\n[txid]:" + resp.TxId)
	return resp, nil
}

// 获取当前版本号的下一个版本号
// @param version 以前的版本
// @return string 下一个版本号
func getNextVersion(version string) string {
	v, err := strconv.Atoi(version)
	if err != nil {
		utils.InfoWarning("version:", version, "is not a number, return default version")
		return "0"
	}
	return strconv.Itoa(v + 1)
}
