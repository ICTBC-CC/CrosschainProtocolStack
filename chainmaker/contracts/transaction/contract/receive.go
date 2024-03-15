// coding:utf-8
// 事务层接收消息
package contract

import (
	"CPS/chainmaker/contracts/common"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 事务层接收到消息
// @param fromapp 源链app
// @param toapp 目的链app
// @param state 事务状态
// @param resource_from 源链操作的资源
// @param resource_to 目的链操作的资源
// @return response
func (t *TRANSACTION) receive() pb.Response {
	msg, err := t.getArgs()
	if err != nil {
		return sdk.Error("get args error:" + err.Error())
	}
	state := msg.TransactionState

	// 判断状态
	if state == common.TRANSACTION_PREPARE {
		// ! 如果是prepare状态,就表示自己是目的链,执行目的链的资源提交操作
		// ! 目的链资源提交操作是BTP协议下锁定want资源
		// 执行目的链资源提交操作
		t.emit_info([]string{"receive prepare操作,执行目的链资源提交操作"})
		// 预锁want资源
		// ! 由于prelockresource函数里面是对lock资源进行预锁,所以这里将want赋值给lock,进行锁定
		msg.AtomicSwap.LockResource = msg.AtomicSwap.WantResource
		if err := t.preLockResource(msg); err != nil {
			t.emit_warning([]string{
				"transaction receive prepare, transfer resource error:",
				err.Error(),
			})
			return sdk.Error("transaction exec btp lock want resource error:" + err.Error())
		}
		return sdk.Success(nil)
	}

	// !如果是commit状态,就表示自己是源链,执行源链资源提交操作
	if state == common.TRANSACTION_COMMIT {
		// 执行源链提交操作
		t.emit_info([]string{"receive commit操作,执行源链资源提交操作"})
		// 预锁并转移资源
		if err := t.commitResource(msg); err != nil {
			t.emit_warning([]string{
				"transaction receive commit, transfer resource error:",
				err.Error(),
			})
			// return sdk.Error("transaction prelock and transfer resource error:" + err.Error())
		}
		return sdk.Success(nil)
	}

	// 其他情况表示有问题
	t.emit_warning([]string{
		"transaction receive invalid state",
		string(state),
		// fmt.Sprintf("%+v", swap),
	})
	// return sdk.Error("invalid transaction state:" + string(state))

	return sdk.Success(nil)
}
