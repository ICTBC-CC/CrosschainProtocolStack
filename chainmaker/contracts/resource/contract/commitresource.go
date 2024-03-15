// coding:utf-8
// 提交资源
package contract

import (
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 提交want资源,即增加
// @return response
func (r *Resource) funcCommitResource() pb.Response {
	// 获取参数
	msg, err := r.getArgs()
	if err != nil {
		return sdk.Error("get args error:" + err.Error())
	}
	app := string(msg.FromApp)

	// 锁定每一个资源
	for _, item := range msg.AtomicSwap.WantResource {
		if err := r.commitOneResource(app, item.ResourceName, item.Count); err != nil {
			return sdk.Error(
				fmt.Sprintf(
					"lock app:%s resource:%s count:%d error:%s",
					app, item.ResourceName, item.Count, err.Error(),
				))
		}
	}

	return sdk.Success(nil)
}

// 提交一种资源,即增加
// @param app app名
// @param name 资源名
// @param count 锁定的数量
// @return response
func (r *Resource) commitOneResource(app, name string, count int) error {
	if count == 0 {
		// 转移数量为0直接返回
		return nil
	}
	// 数值判断
	if count < 0 {
		return fmt.Errorf("count<=0")
	}

	// 获取资源可用数量
	validCount, _ := r.getValidCount(app, name)

	// 增加可用数量
	if err := r.setValidCount(app, name, validCount+count); err != nil {
		return fmt.Errorf("commit one resource set valid count error:" + err.Error())
	}

	r.emit_info([]string{
		"commit", app, name, fmt.Sprint(count),
		"validCount", fmt.Sprint(validCount),
	})

	return nil
}
