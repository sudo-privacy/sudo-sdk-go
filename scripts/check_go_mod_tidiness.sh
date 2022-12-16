#!/bin/bash

set -e

go mod tidy -compat=1.17
if [ -n "$(git status --porcelain go.mod go.sum)" ]; then
    git --no-pager diff -p go.sum go.mod
    echo -e "Please run \`go mod tidy -compat=1.17\` in this module"
    exit 1
fi
