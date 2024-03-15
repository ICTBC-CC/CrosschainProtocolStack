// coding:utf-8
// 执行事务层的receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用事务层receive
// @param transaction_data 事务层数据
// @param transaction_id 事务层id
// @param fromapp 源app
// @param toapp 目的app
// @param msg 整个消息
// @return error err
func (t *ProtocolAggregator) receiveTransaction(msg *common.ProtocolMsg) error {
	// 1. 获取参数
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal atomic swap error:" + err.Error())
	}
	kvs := map[string][]byte{cps_common.KEY_MSG: msgByte}
	// 2. 调用事务层进行资源处理, 返回事务是否执行完成
	resp := sdk.Instance.CallContract(
		common.LAYER_TRANSACTION,
		common.FUNC_RECEIVE,
		kvs,
	)
	// 转发层调用失败就返回错误
	if resp.Status != sdk.OK {
		// ! 好好考虑一下事务层receive失败是否需要回滚
		return fmt.Errorf("call transaction receive error:%s", resp.Message)
	}
	// 没有payload就不进行下一步
	if len(resp.Payload) == 0 {
		// 为0表示匹配没有完成
		t.emit_info([]string{"没有payload"})
		return nil
	}

	// 3. 不为0表示需要进行下一步,进行解码和转发
	mpSwap := make(map[string]*cps_common.AtomicSwap)
	// unmarshal
	if err := json.Unmarshal(resp.Payload, &mpSwap); err != nil {
		t.emit_warning([]string{"unmarshal mtp transaction payload error", err.Error()})
		return fmt.Errorf("unmarshal mtp transaction payload error:" + err.Error())
	}

	// 4. 事务层执行成功以后,通过转发层发送消息出去
	// ! 执行BTP协议的发送
	if msg.TransactionProtocol == cps_common.TRANSACTION_PROTOCOL_TYPE_BTP {
		if msg.TransactionState == common.TRANSACTION_PREPARE {
			// ! 是BTP协议,同时是prepare状态才发送给每一条链
			t.emit_info([]string{"执行中继链BTP的prepare事务状态操作"})
			// 4.1 如果是BTP就发给每一条相关目的链
			for chainname, swap := range mpSwap {
				// 发送每一条链的转发事件
				msg.AtomicSwap = swap
				msg.ToChain = []byte(chainname)

				if err := t.sendTransfer(msg); err != nil {
					return fmt.Errorf("call send transfer error:%s", err)
				}
				t.emit_info([]string{chainname, "call send transfer success"})
			}
		} else if msg.TransactionState == common.TRANSACTION_COMMIT {
			// ! 是BTP协议,同时是commit状态,则转发出去,不做处理
			t.emit_info([]string{"执行中继链BTP的commit事务状态操作"})

			if err := t.sendTransfer(msg); err != nil {
				return fmt.Errorf("call send transfer error:%s", err)
			}
		} else {
			return fmt.Errorf("中继链执行BTP协议,进入了一个不存在的交易状态")
		}
		return nil
	}

	if msg.TransactionProtocol == cps_common.TRANSACTION_PROTOCOL_TYPE_MTP {
		// 4.2 是MTP且COMMIT就调用response发送出去
		for chainname, swap := range mpSwap {
			// 发送每一条链的提交事件
			msg.FromChain = []byte(chainname)
			msg.AtomicSwap = swap

			if err := t.responseTransfer(msg); err != nil {
				return fmt.Errorf(chainname + "call response transfer error:" + err.Error())
			}
			t.emit_info([]string{chainname, "call response transfer success"})
		}
		return nil
	}

	t.emit_warning([]string{"既不是BTP也不是MTP协议,进入了后面", msg.TransactionProtocol})
	return nil
}
