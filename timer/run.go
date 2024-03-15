// coding:utf-8
// 执行relayer
package relayer

import (
	"CPS/utils"
	"sync"
)

// 执行relayer,会一直阻塞
// @param ch 全部错误信息通过这个通道返回
func (r *Relayer) Run(ch chan error) {
	//! 自己订阅处理自己的事件才好办,不然转发不好做
	var wg sync.WaitGroup
	wg.Add(1)
	// 1. 启动每条链的事件监听
	for chainname, client := range r.mpClient {
		// 订阅并这条链
		utils.InfoTips("启动", chainname, "中继程序")
		go r.startRelayer(chainname, client, ch)
	}

	wg.Wait()
}
