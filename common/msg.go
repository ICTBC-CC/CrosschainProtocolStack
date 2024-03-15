// coding:utf-8
// 前后端跨链消息接口定义
package common

// **************** 链下网关消息定义 ****************
// 链下消息类型
type MSG_TYPE uint8

// 链下消息类型定义
const (
	MSG_ERROR       MSG_TYPE = 0x01 // 错误消息
	MSG_CROSS_CHAIN MSG_TYPE = 0x02 // 跨链消息
	MSG_SEND        MSG_TYPE = 0x03 // 跨链send消息
	MSG_RECEIVE     MSG_TYPE = 0x04 // 跨链receive消息
	MSG_INFO        MSG_TYPE = 0x05 // 通知消息
	MSG_WARNING     MSG_TYPE = 0x06 // 警告消息
)

// 链下网关订阅的统一消息接口
type SubscribeMsg struct {
	// ************** 链下区块链客户端加入的信息 **************
	Type         MSG_TYPE // 消息的类型
	ChainName    string   // 消息的来源链
	UserName     string   // 监听到这个消息的用户
	ContractName string   // 监听的合约名
	EventName    string   // 监听到的事件名
	ChainData    []byte   // 监听到的对应链上消息的原始字节码

	// ************** 解析出的链上消息 **************
	// ! 注意, 下面的字段都来自于ChainData的解析结构, 解析ChainData后以下字段才会有值

	// ************* 资源 *************
	// 转移的资源
	AtomicSwap *AtomicSwap

	// ************* 应用层 *************
	AppFrom []byte // 消息发出的app
	AppFunc []byte // 调用的应用功能
	AppData []byte // 应用层跨链数据

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
