package sudoclient

// Copyright 2022 Sudo Technology Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/infra_adapter"
	protoapl "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/apl"
	protoonlinesvc "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/resource"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/task"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/user"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/vtable"
	"google.golang.org/protobuf/proto"
)

func TestUnmarshalUint64(t *testing.T) {
	input := map[string]interface{}{
		"triplets_bit_info": 1,
		"service_id":        "442450667450138891",
		"project_id":        "",
	}
	inputByte, err := json.Marshal(input)
	require.Nil(t, err)
	param := protoonlinesvc.ServiceParam{}
	err = UnmarshalJSONToProto(inputByte, &param)
	require.Nil(t, err)
	require.Equal(t, uint64(442450667450138891), param.ServiceId)
}

func TestUnmarshalInterface(t *testing.T) {
	input := map[string]interface{}{
		"action": "APPROVE",
		"additional_parameters": map[string]interface{}{
			"parameter1": "value1",
			"parameter2": "value2",
		},
	}

	inputByte, err := json.Marshal(input)
	require.Nil(t, err)

	param := protoapl.AplTakeActionParam{}
	err = UnmarshalJSONToProto(inputByte, &param)
	require.Nil(t, err)
	require.Equal(t, 2, len(param.AdditionalParameters))
}

func TestUnmarshalLowerEnum(t *testing.T) {
	input := map[string]interface{}{
		"action":         "terminate",
		"service_id_str": "1234",
	}

	inputByte, err := json.Marshal(input)
	require.Nil(t, err)

	param := protoonlinesvc.ServiceTakeActionParam{}
	err = UnmarshalJSONToProto(inputByte, &param)
	require.Nil(t, err)
}

func TestUnmarshalIntMapKey(t *testing.T) {
	// sliceValue情况
	input := map[string]interface{}{
		"history_progress": map[uint64][]interface{}{
			1: {
				map[string]interface{}{
					"ts":       111,
					"inner_id": 1,
					"party_id": "party1",
					"msg":      "test_msg",
				},
			},
		},
	}

	inputBytes, err := json.Marshal(input)
	require.Nil(t, err)

	stage := task.Stage{}

	err = UnmarshalJSONToProto(inputBytes, &stage)
	require.Nil(t, err)
	assert.Equal(t, 1, len(stage.HistoryProgress))

	// 一般情况
	resourceInput := map[string]interface{}{
		"status": map[uint64]interface{}{
			1: map[string]interface{}{
				"op_host": "test_host",
			},
			2: map[string]interface{}{
				"op_host": "test_host2",
			},
		},
	}

	resourceInputBytes, err := json.Marshal(resourceInput)
	require.Nil(t, err)
	resource := resource.PartyResourceStatus{}
	err = UnmarshalJSONToProto(resourceInputBytes, &resource)
	require.Nil(t, err)
	assert.Equal(t, 2, len(resource.Status))
}

func TestUnmarshalHideKey(t *testing.T) {
	input := map[string]interface{}{
		"ids":             []string{"1", "2", "3"},
		"feature_columns": []string{"c1", "c2"},
		"features": [][]string{
			{
				"c1_value1", "c2_value1",
			},
			{
				"c1_value2", "c2_value2",
			},
			{
				"c1_value3", "c2_value3",
			},
		},
	}
	inputByte, err := json.Marshal(input)
	require.Nil(t, err)

	param := protoonlinesvc.PredictRequest{}
	err = UnmarshalJSONToProto(inputByte, &param)
	require.Nil(t, err)
	require.Equal(t, 3, len(param.Features))
	require.Equal(t, 2, len(param.Features[0].Values))
	require.Nil(t, err)
}

func TestUnmarshalJSON(t *testing.T) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"columns": []interface{}{
			map[string]interface{}{"field": "field 1"},
			map[string]interface{}{"type": "type 2"},
		},
		"name": "vtable name",
		"owner_user": map[string]interface{}{
			"username": "user name",
		},
	})
	vt := &vtable.Vtable{}
	assert.Nil(t, UnmarshalJSONToProto(bytes, vt))
	expVt := &vtable.Vtable{
		Base: &vtable.VtableBase{
			Columns: []*vtable.Column{
				{Field: "field 1"},
				{Type: "type 2"},
			},
			Name: "vtable name",
		},
		OwnerUser: &user.User{
			Base: &user.UserBase{
				Username: "user name",
			},
		},
	}
	assert.True(t, proto.Equal(expVt, vt))
}

func TestUnmarshalJSONWithOneOfField(t *testing.T) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"db_table": map[string]interface{}{
			"db_connection": map[string]interface{}{
				"odbc_connection": map[string]interface{}{
					"connection": "connection",
				},
			},
			"query":     "query",
			"vtable_id": 1,
		},
		"is_tmp_data": true,
		"metadata": map[string]interface{}{
			"file_type": "Csv",
		},
	})
	dataLocation := &infra_adapter.DataLocation{}
	assert.Nil(t, UnmarshalJSONToProto(bytes, dataLocation))
	expDataLocation := &infra_adapter.DataLocation{
		Data: &infra_adapter.DataLocation_DbTable{
			DbTable: &infra_adapter.DataLocation_DB{
				VtableId: 1,
				Path:     &infra_adapter.DataLocation_DB_Query{Query: "query"},
				DbConnection: &infra_adapter.DataLocation_DBConnection{
					Connection: &infra_adapter.DataLocation_DBConnection_OdbcConnection{
						OdbcConnection: &infra_adapter.DataLocation_ODBCConnection{
							Connection: "connection",
						},
					},
				},
			},
		},
		IsTmpData: true,
		Metadata: &infra_adapter.MetaData{
			FileType: 1,
		},
	}
	assert.True(t, proto.Equal(expDataLocation, dataLocation))
}

func TestUnmarshalJSONWithEnumField(t *testing.T) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"db_table": map[string]interface{}{
			"db_connection": map[string]interface{}{
				"odbc_connection": map[string]interface{}{
					"connection": "connection",
				},
			},
			"query":     "query",
			"vtable_id": 1,
		},
		"is_tmp_data": true,
		"metadata": map[string]interface{}{
			"file_type": 1,
		},
	})
	dataLocation := &infra_adapter.DataLocation{}
	expDataLocation := &infra_adapter.DataLocation{
		Data: &infra_adapter.DataLocation_DbTable{
			DbTable: &infra_adapter.DataLocation_DB{
				VtableId: 1,
				Path:     &infra_adapter.DataLocation_DB_Query{Query: "query"},
				DbConnection: &infra_adapter.DataLocation_DBConnection{
					Connection: &infra_adapter.DataLocation_DBConnection_OdbcConnection{
						OdbcConnection: &infra_adapter.DataLocation_ODBCConnection{
							Connection: "connection",
						},
					},
				},
			},
		},
		IsTmpData: true,
		Metadata: &infra_adapter.MetaData{
			FileType: 1,
		},
	}
	assert.Nil(t, UnmarshalJSONToProto(bytes, dataLocation))
	assert.True(t, proto.Equal(expDataLocation, dataLocation))
}
