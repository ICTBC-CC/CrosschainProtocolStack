// coding:utf-8
// 常量定义
package common

// 链类型名字定义
const (
	// 长安链类型
	ChainmakerChainType string = "chainmaker"
	// 以太坊链类型
	EthChainType string = "eth"
	// fabric链类型
	FabricChainType string = "fabric"
)

// 合约名
const (
	CONTRACT_NAME_PROTOCOL_AGGREGATOR string = "protocolaggregator"
	CONTRACT_NAME_TESTAPP             string = "testapp"
	CONTRACT_NAME_RESOURCE            string = "resource"
)

// *********************** 验证层 ***********************

// 验证层验证类型
const (
	VERIFY_TYPE_HAPPY  string = "happy"  // 乐观验证
	VERIFY_TYPE_NOTARY string = "notary" // 公证人
	VERIFY_TYPE_SPV    string = "spv"    // spv验证
)

// *********************** 事务层 ***********************
// 事务层协议类型
const (
	TRANSACTION_PROTOCOL_TYPE_MTP string = "mtp"          // mtp协议
	TRANSACTION_PROTOCOL_TYPE_BTP string = "btp"          // btp
	TRANSACTION_COMPLETE          string = "complete"     // 事务执行完成
	TRANSACTION_NOT_COMPLETE      string = "not_complete" // 事务没有执行完成
)

// ***********************  ***********************

// 参数key
const (
	KEY_VERIFY_TYPE          string = "verify_type"         // 验证层验证类型
	KEY_VERIFY_DATA          string = "verify_data"         // 验证层数据
	KEY_RESOURCE_USE_FROM    string = "resource_use_from"   // 事务处理源链转移的资源
	KEY_RESOURCE_USE_TO      string = "resource_use_to"     // 事务处理目的链转移的资源
	KEY_ATOMIC_SWAP          string = "atomic_swap"         // 进行原子交换的资源
	KEY_TRANSACTION_PROTOCOL string = "transactin_protocol" // 事务层协议
	KEY_MSG                  string = "message"             // 整个完成的消息数据包
)

// 功能名
const (
	FUNC_MTP      string = "func_mpt" // 执行MTP事务
	FUNC_BTP      string = "func_bpt" // 执行MTP事务
	FUNC_SEND     string = "send"     // 本地发送出去消息
	FUNC_RECEIVE  string = "receive"  // 本地接收到一个消息
	FUNC_REVERT   string = "revert"   // 回滚
	FUNC_RESPONSE string = "response" // 本地对接收的消息进行响应
)

// 定义事件名
const (
	EVENT_TRANSFER_SEND    string = "event_transfer_send"    // 转发层发送事件
	EVENT_TRANSFER_RECEIVE string = "event_transfer_receive" // 转发层接收事件
	EVENT_TRANSFER_REVERT  string = "event_transfer_revert"  // 转发层回滚事件
	EVENT_INFO             string = "event_info"             // 通知消息
	EVENT_WARNING          string = "event_warning"          // 运行中的警告消息
	EVENT_ERROR            string = "event_error"            // 运行中的错误消息,用于调试
)
