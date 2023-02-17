#!/bin/bash

# Example: bash sync_protobuf.sh -I=/home/ubuntu/protobuf/include
# where /home/ubuntu/protobuf/include contains google/protobuf/*.proto

set -e
script_path=$(readlink -f "$0")
script_dir=$(dirname "$script_path")
project_dir="${script_dir}/.."
basic_folder="protobuf/basic"

cd $project_dir

protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/grpc-gateway \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/infra_adapter/location/location.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/infra_adapter \
    --go_opt=Msudo/protobuf/enums/enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/common \
    --go_opt=Msudo/protobuf/api_usage/api_usage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/api_usage \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/paginator.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Msudo/protobuf/service/enums/online_svc_enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service/enums \
    --go_opt=Msudo/protobuf/service/online_service.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=module=github.com/sudo-privacy/sudo-sdk-go \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/common.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/enums/enums.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/infra_adapter.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/infra_adapter/location/location.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/api_usage/api_usage.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/service/enums/online_svc_enums.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/service/*.proto

protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/googleapis \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/grpc-gateway \
    --go_opt=Msudo/protobuf/api_usage/api_usage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/api_usage \
    --go_opt=Msudo/protobuf/service/online_service.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/enums/online_svc_enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service/enums \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/service/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/service/common \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/service/furnace.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/service/furnace \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/service/coordinator.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/service/coordinator \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/options.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/options \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/misc.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/misc \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/vtable.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/vtable \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/user.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/user \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/cron.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/cron \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/data_source.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/datasource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/resource.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/resource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/yellowpage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/yellowpage \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/paginator.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/jwt.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/jwt \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/party.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/party \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/pir \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/token.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/token \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/apl.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/apl \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/online_service.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/task.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/task \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/enums/enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/common \
    --go_opt=Msudo/protobuf/infra_adapter/location/location.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/infra_adapter \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=module=github.com/sudo-privacy/sudo-sdk-go \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/mpc_virtual_service/platform/*.proto \
    $project_dir/sudo-apis/protobuf/basic/sudo/protobuf/mpc_virtual_service/platform/service/*.proto

protoc $@ -I=$project_dir/sudo-apis/protobuf/basic \
    -I=$project_dir/sudo-apis \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/googleapis \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/grpc-gateway \
    --go_opt=Msudo/protobuf/service/online_service.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/enums/online_svc_enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service/enums \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/options.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/options \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/misc.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/misc \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/vtable.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/vtable \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/user.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/user \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/cron.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/cron \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/data_source.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/datasource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/resource.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/resource \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/yellowpage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/yellowpage \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/paginator.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/jwt.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/jwt \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/pir \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/token.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/token \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/apl.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/apl \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/online_service.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service \
    --go_opt=Msudo/protobuf/mpc_virtual_service/platform/task.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/task \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/enums/enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/common \
    --go_opt=Msudo/protobuf/infra_adapter/location/location.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/infra_adapter \
    --go_opt=Msudo/protobuf/api_usage/api_usage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/api_usage \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/user.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/user \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/furnacestatus.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/furnacestatus \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/cron.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/cron \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/vtable.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/vtable \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/common \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/perm.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/perm \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/service/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/service/common \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/service/furnace.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/service/furnace \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/yellowpage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/yellowpage \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/paginator.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/paginator \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/pir \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/token.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/token \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/jwt.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/jwt \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/apiusage.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/apiusage \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/task.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/task \
    --go_opt=Mprotobuf/mpc_virtual_service/platform/feature.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/feature \
    --go_opt=Mprotobuf/service/online_service_pir.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service \
    --go_opt=Mprotobuf/service/enums/online_svc_enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service/enums \
    --go_opt=Mprotobuf/service/online_service_common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service \
    --go_opt=Mprotobuf/service/online_service_f3s.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=module=github.com/sudo-privacy/sudo-sdk-go \
    $project_dir/sudo-apis/protobuf/mpc_virtual_service/platform/*.proto \
    $project_dir/sudo-apis/protobuf/mpc_virtual_service/platform/service/common.proto \
    $project_dir/sudo-apis/protobuf/mpc_virtual_service/platform/service/furnace.proto

protoc $@ -I=$project_dir/sudo-apis \
    -I=$project_dir/sudo-apis/protobuf/basic \
    -I=$project_dir/sudo-apis/protobuf/basic/sudo/protobuf/grpc-gateway \
    --go_out=plugins=grpc:$project_dir \
    --go_opt=Msudo/protobuf/service/online_service.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/online_service_common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service \
    --go_opt=Msudo/protobuf/service/enums/online_svc_enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/service/enums \
    --go_opt=Msudo/protobuf/infra_adapter/infra_adapter.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf \
    --go_opt=Msudo/protobuf/common.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/common \
    --go_opt=Msudo/protobuf/enums/enums.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums \
    --go_opt=Msudo/protobuf/infra_adapter/location/location.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/infra_adapter \
    --go_opt=Mprotobuf/service/online_service_f3s.proto=github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service \
    --go_opt=module=github.com/sudo-privacy/sudo-sdk-go \
    $project_dir/sudo-apis/protobuf/service/online_service_f3s.proto
