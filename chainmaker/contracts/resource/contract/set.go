// coding:utf-8
// 设置值
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 设置可用资源数量
// @param app 应用名
// @param name 资源名
// @param count 资源数量
// @return error err
func (r *Resource) setValidCount(app, name string, count int) error {
	// 数值判断
	if count < 0 {
		return fmt.Errorf("count must be greater than 0")
	}
	return r.setFieldCount(common.VALID_FIELD_NAME, app, name, count)
}

// 设置锁定资源数量
// @param app 应用名
// @param name 资源名
// @param count 资源数量
// @return error err
func (r *Resource) setLockCount(app, name string, count int) error {
	if count < 0 {
		return fmt.Errorf("set lock resource count must be greater than 0")
	}
	return r.setFieldCount(common.LOCK_FIELD_NAME, app, name, count)
}

// 设置域中资源数量h
// @param field 域名
// @param app 应用名
// @param name 资源名
// @return int 资源数量
// @return error err
func (r *Resource) setFieldCount(field, app, name string, count int) error {
	// 资源key值
	key, err := r.makeKey(app, name)
	if err != nil {
		return err
	}

	countStr := fmt.Sprintf("%d", count)

	// 设置域中资源数量
	return sdk.Instance.PutState(key, field, countStr)
}
