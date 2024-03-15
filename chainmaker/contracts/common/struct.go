// coding:utf-8
// 结构体
package common

import "CPS/common"

// 协议数据结构体
type ProtocolMsg struct {
	// ************* 资源 *************
	// 转移的资源
	AtomicSwap *common.AtomicSwap

	// ************* 应用层 *************
	FromApp []byte // 消息发出的app
	AppFunc []byte // 调用app的功能
	AppData []byte // 跨链数据

	// ************* 事务层 *************
	TransactionID       []byte // 事务ID
	TransactionState    string // 事务状态
	TransactionData     []byte // 事务层数据
	TransactionProtocol string // 事务协议

	// ************* 验证层 *************
	// ! 中继链验源链,目的链验中继链,源链不验
	VerifyType string // 验证类型
	VerifyData []byte // 验证层数据

	// ************* 转发层 *************
	// ! 注意,到了目的链这个字段会改为对应的链名, 类似于MAC地址
	FromChain     []byte // 消息的源链
	ToChain       []byte // 消息的目的链
	TransferState []byte // 表示处于send或response阶段
}
