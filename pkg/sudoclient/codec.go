package sudoclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"

	"github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/options"
)

func UnmarshalJSONToProto(data []byte, dest proto.Message) error {
	mapstruct := make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&mapstruct); err != nil {
		return err
	}
	return decodeMessageFromMapStruct(mapstruct, dest.ProtoReflect())
}

func decodeMessageFromMapStruct(iVal interface{}, msg protoreflect.Message) error {
	if asVAlue, ok := proto.GetExtension(msg.Descriptor().Options(), options.E_AsValue).(bool); ok && asVAlue {
		if msg.Descriptor().Fields().Len() == 1 {
			iVal = map[string]interface{}{
				string(msg.Descriptor().Fields().Get(0).Name()): iVal,
			}
		}
	}
	val, ok := iVal.(map[string]interface{})
	if !ok {
		return fmt.Errorf("value %+v is not a map struct in field", iVal)
	}
	for idx := 0; idx < msg.Descriptor().Fields().Len(); idx++ {
		fd := msg.Descriptor().Fields().Get(idx)
		if fd.Kind() == protoreflect.GroupKind {
			return fmt.Errorf("group kind field not supported: %+v", fd)
		}

		if squash, ok := proto.GetExtension(fd.Options(), options.E_JsonSquash).(bool); ok && squash {
			newMsg := msg.NewField(fd)
			if err := decodeMessageFromMapStruct(val, newMsg.Message()); err != nil {
				return err
			}
			msg.Set(fd, newMsg)
			continue
		}
		if inner, ok := val[string(fd.Name())]; !ok || inner == nil {
			continue
		}

		err := decodeOneField(fd, msg, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeOneField(fd protoreflect.FieldDescriptor, msg protoreflect.Message, val map[string]interface{}) error {
	switch {
	case fd.IsList():
		protoList := msg.NewField(fd).List()
		innerValues, ok := val[string(fd.Name())].([]interface{})
		if !ok {
			return fmt.Errorf("field %s is expcted to be a list", fd.Name())
		}
		for _, innerValue := range innerValues {
			if fd.Kind() != protoreflect.MessageKind {
				protoValue, err := castPrimitiveValueToProto(fd, innerValue)
				if err != nil {
					return fmt.Errorf("failed to cast primitive value to proto: %+v", err)
				}
				protoList.Append(protoValue)
			} else {
				newMsg := protoList.NewElement()
				if err := decodeMessageFromMapStruct(innerValue, newMsg.Message()); err != nil {
					return err
				}
				protoList.Append(newMsg)
			}
		}
		msg.Set(fd, protoreflect.ValueOfList(protoList))
	case fd.IsMap():
		err := decodeOneMapField(fd, msg, val)
		if err != nil {
			return err
		}
	default:
		if fd.Kind() != protoreflect.MessageKind {
			protoValue, err := castPrimitiveValueToProto(fd, val[string(fd.Name())])
			if err != nil {
				return fmt.Errorf("failed to cast primitive value to proto: %+v", err)
			}
			msg.Set(fd, protoValue)
		} else {
			newMsg := msg.NewField(fd)
			innerValue, ok := val[string(fd.Name())].(map[string]interface{})
			if !ok {
				return fmt.Errorf("field %s is expcted to be a map", fd.Name())
			}
			if err := decodeMessageFromMapStruct(innerValue, newMsg.Message()); err != nil {
				return fmt.Errorf("failed to decode: %+v", err)
			}
			msg.Set(fd, newMsg)
		}
	}
	return nil
}

func decodeOneMapField(fd protoreflect.FieldDescriptor, msg protoreflect.Message, val map[string]interface{}) error {
	protoMap := msg.NewField(fd).Map()
	innerValue, ok := val[string(fd.Name())].(map[string]interface{})
	if !ok {
		return fmt.Errorf("field %s is expcted to be a map of map", fd.Name())
	}
	for k, v := range innerValue {
		keyValue, err := castPrimitiveValueToProto(fd.MapKey(), k)
		if err != nil {
			return fmt.Errorf("failed to cast primitive value to proto: %+v", err)
		}
		if fd.MapValue().Kind() != protoreflect.MessageKind {
			protoValue, err := castPrimitiveValueToProto(fd.MapValue(), v)
			if err != nil {
				return fmt.Errorf("failed to cast primitive value to proto: %+v", err)
			}
			protoMap.Set(keyValue.MapKey(), protoValue)
		} else {
			newMsg := protoMap.NewValue()
			switch newMsg.Message().Interface().(type) {
			case *structpb.Value:
				stripeJNumberV, err := stripJSONNumber(v)
				if err != nil {
					return err
				}
				structpbValue, err := structpb.NewValue(stripeJNumberV)
				if err != nil {
					return err
				}
				protoMap.Set(keyValue.MapKey(), protoreflect.ValueOf(structpbValue.ProtoReflect()))
			case *structpb.ListValue:
				sliceV, ok := v.([]interface{})
				if !ok {
					return fmt.Errorf("value %+v is not slice in field %s", v, fd.Name())
				}
				stripeJNumberSliceV, err := stripJSONNumber(sliceV)
				if err != nil {
					return err
				}
				structpbValue, err := structpb.NewList(stripeJNumberSliceV.([]interface{}))
				if err != nil {
					return err
				}
				protoMap.Set(keyValue.MapKey(), protoreflect.ValueOf(structpbValue.ProtoReflect()))
			default:
				if err := decodeMessageFromMapStruct(v, newMsg.Message()); err != nil {
					return err
				}
				protoMap.Set(keyValue.MapKey(), newMsg)
			}
		}
	}
	msg.Set(fd, protoreflect.ValueOfMap(protoMap))
	return nil
}

// nolint
func castPrimitiveValueToProto(fd protoreflect.FieldDescriptor, val interface{}) (protoreflect.Value, error) {
	var protoValue protoreflect.Value
	switch fd.Kind() {
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.FloatKind, protoreflect.DoubleKind:
		num, ok := val.(json.Number)
		if !ok {
			if numStr, ok := val.(string); ok {
				if numStr == "" {
					numStr = "0"
				}
				num = json.Number(numStr)
			} else {
				return protoreflect.Value{}, fmt.Errorf("value %+enumVal is expected to be a number", val)
			}
		}

		switch fd.Kind() {
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			floatNum, err := num.Float64()
			if err != nil {
				return protoreflect.Value{}, fmt.Errorf("failed to convert json number to float value: %+enumVal", err)
			}
			switch fd.Kind() {
			case protoreflect.FloatKind:
				protoValue = protoreflect.ValueOfFloat32(float32(floatNum))
			case protoreflect.DoubleKind:
				protoValue = protoreflect.ValueOfFloat64(floatNum)
			}
		case protoreflect.Int32Kind, protoreflect.Int64Kind:
			intNum, err := num.Int64()
			if err != nil {
				return protoreflect.Value{}, fmt.Errorf("failed to convert json number to int value: %+enumVal", err)
			}
			switch fd.Kind() {
			case protoreflect.Int32Kind:
				protoValue = protoreflect.ValueOfInt32(int32(intNum))
			case protoreflect.Int64Kind:
				protoValue = protoreflect.ValueOfInt64(intNum)
			}
		case protoreflect.Uint32Kind, protoreflect.Uint64Kind:
			intNum, err := strconv.ParseUint(num.String(), 10, 64)
			if err != nil {
				return protoreflect.Value{}, fmt.Errorf("failed to convert json number to int value: %+enumVal", err)
			}
			switch fd.Kind() {
			case protoreflect.Uint32Kind:
				protoValue = protoreflect.ValueOfUint32(uint32(intNum))
			case protoreflect.Uint64Kind:
				protoValue = protoreflect.ValueOfUint64(uint64(intNum))
			}
		}
	case protoreflect.EnumKind:
		enum, err := protoregistry.GlobalTypes.FindEnumByName(fd.Enum().FullName())
		switch {
		case errors.Is(err, protoregistry.NotFound):
			return protoreflect.Value{}, fmt.Errorf("enum %q is not registered", fd.Enum().FullName())
		case err != nil:
			return protoreflect.Value{}, fmt.Errorf("failed to look up enum: %w", err)
		}
		// Look for enum by name
		strVal, ok := val.(string)
		if !ok {
			numVal, ok := val.(json.Number)
			if !ok {
				return protoreflect.Value{}, fmt.Errorf("cannot cast value %+v to enum value", val)
			}
			strVal = numVal.String()
		}
		enumVal := enum.Descriptor().Values().ByName(protoreflect.Name(strVal))
		if enumVal == nil {
			// 再次使用大写重试
			enumVal = enum.Descriptor().Values().ByName(protoreflect.Name(strings.ToUpper(strVal)))
		}
		if enumVal == nil {
			i, err := strconv.Atoi(strVal)
			if err != nil {
				return protoreflect.Value{}, fmt.Errorf("%q is not a valid value", strVal)
			}
			// Look for enum by number
			enumVal = enum.Descriptor().Values().ByNumber(protoreflect.EnumNumber(i))
			if enumVal == nil {
				return protoreflect.Value{}, fmt.Errorf("%q is not a valid value", strVal)
			}
		}
		protoValue = protoreflect.ValueOfEnum(enumVal.Number())
	case protoreflect.BytesKind:
		strVal, ok := val.(string)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("cannot cast value %+v to string", val)
		}
		bytesVal, err := runtime.Bytes(strVal)
		if err != nil {
			return protoreflect.Value{}, err
		}
		protoValue = protoreflect.ValueOfBytes(bytesVal)
	case protoreflect.BoolKind, protoreflect.StringKind:
		protoValue = protoreflect.ValueOf(val)
	default:
		return protoreflect.Value{}, fmt.Errorf("unsupported value %+enumVal", val)
	}

	return protoValue, nil
}

func stripJSONNumber(input interface{}) (interface{}, error) {
	switch input.(type) {
	case map[string]interface{}:
		inputBytes, err := json.Marshal(input)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal input %v", input)
		}
		resp := make(map[string]interface{})
		err = json.Unmarshal(inputBytes, &resp)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal inputBytes %s", string(inputBytes))
		}
		return resp, nil
	case []interface{}:
		inputBytes, err := json.Marshal(input)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal input %v", input)
		}
		resp := make([]interface{}, 0)
		err = json.Unmarshal(inputBytes, &resp)
		if err != nil {
			return nil, fmt.Errorf("fail to marshal inputBytes %s", string(inputBytes))
		}
		return resp, nil
	}
	if jsonNumber, ok := input.(json.Number); ok {
		if int64Resp, err := jsonNumber.Int64(); err == nil {
			return int64Resp, nil
		}
		if uint64Resp, err := strconv.ParseInt(jsonNumber.String(), 10, 64); err == nil {
			return uint64Resp, nil
		}
		if float64Resp, err := jsonNumber.Float64(); err == nil {
			return float64Resp, nil
		}
		return jsonNumber.String(), nil
	}
	return input, nil
}
