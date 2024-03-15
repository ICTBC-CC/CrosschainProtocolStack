// coding:utf-8
// 常量
package common

import "CPS/common"

// *********************** 资源合约 ***********************

// 资源合约用到的参数KEY
const (
	KEY_RESOURCE_USE_FROM string = common.KEY_RESOURCE_USE_FROM // 事务处理源链使用的资源
	KEY_RESOURCE_USE_TO   string = common.KEY_RESOURCE_USE_TO   // 事务处理目的链使用的资源
)

// 资源合约提供的方法
const (
	// 资源的四个名字
	FUNC_RESOURCE_NEW      string = "resourcenew" // 新建资源或者增加资源
	FUNC_RESOURCE_LOCK     string = "resourcelock"
	FUNC_RESOURCE_ROLLBACK string = "resourcerollback"
	FUNC_RESOURCE_COMMIT   string = "resource_commit" // 提交消息里的资源
	FUNC_RESOURCE_DESTROY  string = "resourcedestroy" // 销毁一定的资源
	// 获取锁定的资源个数
	FUNC_RESOURCE_GET_LOCKCOUNT string = "get_lock_count"
	// 获取可用的资源个数
	FUNC_RESOURCE_GET_VALID_COUNT string = "get_valid_count"
)

// *********************** 应用层 ***********************

// *********************** 事务层 ***********************

// 定义事务状态
const (
	TRANSACTION_PRE_PREPARE string = "transaction_pre_prepare" // 刚发起
	TRANSACTION_PREPARE     string = "transaction_prepare"     // 发起链/目的链预锁成功
	TRANSACTION_COMMIT      string = "transaction_commit"      // 发起链/目的链已经提交
	TRANSACTION_FAIL        string = "transaction_fail"        // 事务执行失败
	TRANSACTION_COMMIT_WAIT string = "transaction_commit_wait" // 中继链等待事务提交
)

// 事务层用到的数据库KEY和VALUE相关值
const (
	KEY_FIELD_TRANSACTION_ID  string = "field_transaction_id" // 保存交易id的数据库域
	VALUE_TRANSACTION_EXIST   string = "y"                    // 事务id对应的value,表示事务存在
	KEY_TRANSACTION_NEWEST_ID string = "key_newest_id"        // 保存事务层最新id的key
)

// ! 注意这里定义了事务类型,因为后续可能会改变类型,比如改为math.bigInt
// 事务ID类型
type TypeTransactionID uint64

const (
	// ! 注意这里的空事务ID类型的数字应该和事务ID类型同步更改
	// ! 注意一定要重载+=
	// 定义空事务ID类型
	NIL_TRANSACTION_ID = TypeTransactionID(0)
)

// *********************** 验证层 ***********************

// 验证层验证类型
var VERIFY_TYPE_HAPPY string = common.VERIFY_TYPE_HAPPY   // 乐观验证
var VERIFY_TYPE_NOTARY string = common.VERIFY_TYPE_NOTARY // 公证人
var VERIFY_TYPE_SPV string = common.VERIFY_TYPE_SPV       // spv验证

// 验证层用到的数据库KEY和VALUE相关值
const (
	KEY_FIELD_NOTARY    string = "key_field_notary" // 公证人数据库域
	VALUE_VERIFY_NOTARY string = "y"                // 公证人存储的value
)

// 验证层用到的参数KEY
const (
	KEY_VERIFY_TYPE string = common.KEY_VERIFY_TYPE // 验证层验证类型
	KEY_VERIFY_DATA string = common.KEY_VERIFY_DATA // 验证层数据
)

// 验证层特有功能名
const (
	FUNC_VERIFY_ADD_NOTORY string = "add_notary"
)

// *********************** 转发层 ***********************
// ACK状态字节码
var TRANSFER_STATE_SEND []byte = []byte("state_send")         // 表示当前消息为第一阶段send消息
var TRANSFER_STATE_RESPONSE []byte = []byte("state_response") // 表示当前消息为第二阶段响应消息

// *********************** 常量定义 ***********************

// 定义参数的key
const (
	KEY_FROM_CHAIN     string = "fromchain"     // 转发层源链
	KEY_TO_CHAIN       string = "tochain"       // 转发层目的链
	KEY_TRANSFER_STATE string = "transferstate" // 转发层状态
	KEY_TRANSFER_DATA  string = "transferdata"  // 完整数据包

	KEY_FROM_APP string = "fromapp" // 应用层源app
	KEY_APP_FUNC string = "appfunc" // 应用层app调用的函数
	KEY_APP_DATA string = "appdata" // 应用层数据

	KEY_TRANSACTION_DATA  string = "transactiondata"  // 事务层数据key
	KEY_TRANSACTION_ID    string = "transactionid"    // 事务层id key
	KEY_TRANSACTION_STATE string = "transactionstate" // 事务层状态key

	KEY_ATOMIC_SWAP string = common.KEY_ATOMIC_SWAP // 交换的资源
)

// 定义每一层合约的名字
const (
	// ******* 聚合器相关合约名 *******
	LAYER_PROTOCOL_AGGREGATOR string = "protocolaggregator" // 协议聚合器

	// ******* 事务层相关合约名 *******
	LAYER_TRANSACTION string = "transaction" // 事务层
	LAYER_RESOURCE    string = "resource"    // 资源合约

	// ******* 验证层相关合约名 *******
	LAYER_VERIFY string = "verify" // 验证层

	// ******* 转发层相关合约名 *******
	LAYER_TRANSFER string = "transfer" // 转发层
)

// 定义基本的方法名
const (
	FUNC_SEND     string = common.FUNC_SEND     // 本地发送出去消息
	FUNC_RECEIVE  string = common.FUNC_RECEIVE  // 本地接收到一个消息
	FUNC_REVERT   string = common.FUNC_REVERT   // 回滚
	FUNC_RESPONSE string = common.FUNC_RESPONSE // 本地对接收的消息进行响应
)

// 定义事件名
const (
	EVENT_TRANSFER_SEND    string = common.EVENT_TRANSFER_SEND    // 转发层发送事件
	EVENT_TRANSFER_RECEIVE string = common.EVENT_TRANSFER_RECEIVE // 转发层接收事件
	EVENT_TRANSFER_REVERT  string = common.EVENT_TRANSFER_REVERT  // 转发层回滚事件
	EVENT_INFO             string = common.EVENT_INFO             // 通知消息
	EVENT_WARNING          string = common.EVENT_WARNING          // 运行中的警告消息
	EVENT_ERROR            string = common.EVENT_ERROR            // 运行中的错误消息,用于调试
)

// 资源合约常量定义
const (
	// 可用资源的域名
	VALID_FIELD_NAME string = "field_valid"
	// 锁定资源的域名
	LOCK_FIELD_NAME string = "field_lock"
	// 现有待匹配资源域名
	MARCH_FIELD_NAME string = "field_march"
)
