// coding:ut-8
// 新建各个链的客户端
package relayer

import (
	cps_client "CPS/chainmaker/client"
	"CPS/common"
	cps_common "CPS/common"
	"CPS/utils"
	"fmt"
)

// 新建客户端
// @param chaintype 链类型
// @param config_path 配置文件
// @param chain_name 链名
// @return client 客户端
// @return error err
func (r *Relayer) newClient(chaintype, config_path, chain_name string) (common.Client, error) {
	switch chaintype {
	case cps_common.ChainmakerChainType:
		// 长安链
		client, err := cps_client.NewChainmakerClient(chain_name, config_path)
		if err != nil {
			info := utils.InfoError(err)
			return nil, fmt.Errorf(info)
		}
		return client, nil
	}
	return nil, fmt.Errorf("invalid chain name")
}
