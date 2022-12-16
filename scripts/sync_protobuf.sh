#!/bin/bash

# Example: bash sync_protobuf.sh -I=/home/ubuntu/protobuf/include
# where /home/ubuntu/protobuf/include contains google/protobuf/*.proto

set -e -x
script_path=$(readlink -f "$0")
script_dir=$(dirname "$script_path")
project_dir="${script_dir}/.."
basic_folder="protobuf/basic"

cd $project_dir

protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/infra_adapter/service.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/infra_adapter/location/location.proto=sudosdk/protobuf/basic/protobuf/infra_adapter \
    --go_opt=Msudo/protobuf/enums/enums.proto=sudosdk/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/api/api.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/api_usage/api_usage.proto=sudosdk/protobuf/basic/protobuf/api_usage \
    --go_opt=Msudo/protobuf/common.proto=sudosdk/protobuf/basic/protobuf/common \
    --go_opt=module=sudosdk \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/common.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/enums/enums.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/api/api.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/infra_adapter.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/service.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/api_usage/api_usage.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/location/location.proto

protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/googleapis \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/service.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/options.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/options \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/misc.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/misc \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/vtable.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/vtable \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/user.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/user \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/cron.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/cron \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/data_source.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/datasource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/resource.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/resource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/yellowpage.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/yellowpage \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/paginator.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/jwt.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/jwt \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/party.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/party \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/enums/enums.proto=sudosdk/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/common.proto=sudosdk/protobuf/basic/protobuf/common \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=module=sudosdk \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/mpc_virtual_service/*/*.proto

protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    -I=$project_dir/sudo-apis \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/googleapis \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/options.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/options \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/misc.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/misc \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/vtable.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/vtable \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/user.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/user \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/cron.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/cron \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/data_source.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/datasource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/resource.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/resource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/yellowpage.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/yellowpage \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/paginator.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/jwt.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/jwt \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/enums/enums.proto=sudosdk/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/common.proto=sudosdk/protobuf/basic/protobuf/common \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/service.proto=sudosdk/protobuf/virtualservice/platformpb \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/user.proto=sudosdk/protobuf/virtualservice/platformpb/user \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/furnacestatus.proto=sudosdk/protobuf/virtualservice/platformpb/furnacestatus \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/cron.proto=sudosdk/protobuf/virtualservice/platformpb/cron \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/vtable.proto=sudosdk/protobuf/virtualservice/platformpb/vtable \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/common.proto=sudosdk/protobuf/virtualservice/platformpb/common \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/perm.proto=sudosdk/protobuf/virtualservice/platformpb/perm \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/service.proto=sudosdk/protobuf/virtualservice/platformpb \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/yellowpage.proto=sudosdk/protobuf/virtualservice/platformpb/yellowpage \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/paginator.proto=sudosdk/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/pir.proto=sudosdk/protobuf/virtualservice/platformpb/pir \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/token.proto=sudosdk/protobuf/virtualservice/platformpb/token \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/jwt.proto=sudosdk/protobuf/virtualservice/platformpb/jwt \
    --go_opt=Mprotobuf/service/online_service_pir.proto=sudosdk/protobuf/online_service \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=module=sudosdk \
    $project_dir/sudo-apis/protobuf/mpc_virtual_service/*/*.proto

protoc $@ -I=$project_dir/sudo-apis \
    -I=$project_dir/sudo-apis/protobuf/basic \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/infra_adapter/service.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/infra_adapter/location/location.proto=sudosdk/protobuf/basic/protobuf/infra_adapter \
    --go_opt=Msudo/protobuf/enums/enums.proto=sudosdk/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/api/api.proto=sudosdk/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/common.proto=sudosdk/protobuf/basic/protobuf/common \
    --go_opt=Mprotobuf/distributed_compute/distributed_compute.proto=sudosdk/protobuf/distributed_compute \
    --go_opt=Mprotobuf/service/online_service_pir.proto=sudosdk/protobuf/online_service \
    --go_opt=module=sudosdk \
    $project_dir/sudo-apis/protobuf/distributed_compute/distributed_compute.proto \
    $project_dir/sudo-apis/protobuf/service/online_service_pir.proto

# grpcgw api deepcopy
protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    --deepcopy_out=$project_dir \
    --deepcopy_opt=module=sudosdk \
    --deepcopy_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=sudosdk/protobuf/basic/protobuf \
    --deepcopy_opt=Msudo/protobuf/infra_adapter/service.proto=sudosdk/protobuf/basic/protobuf \
    --deepcopy_opt=Msudo/protobuf/infra_adapter/location/location.proto=sudosdk/protobuf/basic/protobuf/infra_adapter \
    --deepcopy_opt=Msudo/protobuf/enums/enums.proto=sudosdk/protobuf/basic/protobuf/enums \
    --deepcopy_opt=Msudo/protobuf/api/api.proto=sudosdk/protobuf/basic/protobuf \
    --deepcopy_opt=Msudo/protobuf/common.proto=sudosdk/protobuf/basic/protobuf/common \
    --deepcopy_opt=Msudo/protobuf/mpc_virtual_service/platform/resource.proto=sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/resource \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/common.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/enums/enums.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/api/api.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/infra_adapter.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/service.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/api/api.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/location/location.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/mpc_virtual_service/platform/resource.proto
