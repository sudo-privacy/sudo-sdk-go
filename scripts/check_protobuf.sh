#!/bin/bash

set -e -x

script_dir="$(dirname "${BASH_SOURCE[0]}")"
project_dir="${script_dir}/.."
cd $project_dir

./scripts/sync_protobuf.sh
if [ -n "$(git status --porcelain ./protobuf)" ]; then
    echo -e "protobuf need update; please use \`./scripts/sync_protobuf.sh\` to update it."
    git diff
    exit 1
fi
