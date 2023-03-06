#!/bin/bash

set -e -x
script_path=$(readlink -f "$0")
script_dir=$(dirname "$script_path")
project_dir="${script_dir}/.."

go test ./... -v -coverpkg=./... -coverprofile=profile.cov
cat profile.cov | grep -v "pb.go" | grep -v "pb.gw.go" | grep -v "deepcopy.go" | grep -v "/generated/" |
    grep -v "/experimental/" >>coverage.out && rm profile.cov

go tool cover -func coverage.out | tail -n 1
go tool cover -html=coverage.out -o coverage.html
rm -f profile.cov coverage.out

cd -
