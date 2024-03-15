// coding:utf-8
// 统一的链下客户端接口
package common

type Client interface {
	// 订阅消息
	// @param contractName 合约名
	// @param eventName 事件名
	// @param msg_type 消息类型,send or receive or info
	// @param resultChan 保存事件的通道
	SubscribeEvent(
		contractName,
		eventName string,
		msg_type MSG_TYPE,
		resultChan chan<- *SubscribeMsg,
	)

	// 调用合约方法
	// @param contractName 合约名称
	// @param method 方法名
	// @param args 参数列表
	// @return any result
	// @return error err
	InvokeContract(
		contractName string,
		method string,
		args ...any,
	) (any, error)

	// 客户端内部消息转化为订阅MSG格式
	// @param msg 客户端内部消息
	// @return subscribemsg 订阅消息
	// @return error 报错
	// todo:这个应该可以不要
	SelfMsg2SubscribeMsg(msg any) (*SubscribeMsg, error)

	// 订阅的MSG消息格式转为发送到链上receive函数的参数格式,配合invoke函数
	// @param msg 被转化的订阅MSG消息
	// @return any 客户端内部消息
	// @return error 错误消息
	SubscribeMsg2ReceiveArgs(msg *SubscribeMsg) (any, error)

	// 从订阅的MSG中解析出链上协议栈消息
	// @param msg 需要被解析的数据格式,通过指针改变对象
	// @return error 错误消息
	ParseEvent(msg *SubscribeMsg) error
	// 从订阅的MSG中解析出数据的源链和目的链信息
	// @param data 订阅监听到的消息
	// @return string 解析出的源链
	// @return string 解析出的目的链
	// @return []byte 解析出的数据包
	// @return error 错误消息
	// ParseEvent(data []byte) (string, string, []byte, error)

	// 获取链名
	// @return string 链名
	GetChainname() string

	// 关闭客户端
	Close()
}
