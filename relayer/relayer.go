// coding:utf-8
// 链下网关
package relayer

import (
	"CPS/common"
	"CPS/utils"
	"fmt"
)

type Relayer struct {
	// 保存全部网关, chainname->client
	// ! 目前是每条链一个网关
	mpClient map[string]common.Client
	// 输入的配置文件,也保存下来
	mpConfig map[string]string
	// 每条链是不是中继链
	mpIsRelayer map[string]bool
	// 每条链的链类型
	mpChainType map[string]string
	// 当前网关拥有的链名
	mpChainName map[string]bool
}

// 新建网关
// @param lsChainConfig 每条链的配置文件
// @param lsChainName 每条链的链名
// @param lsChainType 每条链的类型
// @param lsIsrelayer 是否是中继链
// @return *relayer 中继对象
// @return error 错误信息
func NewRelayer(
	lsChainConfig, lsChainName, lsChainType []string,
	lsIsrelayer []bool,
) (*Relayer, error) {
	p := &Relayer{}
	// 实例化
	p.mpClient = make(map[string]common.Client)
	p.mpConfig = make(map[string]string)
	p.mpIsRelayer = make(map[string]bool)
	p.mpChainType = make(map[string]string)
	p.mpChainName = make(map[string]bool)

	if len(lsChainConfig) != len(lsChainName) {
		return nil, fmt.Errorf("length of config must be euqal to chainname")
	}
	if len(lsChainConfig) != len(lsChainType) {
		return nil, fmt.Errorf("length of config must be euqal to chaintype")
	}

	// 新建客户端
	for i := 0; i < len(lsChainConfig); i++ {
		config := lsChainConfig[i]
		chainname := lsChainName[i]
		chaintype := lsChainType[i]
		client, err := p.newClient(chaintype, config, chainname)
		if err != nil {
			info := utils.InfoError(err)
			return nil, fmt.Errorf(info)
		}
		p.mpClient[chainname] = client
		p.mpConfig[chainname] = config
		p.mpChainType[chainname] = chaintype
		p.mpIsRelayer[chainname] = lsIsrelayer[i]
		p.mpChainName[chainname] = true
		utils.InfoTips("实例化", client.GetChainname(), "网关对象成功")
	}

	return p, nil
}
