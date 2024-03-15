// coding:utf-8
// 测试应用层合约
package main

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"
	"log"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type TESTAPP struct{}

func (t *TESTAPP) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *TESTAPP) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *TESTAPP) InvokeContract(method string) pb.Response {
	// according method segment to select contract functions
	switch method {
	case cps_common.FUNC_SEND:
		// 执行协议聚合器的调用,返回应用层跨链数据
		return t.funcSend()
	case cps_common.FUNC_RECEIVE:
		// 执行协议聚合器的调用,返回应用层跨链数据
		return t.funcReceive()
	default:
		return sdk.Error("invalid function " + method)
	}
}

// *************************** receive ***************************

// 执行合约应用层receive
func (t *TESTAPP) funcReceive() pb.Response {
	t.emit_info([]string{"进入test app func receive"})
	msg, err := t.getArgs()
	if err != nil {
		return sdk.Error(err.Error())
	}

	// 调用不同的函数
	appFunc := string(msg.AppFunc)
	var payload []byte
	switch appFunc {
	case cps_common.FUNC_BTP:
		{
			payload, err = t.receiveBtp(msg)
		}
	case cps_common.FUNC_MTP:
		{
			payload, err = t.receiveMtp(msg)
		}
	default:
		{
			payload = nil
			err = fmt.Errorf("invalid app func:" + appFunc)
		}
	}
	// 返回结果
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(payload)
}

// 执行mtp测试
// @param msg msg
// @return []byte 跨链应用层数据
// @return error err
func (t *TESTAPP) receiveMtp(msg *common.ProtocolMsg) ([]byte, error) {
	t.emit_info([]string{"进入testapp func mtp"})

	return []byte("this is a response mtp message from test app"), nil
}

// 执行btp测试
// @param msg msg
// @return []byte 跨链应用层数据
// @return error err
func (t *TESTAPP) receiveBtp(msg *common.ProtocolMsg) ([]byte, error) {
	t.emit_info([]string{"进入testapp func btp"})

	return []byte("this is a response btp message from test app"), nil
}

// *************************** send ***************************

// 执行应用层合约send
func (t *TESTAPP) funcSend() pb.Response {
	t.emit_info([]string{"进入test app func send"})
	msg, err := t.getArgs()
	if err != nil {
		return sdk.Error(err.Error())
	}

	// 调用不同的函数
	appFunc := string(msg.AppFunc)
	var payload []byte
	switch appFunc {
	case cps_common.FUNC_BTP:
		{
			payload, err = t.sendBtp(msg)
		}
	case cps_common.FUNC_MTP:
		{
			payload, err = t.sendMtp(msg)
		}
	default:
		{
			payload = nil
			err = fmt.Errorf("invalid app func:" + appFunc)
		}
	}
	// 返回结果
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(payload)
}

// 执行mtp测试
// @param msg msg
// @return []byte 跨链应用层数据
// @return error err
func (t *TESTAPP) sendMtp(msg *common.ProtocolMsg) ([]byte, error) {
	t.emit_info([]string{"进入testapp func mtp"})

	return []byte("this is a send mtp message from test app"), nil
}

// 执行btp测试
// @param msg msg
// @return []byte 跨链应用层数据
// @return error err
func (t *TESTAPP) sendBtp(msg *common.ProtocolMsg) ([]byte, error) {
	t.emit_info([]string{"进入testapp func btp"})

	return []byte("this is a send btp message from test app"), nil
}

// *************************** 其他 ***************************

// 抛出警告事件
// @param data 警告信息
func (t *TESTAPP) emit_warning(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_WARNING, data)
}

// 抛出通知事件
// @param data 通知事件
func (t *TESTAPP) emit_info(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_INFO, data)
}

// 获取参数
// @return msg msg
// @return error err
func (t *TESTAPP) getArgs() (*common.ProtocolMsg, error) {
	// 获取msg
	msgByte := sdk.Instance.GetArgs()[cps_common.KEY_MSG]
	if len(msgByte) == 0 {
		t.emit_warning([]string{"get args error", "invalid length of msg byte"})
		return nil, fmt.Errorf("invalid length of msg byte")
	}

	var msg common.ProtocolMsg
	if err := json.Unmarshal(msgByte, &msg); err != nil {
		t.emit_warning([]string{"get args unmarshal msg error", err.Error()})
		return nil, fmt.Errorf("unmarshal msg error:" + err.Error())
	}

	return &msg, nil
}

// main
func main() {
	err := sandbox.Start(new(TESTAPP))
	if err != nil {
		log.Fatal(err)
	}
}
