#! /bin/bash
set -euo pipefail

workspace=${PWD}/`basename $(dirname $0)`
repospace=`dirname $(dirname $workspace)`

source ${workspace}/log.sh

cd ${repospace}

if [ -d "chainmaker-go" ];then
    info "chainmaker-go exist."
else
    info "from git.chainmaker.org.cn clone chainmaker-go..."
    git clone -b v2.3.2 https://git.chainmaker.org.cn/chainmaker/chainmaker-go.git
fi

if [ -d "chainmaker-cryptogen" ];then
    info "chainmaker-cryptogen exist."
else
    info "from git.chainmaker.org.cn clone chainmaker-cryptogen..."
    git clone -b v2.3.0 https://git.chainmaker.org.cn/chainmaker/chainmaker-cryptogen.git
fi
