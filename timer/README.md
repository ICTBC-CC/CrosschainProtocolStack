# 链下网关
* 链下网关的逻辑处理
* **注意,只在两条链中做中继**
* 链下网关只监听`send`事件，并将监听到的事件做解码
* 如果从端链监听到`send`事件, 就直接发给中继链
* 如果从中继链监听到`send`事件, 就判断目的链是不是自己的端链
  * 如果是自己的端链，就发送到端链，否则忽略

## 需要解决的问题
1. 链下网关会不会作恶，如果网关不按照既定规则进行交易的转发，会对系统造成什么影响，如何从协议栈层面进行规避
2. 是否可以多网关转发，多网关转发后如何去重
3. 或许不是作恶，而是单纯的参数错误，某一个网关将端链标识为了中继链，如此会发生什么问题，如何在协议栈层面解决，同`1`