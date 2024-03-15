#!/bin/bash

contract_name="transaction"

# echo "please input contract name: "
# read contract_name
go build -ldflags="-s -w" -o $contract_name
7z a $contract_name $contract_name