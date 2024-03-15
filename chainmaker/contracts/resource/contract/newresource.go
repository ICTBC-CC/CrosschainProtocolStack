// coding:utf-8
// 新建资源
package contract

import (
	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 新建资源
// @param app 新建app的资源
// @param name 新建app的name资源
// @param count 新建app的name资源count个
// @return response
func (r *Resource) newResource(app, name string, count int) pb.Response {
	// 数值判断
	if count < 0 {
		return sdk.Error("count<0")
	}
	// 获取已有资源
	validCount, err := r.getValidCount(app, name)
	if err != nil {
		// 表示资源不存在
		validCount = 0
	}
	// 不论是否存在, 都增加资源
	if err := r.setValidCount(app, name, validCount+count); err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(nil)
}
