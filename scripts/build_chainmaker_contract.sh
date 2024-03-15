# bin/bash
# 编译长安链合约
set -eu

INFO(){
    local content=$*
    local DATE_NOW=`date "+%Y-%m-%d %H:%M:%S"`
    echo -e "\033[32m[INFO ${DATE_NOW}][chainmaker] ${content} \033[0m"
}

INFO "编译端链合约"

CURR_PATH="$(pwd)"
INFO "编译 testapp 合约"
cd $CURR_PATH/chainmaker/contracts/app/testapp && bash ./build.sh

INFO "编译 transfer 合约"
cd $CURR_PATH/chainmaker/contracts/transfer && bash ./build.sh

INFO "编译 verify 合约"
cd $CURR_PATH/chainmaker/contracts/verify && bash ./build.sh

INFO "编译 transaction 合约"
cd $CURR_PATH/chainmaker/contracts/transaction && bash ./build.sh

INFO "编译 resource 合约"
cd $CURR_PATH/chainmaker/contracts/resource && bash ./build.sh

INFO "编译 protocolaggregator 合约"
cd $CURR_PATH/chainmaker/contracts/protocolaggregator && bash ./build.sh


# INFO "编译中继链合约"

INFO "编译 transfer 合约"
cd $CURR_PATH/chainmaker/contracts_relayer/transfer && bash ./build.sh

INFO "编译 verify 合约"
cd $CURR_PATH/chainmaker/contracts_relayer/verify && bash ./build.sh

INFO "编译 transaction 合约"
cd $CURR_PATH/chainmaker/contracts_relayer/transaction && bash ./build.sh

INFO "编译 protocolaggregator 合约"
cd $CURR_PATH/chainmaker/contracts_relayer/protocolaggregator && bash ./build.sh
