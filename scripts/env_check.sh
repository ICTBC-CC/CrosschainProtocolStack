#! /bin/bash

set -euo pipefail

workspace=${PWD}/`basename $(dirname $0)`
source ${workspace}/log.sh

command_check(){
    command -v $1 > /dev/null 2>&1 && echo "check $1 installed." || { echo "command $1 not exist";exit 1; }
}

# 版本检查
version_greater_equal() { 
    let result=1
    ls_n1=($(echo $1 | tr "." "\n"))
    ls_n2=($(echo $2 | tr "." "\n"))
    len1=${#ls_n1[@]}
    len2=${#ls_n2[@]}
    # echo $len1
    # echo $len2
    min_len=`echo "$len1 $len2" | tr " " "\n" | sort -V | head -n 1`
    # echo "min len $min_len"

    for ((i=0;i<$min_len;i++))
    do
        # echo "${ls_n1[$i]} ${ls_n2[$i]}"
        if [ ${ls_n1[$i]} -gt ${ls_n2[$i]} ]; then
            # 大于就表示正确
            # echo "大于"
            result=1
            break
        fi
        if [ ${ls_n1[$i]} -lt ${ls_n2[$i]} ]; then
            # 小于表示错误
            # echo "小于"
            result=0
            break
        fi
        if [ ${ls_n1[$i]} -eq ${ls_n2[$i]} ];then
            # 相同就继续
            # echo "继续"
            continue
        fi
    done
    echo $result
}

# 断言版本
assert_version(){
    result=$(version_greater_equal $1 $2)
    # echo "result $result"
    if test $result -eq 1
    then
        # echo "$1 is greater than $2"
        info "version check pass"
    else
        error "$version小于$1"
        exit 1
    fi
}

go_version_check(){
    command_check go

    # go 版本必须是1.18, 1.19和1.20都不行
    info "check go version. need $1"
    info "go : current version $version"
    go version | grep "go$1" > /dev/null 2>&1 && info "check go version pass." || { error "go version error. need go$1";exit 1; }
    info "go version check pass"
}

# 检查docker版本
docker_version_check(){
    command_check docker

    info "check docker version, should be greater than $1"
    version=`docker -v | tr "," "\n" | head -n 1 | tr " " "\n" | tail -n 1`
    info "docker : current version $version"
    assert_version $version $1
    info "docker version check pass"
}

# 检查docker-compose版本
docker_compose_version_check(){
    command_check docker-compose

    info "check docker-compose version, should be greater than $1"
    version=`docker-compose -v | tr "v" "\n" | tail -n 1`
    info "docker-compose : current version $version"
    assert_version $version $1
    info "docker-compose version check pass"
}

# 检查gcc版本
gcc_version_check(){
    command_check gcc

    info "check gcc version, should be greater than $1"
    version=`gcc --version | head -n 1 | tr " " "\n" |tail -n 1`
    info "gcc : current version $version"
    assert_version $version $1
    info "gcc version check pass"
}

docker_version_check 20.10
# ! go版本必须1.18, 不然无法编译chainmaker
go_version_check 1.18
command_check python3
docker_compose_version_check 2.10
gcc_version_check 7.3
command_check 7z
