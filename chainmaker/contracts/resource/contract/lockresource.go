// coding:utf-8
// 锁定资源
package contract

import (
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 锁定资源
// @return response
func (r *Resource) lockResource() pb.Response {
	// 获取参数
	msg, err := r.getArgs()
	if err != nil {
		return sdk.Error("get args error:" + err.Error())
	}
	app := string(msg.FromApp)

	// 锁定每一个资源
	for _, item := range msg.AtomicSwap.LockResource {
		if err := r.lockOneResource(app, item.ResourceName, item.Count); err != nil {
			return sdk.Error(
				fmt.Sprintf(
					"lock app:%s resource:%s count:%d error:%s",
					app, item.ResourceName, item.Count, err.Error(),
				))
		}
	}

	return sdk.Success(nil)
}

// 锁定一种资源
// @param app app名
// @param name 资源名
// @param count 锁定的数量
// @return response
func (r *Resource) lockOneResource(app, name string, count int) error {
	if count == 0 {
		// 转移数量为0直接返回
		return nil
	}
	// 数值判断
	if count < 0 {
		return fmt.Errorf("count<=0")
	}

	// 获取资源可用数量
	validCount, err := r.getValidCount(app, name)
	if err != nil {
		// 没有就新建资源
		// todo:仅用作测试
		validCount = 1000
		if resp := r.newResource(app, name, validCount); resp.Status != sdk.OK {
			return fmt.Errorf(resp.Message)
		}
		// return sdk.Error(err.Error())
	}

	if count > validCount {
		return fmt.Errorf("need %d, have %d", count, validCount)
	}

	// 减少可用数量
	if err := r.setValidCount(app, name, validCount-count); err != nil {
		return err
	}

	// 获取已经锁定的数量
	lockCount, err := r.getLockCount(app, name)
	if err != nil {
		return err
	}

	// 增加锁定数量
	if err := r.setLockCount(app, name, lockCount+count); err != nil {
		return err
	}

	r.emit_info([]string{
		"lock", app, name, fmt.Sprint(count),
		"validCount", fmt.Sprint(validCount),
		"lockCount", fmt.Sprint(lockCount),
	})

	return nil
}
