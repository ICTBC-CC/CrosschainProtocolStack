// coding:utf-8
// 获取数据
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"
	"strconv"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 获取可用资源数量
// @param app 应用名
// @param name 资源名
// @return int 资源数量
// @return error err
func (r *Resource) getValidCount(app, name string) (int, error) {
	count, err := r.getFieldCount(common.VALID_FIELD_NAME, app, name)
	return count, err
}

// 获取锁定资源数量
// @param app 应用名
// @param name 资源名
// @return int 资源数量
// @return error err
func (r *Resource) getLockCount(app, name string) (int, error) {
	// 首先判断资源key对不对
	if _, err := r.makeKey(app, name); err != nil {
		return 0, fmt.Errorf("invalid app and name error:" + err.Error())
	}

	// 数据库中查找,看看能不能找到key
	count, err := r.getFieldCount(common.LOCK_FIELD_NAME, app, name)
	if err != nil {
		// 表示这个资源还没有锁定过,就直接返回
		sdk.Instance.EmitEvent(common.EVENT_WARNING, []string{
			"get lock field name warning:" + err.Error(),
		})
		return 0, nil
	}
	return count, nil
}

// 获取域中资源数量h
// @param field 域名
// @param app 应用名
// @param name 资源名
// @return int 资源数量
// @return error err
func (r *Resource) getFieldCount(field, app, name string) (int, error) {
	// 资源key值
	key, err := r.makeKey(app, name)
	if err != nil {
		return 0, err
	}

	// 获取可域中资源数量
	countStr, b, err := sdk.Instance.GetStateWithExists(key, field)
	if err != nil || !b {
		return 0, fmt.Errorf("invalid resource %s %s", app, name)
	}

	// 变为整数
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}
