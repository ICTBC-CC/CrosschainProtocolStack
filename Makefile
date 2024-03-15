CHAINMAKER_CRYPTOGEN_PATH = $(shell dirname ${PWD})/chainmaker-cryptogen
CHAINMAKER_GO_PATH = $(shell dirname ${PWD})/chainmaker-go
CURR_PATH = $(shell pwd)

# ! 证书模式
chainmaker: environment-check chainmaker-go-build build_link execution_permission
# ! 证书模式,指定端口
# chainmaker: environment-check chainmaker-go-build_port build_link execution_permission
# !公钥模式
# chainmaker: environment-check chainmaker-go-build_pk build_link execution_permission
# chainmaker: build_link

environment-check:
	@bash ./scripts/env_check.sh
	@bash ./scripts/download_repo.sh

chainmaker-go-build: chainmaker-cryptogen-build
	@echo "build chainmaker-go..."
# @cd $(CHAINMAKER_GO_PATH) && git checkout v2.3.2
	@cd $(CHAINMAKER_GO_PATH) && git checkout v2.3.1
	@cd $(CHAINMAKER_GO_PATH)/tools && ln -fs ../../chainmaker-cryptogen/ .
	@cd $(CHAINMAKER_GO_PATH)/scripts && echo -e "1\nINFO\nYES\n\n" | ./prepare.sh 4 4
	@cd $(CHAINMAKER_GO_PATH)/scripts && ./build_release.sh

# 指定端口
# 										11301		12301			32351				22351
# prepare.sh NODE_CNT=$1 CHAIN_CNT=$2 P2P_PORT=$3 RPC_PORT=$4 VM_GO_RUNTIME_PORT=$5 VM_GO_ENGINE_PORT=$6
chainmaker-go-build_port: chainmaker-cryptogen-build
	@echo "build chainmaker-go..."
	@cd $(CHAINMAKER_GO_PATH) && git checkout v2.3.1
	@cd $(CHAINMAKER_GO_PATH)/tools && ln -fs ../../chainmaker-cryptogen/ .
	@cd $(CHAINMAKER_GO_PATH)/scripts && echo -e "1\nINFO\nYES\n\n" | ./prepare.sh 4 3 31301 32301
	@cd $(CHAINMAKER_GO_PATH)/scripts && ./build_release.sh

# !长安链证书模式启动
chainmaker-go-build_pk: chainmaker-cryptogen-build
	@echo "build chainmaker-go with pk..."
	@cd $(CHAINMAKER_GO_PATH) && git checkout v2.3.1
	@cd $(CHAINMAKER_GO_PATH)/tools && ln -fs ../../chainmaker-cryptogen/ .
	@cd $(CHAINMAKER_GO_PATH)/scripts && echo -e "5\nINFO\n\nYES\n\n" | ./prepare_pk.sh 4 4
	@cd $(CHAINMAKER_GO_PATH)/scripts && ./build_release.sh

chainmaker-cryptogen-build:
	@echo "build chainmaker-cryptogen..."
	@cd $(CHAINMAKER_CRYPTOGEN_PATH) && make
	@cp -r $(CHAINMAKER_CRYPTOGEN_PATH)/bin .

build_link:
	@echo "build link..."
	@if [ ! -d $(CURR_PATH)/config_files/chain1 ]; then mkdir -p $(CURR_PATH)/config_files/chain1; fi
	@cd $(CURR_PATH)/config_files/chain1 && if [ -d crypto-config ]; then rm -rf crypto-config; fi && ln -s ../../../chainmaker-go/build/crypto-config .
	@if [ ! -d $(CURR_PATH)/config_files/chain2 ]; then mkdir -p $(CURR_PATH)/config_files/chain2; fi
	@cd $(CURR_PATH)/config_files/chain2 && if [ -d crypto-config ]; then rm -rf crypto-config; fi && ln -s ../../../chainmaker-go/build/crypto-config .
	@if [ ! -d $(CURR_PATH)/config_files/chain3 ]; then mkdir -p $(CURR_PATH)/config_files/chain3; fi
	@cd $(CURR_PATH)/config_files/chain3 && if [ -d crypto-config ]; then rm -rf crypto-config; fi && ln -s ../../../chainmaker-go/build/crypto-config .
	@if [ ! -d $(CURR_PATH)/config_files/chain4 ]; then mkdir -p $(CURR_PATH)/config_files/chain4; fi
	@cd $(CURR_PATH)/config_files/chain4 && if [ -d crypto-config ]; then rm -rf crypto-config; fi && ln -s ../../../chainmaker-go/build/crypto-config .

execution_permission:
	@echo "execution permission..."
	@if [ -d $(CURR_PATH)/bin ]; then cd $(CURR_PATH)/bin && chmod 777 * ; fi
	@cd $(CURR_PATH)/scripts && chmod 777 *
	@cd $(CURR_PATH)
	@chmod 777 restart.sh
	@chmod 777 stop.sh
	@chmod 777 build_install_contract.sh

fileExist = $(shell if [ -d ${CHAINMAKER_GO_PATH}/build ]; then echo "exist"; else echo "noexist"; fi)
.PHONY:clean
clean:
ifeq ("$(fileExist)", "exist")
	@cd $(CHAINMAKER_GO_PATH)/scripts && ./cluster_quick_stop.sh
	@echo "y" | docker container prune
	@cd $(CHAINMAKER_GO_PATH) && sudo rm -rf $(CHAINMAKER_GO_PATH)/build/release/chainmaker*.org
endif
