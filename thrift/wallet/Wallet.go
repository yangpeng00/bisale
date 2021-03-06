// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package wallet

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type Wallet interface {
  // Parameters:
  //  - ModuleName
  //  - MethodName
  //  - JsonStr
  Execute(ctx context.Context, moduleName string, methodName string, jsonStr string) (r string, err error)
}

type WalletClient struct {
  c thrift.TClient
}

// Deprecated: Use NewWallet instead
func NewWalletClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *WalletClient {
  return &WalletClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewWallet instead
func NewWalletClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *WalletClient {
  return &WalletClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewWalletClient(c thrift.TClient) *WalletClient {
  return &WalletClient{
    c: c,
  }
}

// Parameters:
//  - ModuleName
//  - MethodName
//  - JsonStr
func (p *WalletClient) Execute(ctx context.Context, moduleName string, methodName string, jsonStr string) (r string, err error) {
  var _args0 WalletExecuteArgs
  _args0.ModuleName = moduleName
  _args0.MethodName = methodName
  _args0.JsonStr = jsonStr
  var _result1 WalletExecuteResult
  if err = p.c.Call(ctx, "execute", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type WalletProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Wallet
}

func (p *WalletProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *WalletProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *WalletProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewWalletProcessor(handler Wallet) *WalletProcessor {

  self2 := &WalletProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["execute"] = &walletProcessorExecute{handler:handler}
return self2
}

func (p *WalletProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type walletProcessorExecute struct {
  handler Wallet
}

func (p *walletProcessorExecute) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := WalletExecuteArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("execute", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := WalletExecuteResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Execute(ctx, args.ModuleName, args.MethodName, args.JsonStr); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing execute: " + err2.Error())
    oprot.WriteMessageBegin("execute", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("execute", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - ModuleName
//  - MethodName
//  - JsonStr
type WalletExecuteArgs struct {
  ModuleName string `thrift:"moduleName,1" db:"moduleName" json:"moduleName"`
  MethodName string `thrift:"methodName,2" db:"methodName" json:"methodName"`
  JsonStr string `thrift:"jsonStr,3" db:"jsonStr" json:"jsonStr"`
}

func NewWalletExecuteArgs() *WalletExecuteArgs {
  return &WalletExecuteArgs{}
}


func (p *WalletExecuteArgs) GetModuleName() string {
  return p.ModuleName
}

func (p *WalletExecuteArgs) GetMethodName() string {
  return p.MethodName
}

func (p *WalletExecuteArgs) GetJsonStr() string {
  return p.JsonStr
}
func (p *WalletExecuteArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *WalletExecuteArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ModuleName = v
}
  return nil
}

func (p *WalletExecuteArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.MethodName = v
}
  return nil
}

func (p *WalletExecuteArgs)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.JsonStr = v
}
  return nil
}

func (p *WalletExecuteArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("execute_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *WalletExecuteArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("moduleName", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:moduleName: ", p), err) }
  if err := oprot.WriteString(string(p.ModuleName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.moduleName (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:moduleName: ", p), err) }
  return err
}

func (p *WalletExecuteArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("methodName", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:methodName: ", p), err) }
  if err := oprot.WriteString(string(p.MethodName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.methodName (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:methodName: ", p), err) }
  return err
}

func (p *WalletExecuteArgs) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("jsonStr", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:jsonStr: ", p), err) }
  if err := oprot.WriteString(string(p.JsonStr)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.jsonStr (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:jsonStr: ", p), err) }
  return err
}

func (p *WalletExecuteArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("WalletExecuteArgs(%+v)", *p)
}

// Attributes:
//  - Success
type WalletExecuteResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewWalletExecuteResult() *WalletExecuteResult {
  return &WalletExecuteResult{}
}

var WalletExecuteResult_Success_DEFAULT string
func (p *WalletExecuteResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return WalletExecuteResult_Success_DEFAULT
  }
return *p.Success
}
func (p *WalletExecuteResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *WalletExecuteResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *WalletExecuteResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *WalletExecuteResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("execute_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *WalletExecuteResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *WalletExecuteResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("WalletExecuteResult(%+v)", *p)
}


