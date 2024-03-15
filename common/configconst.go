// coding:utf-8
// 配置文件常量
package common

import "fmt"

// 每条链的链id
const (
	CHAIN_ID_CHAIN1  string = "chain1"
	CHAIN_ID_CHAIN2  string = "chain2"
	CHAIN_ID_CHAIN3  string = "chain3"
	CHAIN_ID_RELAYER string = "chain_relayer"
)

// 所有链1上的管理员用户信息
var Chain1Admins = []string{Chain1Org1Admin1, Chain1Org2Admin1, Chain1Org3Admin1, Chain1Org4Admin1}

// 所有链2上的管理员用户信息
var Chain2Admins = []string{Chain2Org1Admin1, Chain2Org2Admin1, Chain2Org3Admin1, Chain2Org4Admin1}

// 所有链2上的管理员用户信息
var Chain3Admins = []string{Chain3Org1Admin1, Chain3Org2Admin1, Chain3Org3Admin1, Chain3Org4Admin1}

// 所有链2上的管理员用户信息
var ChainRelayerAdmins = []string{Chain4Org1Admin1, Chain4Org2Admin1, Chain4Org3Admin1, Chain4Org4Admin1}

// 每条链对应的基本用户
var MpConfig = map[string]string{
	CHAIN_ID_CHAIN1:  "./config_files/sdkconfigs/chain1_sdkconfig1.yml",
	CHAIN_ID_CHAIN2:  "./config_files/sdkconfigs/chain2_sdkconfig1.yml",
	CHAIN_ID_CHAIN3:  "./config_files/sdkconfigs/chain3_sdkconfig1.yml",
	CHAIN_ID_RELAYER: "./config_files/sdkconfigs/chain4_sdkconfig1.yml",
}

// User用户结构定义了用户基础信息
type User struct {
	TlsKeyPath, TlsCrtPath   string
	SignKeyPath, SignCrtPath string
}

//名称信息
const (
	//组织管理员信息
	Chain1Org1Admin1 = "chain1org1admin1" //链1组织1的管理员名称
	Chain1Org2Admin1 = "chain1org2admin1" //链1组织2的管理员名称
	Chain1Org3Admin1 = "chain1org3admin1" //链1组织3的管理员名称
	Chain1Org4Admin1 = "chain1org4admin1" //链1组织4的管理员名称
	Chain2Org1Admin1 = "chain2org1admin1" //链2组织1的管理员名称
	Chain2Org2Admin1 = "chain2org2admin1" //链2组织2的管理员名称
	Chain2Org3Admin1 = "chain2org3admin1" //链2组织3的管理员名称
	Chain2Org4Admin1 = "chain2org4admin1" //链2组织4的管理员名称
	Chain3Org1Admin1 = "chain3org1admin1" //链2组织1的管理员名称
	Chain3Org2Admin1 = "chain3org2admin1" //链2组织2的管理员名称
	Chain3Org3Admin1 = "chain3org3admin1" //链2组织3的管理员名称
	Chain3Org4Admin1 = "chain3org4admin1" //链2组织4的管理员名称
	Chain4Org1Admin1 = "chain4org1admin1" //链2组织1的管理员名称
	Chain4Org2Admin1 = "chain4org2admin1" //链2组织2的管理员名称
	Chain4Org3Admin1 = "chain4org3admin1" //链2组织3的管理员名称
	Chain4Org4Admin1 = "chain4org4admin1" //链2组织4的管理员名称
	//客户端信息
	Chain1Org1Client1 = "chain1org1client1" //链1组织1的客户端名称
	Chain1Org2Client1 = "chain1org2client1" //链1组织2的客户端名称
	Chain1Org3Client1 = "chain1org3client1" //链1组织3的客户端名称
	Chain1Org4Client1 = "chain1org4client1" //链1组织4的客户端名称
	Chain2Org1Client1 = "chain2org1client1" //链2组织1的客户端名称
	Chain2Org2Client1 = "chain2org2client1" //链2组织2的客户端名称
	Chain2Org3Client1 = "chain2org3client1" //链2组织3的客户端名称
	Chain2Org4Client1 = "chain2org4client1" //链2组织4的客户端名称
	Chain3Org1Client1 = "chain3org1client1" //链2组织1的客户端名称
	Chain3Org2Client1 = "chain3org2client1" //链2组织2的客户端名称
	Chain3Org3Client1 = "chain3org3client1" //链2组织3的客户端名称
	Chain3Org4Client1 = "chain3org4client1" //链2组织4的客户端名称
	Chain4Org1Client1 = "chain4org1client1" //链2组织1的客户端名称
	Chain4Org2Client1 = "chain4org2client1" //链2组织2的客户端名称
	Chain4Org3Client1 = "chain4org3client1" //链2组织3的客户端名称
	Chain4Org4Client1 = "chain4org4client1" //链2组织4的客户端名称
)

// 获取固定路径下管理员证书信息
func GetAdmins() map[string]*User {
	var admins = map[string]*User{
		Chain1Org1Admin1: getConfigPath(1, 1),
		Chain1Org2Admin1: getConfigPath(1, 2),
		Chain1Org3Admin1: getConfigPath(1, 3),
		Chain1Org4Admin1: getConfigPath(1, 4),

		Chain2Org1Admin1: getConfigPath(2, 1),
		Chain2Org2Admin1: getConfigPath(2, 2),
		Chain2Org3Admin1: getConfigPath(2, 3),
		Chain2Org4Admin1: getConfigPath(2, 4),

		Chain3Org1Admin1: getConfigPath(3, 1),
		Chain3Org2Admin1: getConfigPath(3, 2),
		Chain3Org3Admin1: getConfigPath(3, 3),
		Chain3Org4Admin1: getConfigPath(3, 4),

		Chain4Org1Admin1: getConfigPath(4, 1),
		Chain4Org2Admin1: getConfigPath(4, 2),
		Chain4Org3Admin1: getConfigPath(4, 3),
		Chain4Org4Admin1: getConfigPath(4, 4),
	}
	return admins
}

// 获取链的配置文件的路径
// @param chainid 链id
// @param orgid 组织id
// @return *user 用户信息
func getConfigPath(chainid, orgid int) *User {
	CONFIG_PATH := "config_files/"
	return &User{
		fmt.Sprintf(CONFIG_PATH+"chain%d/crypto-config/wx-org%d.chainmaker.org/user/admin1/admin1.tls.key", chainid, orgid),
		fmt.Sprintf(CONFIG_PATH+"chain%d/crypto-config/wx-org%d.chainmaker.org/user/admin1/admin1.tls.crt", chainid, orgid),
		fmt.Sprintf(CONFIG_PATH+"chain%d/crypto-config/wx-org%d.chainmaker.org/user/admin1/admin1.sign.key", chainid, orgid),
		fmt.Sprintf(CONFIG_PATH+"chain%d/crypto-config/wx-org%d.chainmaker.org/user/admin1/admin1.sign.crt", chainid, orgid),
	}
}
