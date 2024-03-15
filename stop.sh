# bin/bash
# 关闭链
set -eu

cd scripts
pwd
source log.sh

# 关闭长安链
info "关闭长安链"
bash chainmaker_stop.sh
info "长安链关闭成功"

# 关闭以太坊
# info "关闭以太坊"
