package leetcode

import (
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

var globalProtoMap map[string]*desc.FileDescriptor
var IsCached = true
var lk sync.RWMutex

func init() {
	globalProtoMap = make(map[string]*desc.FileDescriptor)
}

func getProto(path string) *desc.FileDescriptor {
	lk.Lock()
	defer lk.Unlock()

	fd, ok := globalProtoMap[path]
	if ok {
		return fd
	}
	p := protoparse.Parser{}
	fds, err := p.ParseFiles(path)
	if err != nil {
		return nil
	}
	fd = fds[0]
	globalProtoMap[path] = fd
	return fd
}

// JsonToPb 传入proto文件的path, proto中对应的message.name，js的原始数据
// 返回生成的proto.Marshal的[]byte
// example:
// path := "$PROTOPATH/helloworld.proto"
// messageName "helloworld.HelloRequest"
// JsonToPb(path,"helloworld.HelloRequest", []byte(`{"name":"yzh"}`))
func JsonToPb(protoPath, messageName string, jsonStr []byte) ([]byte, error) {

	fd := getProto(protoPath)

	msg := fd.FindMessage(messageName)

	dymsg := dynamic.NewMessage(msg)
	err := dymsg.UnmarshalJSON(jsonStr)
	if err != nil {
		return nil, nil
	}

	any, err := ptypes.MarshalAny(dymsg)
	if err != nil {
		return nil, nil
	}
	return any.Value, nil
}

// PbToJson 传入proto的byte数据，返回它对应的json数据
// example:
// path := "$PROTOPATH/helloworld.proto"
// messageName "helloworld.HelloRequest"
// jsonByte, err := PbToJson(path, messageName, pbByte)
func PbToJson(protoPath, messageName string, protoData []byte) ([]byte, error) {
	fd := getProto(protoPath)
	msg := fd.FindMessage(messageName)
	dymsg := dynamic.NewMessage(msg)

	err := proto.Unmarshal(protoData, dymsg)

	jsonByte, err := dymsg.MarshalJSON()
	return jsonByte, err
}
