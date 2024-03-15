# bin/bash
# 长安链启动脚本
set -euo pipefail

CHAINMAKER_GO_PATH=`dirname $(dirname ${PWD})`/chainmaker-go

echo $CHAINMAKER_GO_PATH

cd ${CHAINMAKER_GO_PATH}/scripts && ./cluster_quick_start.sh normal
