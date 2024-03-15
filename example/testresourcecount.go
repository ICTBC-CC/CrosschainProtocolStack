// coding:utf-8
// 多链事务协同
package example

import (
	cps_common "CPS/common"
)

// 获取测试的所有资源
// @param client 客户端1
// @return error err
func TestResourceCount(ls_client []cps_common.Client) error {
	for _, client := range ls_client {
		// 获取锁定余额
		if err := getLockCount(client, cps_common.CONTRACT_NAME_TESTAPP, "a"); err != nil {
			panic(err)
		}
		// 获取锁定余额
		if err := getLockCount(client, cps_common.CONTRACT_NAME_TESTAPP, "b"); err != nil {
			panic(err)
		}
		// 获取锁定余额
		if err := getLockCount(client, cps_common.CONTRACT_NAME_TESTAPP, "c"); err != nil {
			panic(err)
		}

		// 获取可用余额
		if err := getValidCount(client, cps_common.CONTRACT_NAME_TESTAPP, "a"); err != nil {
			panic(err)
		}
		// 获取可用余额
		if err := getValidCount(client, cps_common.CONTRACT_NAME_TESTAPP, "b"); err != nil {
			panic(err)
		}
		// 获取可用余额
		if err := getValidCount(client, cps_common.CONTRACT_NAME_TESTAPP, "c"); err != nil {
			panic(err)
		}

	}

	return nil
}
