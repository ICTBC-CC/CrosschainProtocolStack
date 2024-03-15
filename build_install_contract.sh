# bin/bash
# 编译并部署合约
set -eu

source ./scripts/log.sh

# 编译部署长安链合约
info "编译部署长安链合约"
bash ./scripts/build_chainmaker_contract.sh
go mod tidy
go run chainmaker/deploy.go
info "长安链合约编译部署完成"