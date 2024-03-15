// coding:utf-8
// 长安链客户端
package client

import (
	"CPS/utils"
	"fmt"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

// 封装的长安链客户端
type Client struct {
	// 长安链SDK客户端
	ChainmakerSDK *sdk.ChainClient
	// 链名
	Chainname string
	// 配置文件
	ConfigPath string
}

// 创建链客户端
// @param chainname 链名
// @param sdkConfPath:客户端sdk路径
// @return client 客户端对象
// @return error err
func NewChainmakerClient(chainname, sdkConfPath string) (*Client, error) {
	result := Client{}
	result.Chainname = chainname
	result.ConfigPath = sdkConfPath

	// 新建客户端
	cc, err := sdk.NewChainClient(
		sdk.WithConfPath(sdkConfPath),
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

	result.ChainmakerSDK = cc
	return &result, nil
}
