// coding:utf-8
// 日志
package utils

import (
	"fmt"
	"path"
	"runtime"
)

// 获取函数调用者的名字
func getCallerInfo() (info string) {
	pc, file, lineNo, ok := runtime.Caller(2)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}

// 在控制台打印出不同颜色的信息
// 0 - 黑色
// 1 - 红色
// 2 - 绿色
// 3 - 黄色
// 4 - 蓝色
// 5 - 紫红色
// 6 - 青蓝色;
func ColorPrint(color int, message string) {
	fmt.Printf("\033[0;%dm%s\033[0m\n", color+30, message)
}

// 打印输出一条错误消息
// @param funcName 哪一个函数名报错, 自己指定
// @param err 错误信息
func InfoError(err error) string {
	info := fmt.Sprintln("funcName->" + getCallerInfo() + " error->" + err.Error())
	ColorPrint(1, info)
	return info
}

// 打印输出消息
// @param info 任意消息
// @return 合并后的消息
func Info(info ...any) string {
	result := fmt.Sprint("info: ", info)
	ColorPrint(2, result)
	return result
}

// 打印输出消息, 不加前缀
// @param info 任意消息
// @return 合并后的消息
func Print(info ...any) string {
	result := fmt.Sprint(info...)
	ColorPrint(0, result)
	return result
}

// 打印警告消息
// @param info 任意消息
// @return 合并后的消息
func InfoWarning(info ...any) string {
	result := fmt.Sprint("warning: ", info)
	ColorPrint(3, result)
	return result
}

// 打印提示消息
// @param info 任意消息
// @return 合并后的消息
func InfoTips(info ...any) string {
	result := fmt.Sprint("tips: ", info)
	ColorPrint(4, result)
	return result
}
