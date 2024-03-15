// coding:utf-8
// 资源和余额
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 资源合约
type Resource struct {
	// 有四个状态数据库
	// 1. 每个app可用的资源名以及数量, valid->app_name->count
	// 2. 每个app锁定的资源名以及数量, lock->app_name->count
	// ! 不应该是以app为单位进行资源管理,应该是以个人为单位进行资源管理
	// ! 不然资源永远都处在一个app里面,没有任何意义
	// 有两个状态数据库
	// 1. 每个用户可用的资源名以及数量, 可用资源域->用户名_资源名(作为key)->资源数量(作为value), valid->userName_resourceName->count
	// 2. 每个用户锁定的资源名以及数量, 锁定资源域->用户名_资源名(作为key)->资源数量(作为value), lock->userName_resourceName->count
	// todo:应该以用户为单位分配资源,以app为单位锁定资源
	// todo:资源流转过程:user新建资源->user授权app使用资源(锁定资源到app)
	// todo:			->user使用app进行跨链资源转移->app转移资源(无差别)
	// todo:			->结束
}

func (r *Resource) InitContract() pb.Response {
	return sdk.Success(nil)
}

func (r *Resource) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

func (r *Resource) InvokeContract(method string) pb.Response {
	switch method {
	case common.FUNC_RESOURCE_NEW:
		// 新建资源
		r.emit_info([]string{"call resource new"})
		return r.newResource("app", "name", 1)
	case common.FUNC_RESOURCE_LOCK:
		// 锁定/预锁资源
		r.emit_info([]string{"call resource lock"})
		return r.lockResource()
	case common.FUNC_RESOURCE_COMMIT:
		// 提交事务资源
		r.emit_info([]string{"call resource commit"})
		return r.funcCommitResource()
	case common.FUNC_RESOURCE_ROLLBACK:
		// 资源回滚,回退
		r.emit_info([]string{"call resource rollback"})
		return sdk.Success(nil)
	case common.FUNC_RESOURCE_GET_LOCKCOUNT:
		{
			// 获取锁定的资源数量
			app := string(sdk.Instance.GetArgs()[common.KEY_FROM_APP])
			name := string(sdk.Instance.GetArgs()[common.KEY_RESOURCE_USE_FROM])
			count, err := r.getLockCount(app, name)
			if err != nil {
				return sdk.Error(err.Error())
			}
			r.emit_info([]string{app, name, "锁定余额", fmt.Sprint(count)})
			return sdk.Success(nil)
		}
	case common.FUNC_RESOURCE_GET_VALID_COUNT:
		{
			// 获取可用的资源数量
			app := string(sdk.Instance.GetArgs()[common.KEY_FROM_APP])
			name := string(sdk.Instance.GetArgs()[common.KEY_RESOURCE_USE_FROM])
			count, _ := r.getValidCount(app, name)
			r.emit_info([]string{app, name, "可用余额", fmt.Sprint(count)})
			return sdk.Success(nil)
		}
	}
	return sdk.Success(nil)
}
