# 中继链的合约说明
* `transaction`事务层合约
* `verify`验证层合约
* `transfer`转发层合约
* `common`一些公共组件,包括各层结构体以及各种类型定义
* `protocolaggregator`协议聚合器
* `test`文件夹是开发代码的时候,进行测试的合约
* **中继链没有应用层**, 只负责事务的调度,所以也没有资源合约

## `resource`资源合约
* 跨链涉及到的资源必须定义在`resource`合约当中
* 应用资源的初始状态必须保存在`resource`合约,事务执行成功则直接`改变资源状态`,并由`协议栈`通知`应用`事务执行成功

## 整体执行流程
#### 长安链send
* 1. 应用app定义好各层使用的协议，以及传递的数据，保存到`ProtocolMsg`结构体中
* 2. 将`ProtocolMsg`进行`marshal`以字节数组的形式保存在参数中的`data`字段并调用`protocolAggregator`合约
* 3. 由`protocalaggregator`进行`unmarshal`,并执行协议栈
* 4. 链下网关监听转发层事件,获取完整`protocalMsg`,随后转发

#### 长安链receive
* 1. 链下网关监听到转发层事件,解码消息并发送到对应的链
* 2. 参数的`data`是`marshal`后的字节数组,调用`protocalaggregator`
* 3. 到达协议聚合器后,进行`unmarshal`,然后执行协议栈并返回`response`