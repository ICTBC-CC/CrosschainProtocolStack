#! /bin/bash
set -euo pipefail

info(){
    echo -e "\033[32m[$0 INFO] \033[0m$1"
}


error(){
    echo -e "\033[31m[$0 ERROR] \033[0m$1"
}