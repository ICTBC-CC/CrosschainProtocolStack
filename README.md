# 基于中继链的跨链交互协议栈
* `chainmaker`长安链模块全部逻辑及代码
* `eth`以太坊模块全部逻辑及代码
* `fabric``fabric`模块全部逻辑及代码
* `utils`全部工具函数
* `scripts`全部`shell`脚本
* `example`协议栈使用的示例
* `img`是文档当中使用到的图片
* `config_files`是全部的链配置文件和用户配置文件
* `common`是链上链下公用的常量定义
* `timmer`是事务层链下计时器
* `relayer`是中继网关
* **仅在`ubuntu`以及`centeros`系统下测试过**,理论上可以兼容所有linux发行版
* **本示例使用`chain1`和`chain2`作为端链,`chain3`作为中继链**

## 快速启动
* 按照[长安链官网安装好各种必须的软件](https://docs.chainmaker.org.cn/v2.3.2/html/quickstart/%E9%80%9A%E8%BF%87%E5%91%BD%E4%BB%A4%E8%A1%8C%E4%BD%93%E9%AA%8C%E9%93%BE.html),并注册[长安链github账号](https://git.chainmaker.org.cn/)
* 执行命令,下载编译`长安链`,`以太坊`, 并启动链, 会在本地启动区块链,包括`两条长安链`,`两个poa以太坊`,`两个fabric`
```bash
make && bash ./restart.sh
```
* 编译并部署合约
```bash
bash ./build_install_contract.sh
```
* 启动网关
```bash
go run main.go -t relayer
```
* 监听通知事件(可以不启动,仅用作调试)
```bash
go run main.go -t info
```
* 执行示例
```bash
go run main.go -t testapp
```
* 运行mtp示例
```bash
go run main.go -t mtp
```
* 运行btp示例,进行信息交换
```bash
# btp不进行资源交换,只进行信息传递
go run main.go -t btp
```
* 运行示例
```bash
go run example/sendmsg2peer.go
```
* 运行资源合约测试
```bash
go run main.go -t testresource
```
* 运行资源合约测试数量
```bash
go run main.go -t testresourcecount
```

#### 测试各个代码的结果
* 因为项目涉及到链上链下的交互,所以用`testing`比较麻烦,需要打桩,还需要考虑到文件路径,因此,本项目的功能测试没有使用`testing`,而是直接在`test`文件夹下实现的.
* 执行测试
```bash
go run main.go -t 1
```

## 测试BTP协议
* 启动事件监听
```bash
go run main.go -t info
```
* 启动中继
```bash
go run main.go -t relayer
```
* 发起一条BTP跨链交易
```bash
go run main.go -t btp
```
* 会看到输出结果

## 测试MTP协议
* 启动事件监听
```bash
go run main.go -t info
```
* 启动中继
```bash
go run main.go -t relayer
```
* 发起一条BTP跨链交易
```bash
go run main.go -t mtp
```
* 会看到输出结果

## 实现进度
&#9745;建立项目文件夹路径

&#9745;抽象接口

&#9744;长安链协议栈框架

&#9745;长安链一键启动脚本

&#9745;长安链协议聚合器

&#9745;长安链客户端

&#9744;长安链验证层

&#9744;长安链转发层

&#9745;链下网关逻辑确定

&#9744;长安链状态定义

&#9745;长安链状态锁定方案

&#9745;长安链中继链的事务层实现

&#9745;长安链事务层实现

&#9745;长安链链下网关

&#9745;长安链协议栈实现

&#9745;长安链跨长安链

&#9744;以太坊客户端

&#9744;以太坊状态定义

&#9744;以太坊协议聚合器

&#9744;以太坊状态锁定方案

&#9744;以太坊验转发层

&#9744;以太坊验验证层

&#9744;以太坊事务层实现

&#9744;以太坊链下网关

&#9744;以太坊协议栈实现

&#9744;以太坊跨以太坊

&#9744;异构跨链:以太坊跨长安链

&#9744;多中继异构跨链

&#9744;长安链spv

&#9744;以太坊spv

&#9745;项目一键启动脚本

&#9745;环境检查脚本

&#9745;一键部署脚本

&#9744;项目各个模块介绍图片

&#9744;代码注释率检测脚本

&#9744;代码注释率15%以上
