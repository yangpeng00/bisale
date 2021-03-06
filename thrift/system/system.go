// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package system

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

// Attributes:
//  - ID
//  - KeyWord
//  - KeyValue
//  - Mark
//  - CreateTime
//  - UpdateTime
//  - Empty
type TSystemConfig struct {
  ID int32 `thrift:"id,1" db:"id" json:"id"`
  KeyWord string `thrift:"keyWord,2" db:"keyWord" json:"keyWord"`
  KeyValue string `thrift:"keyValue,3" db:"keyValue" json:"keyValue"`
  Mark string `thrift:"mark,4" db:"mark" json:"mark"`
  CreateTime string `thrift:"createTime,5" db:"createTime" json:"createTime"`
  UpdateTime string `thrift:"updateTime,6" db:"updateTime" json:"updateTime"`
  Empty bool `thrift:"empty,7" db:"empty" json:"empty,omitempty"`
}

func NewTSystemConfig() *TSystemConfig {
  return &TSystemConfig{}
}


func (p *TSystemConfig) GetID() int32 {
  return p.ID
}

func (p *TSystemConfig) GetKeyWord() string {
  return p.KeyWord
}

func (p *TSystemConfig) GetKeyValue() string {
  return p.KeyValue
}

func (p *TSystemConfig) GetMark() string {
  return p.Mark
}

func (p *TSystemConfig) GetCreateTime() string {
  return p.CreateTime
}

func (p *TSystemConfig) GetUpdateTime() string {
  return p.UpdateTime
}
var TSystemConfig_Empty_DEFAULT bool = false

func (p *TSystemConfig) GetEmpty() bool {
  return p.Empty
}
func (p *TSystemConfig) IsSetEmpty() bool {
  return p.Empty != TSystemConfig_Empty_DEFAULT
}

func (p *TSystemConfig) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.I32 {
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
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 7:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField7(iprot); err != nil {
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

func (p *TSystemConfig)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *TSystemConfig)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.KeyWord = v
}
  return nil
}

func (p *TSystemConfig)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.KeyValue = v
}
  return nil
}

func (p *TSystemConfig)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Mark = v
}
  return nil
}

func (p *TSystemConfig)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.CreateTime = v
}
  return nil
}

func (p *TSystemConfig)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.UpdateTime = v
}
  return nil
}

func (p *TSystemConfig)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  p.Empty = v
}
  return nil
}

func (p *TSystemConfig) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TSystemConfig"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField7(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TSystemConfig) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("id", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteI32(int32(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *TSystemConfig) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("keyWord", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:keyWord: ", p), err) }
  if err := oprot.WriteString(string(p.KeyWord)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.keyWord (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:keyWord: ", p), err) }
  return err
}

func (p *TSystemConfig) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("keyValue", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:keyValue: ", p), err) }
  if err := oprot.WriteString(string(p.KeyValue)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.keyValue (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:keyValue: ", p), err) }
  return err
}

func (p *TSystemConfig) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("mark", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:mark: ", p), err) }
  if err := oprot.WriteString(string(p.Mark)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.mark (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:mark: ", p), err) }
  return err
}

func (p *TSystemConfig) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("createTime", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:createTime: ", p), err) }
  if err := oprot.WriteString(string(p.CreateTime)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.createTime (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:createTime: ", p), err) }
  return err
}

func (p *TSystemConfig) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("updateTime", thrift.STRING, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:updateTime: ", p), err) }
  if err := oprot.WriteString(string(p.UpdateTime)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.updateTime (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:updateTime: ", p), err) }
  return err
}

func (p *TSystemConfig) writeField7(oprot thrift.TProtocol) (err error) {
  if p.IsSetEmpty() {
    if err := oprot.WriteFieldBegin("empty", thrift.BOOL, 7); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:empty: ", p), err) }
    if err := oprot.WriteBool(bool(p.Empty)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.empty (7) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 7:empty: ", p), err) }
  }
  return err
}

func (p *TSystemConfig) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSystemConfig(%+v)", *p)
}

// Attributes:
//  - KeyWord
//  - KeyValue
//  - Empty
type TSystemConfigResult_ struct {
  KeyWord string `thrift:"keyWord,1" db:"keyWord" json:"keyWord"`
  KeyValue string `thrift:"keyValue,2" db:"keyValue" json:"keyValue"`
  Empty bool `thrift:"empty,3" db:"empty" json:"empty,omitempty"`
}

func NewTSystemConfigResult_() *TSystemConfigResult_ {
  return &TSystemConfigResult_{}
}


func (p *TSystemConfigResult_) GetKeyWord() string {
  return p.KeyWord
}

func (p *TSystemConfigResult_) GetKeyValue() string {
  return p.KeyValue
}
var TSystemConfigResult__Empty_DEFAULT bool = false

func (p *TSystemConfigResult_) GetEmpty() bool {
  return p.Empty
}
func (p *TSystemConfigResult_) IsSetEmpty() bool {
  return p.Empty != TSystemConfigResult__Empty_DEFAULT
}

func (p *TSystemConfigResult_) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.BOOL {
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

func (p *TSystemConfigResult_)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.KeyWord = v
}
  return nil
}

func (p *TSystemConfigResult_)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.KeyValue = v
}
  return nil
}

func (p *TSystemConfigResult_)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Empty = v
}
  return nil
}

func (p *TSystemConfigResult_) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TSystemConfigResult"); err != nil {
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

func (p *TSystemConfigResult_) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("keyWord", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:keyWord: ", p), err) }
  if err := oprot.WriteString(string(p.KeyWord)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.keyWord (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:keyWord: ", p), err) }
  return err
}

func (p *TSystemConfigResult_) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("keyValue", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:keyValue: ", p), err) }
  if err := oprot.WriteString(string(p.KeyValue)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.keyValue (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:keyValue: ", p), err) }
  return err
}

func (p *TSystemConfigResult_) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetEmpty() {
    if err := oprot.WriteFieldBegin("empty", thrift.BOOL, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:empty: ", p), err) }
    if err := oprot.WriteBool(bool(p.Empty)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.empty (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:empty: ", p), err) }
  }
  return err
}

func (p *TSystemConfigResult_) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSystemConfigResult_(%+v)", *p)
}

type TSystemService interface {
  // Parameters:
  //  - TraceId
  //  - KeyWords
  SelectSystemConfigByKeyWords(ctx context.Context, traceId string, keyWords string) (r []*TSystemConfigResult_, err error)
  // Parameters:
  //  - TraceId
  //  - JsonValue
  UpdateSystemConfigByJsonValue(ctx context.Context, traceId string, jsonValue string) (r bool, err error)
}

type TSystemServiceClient struct {
  c thrift.TClient
}

// Deprecated: Use NewTSystemService instead
func NewTSystemServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TSystemServiceClient {
  return &TSystemServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewTSystemService instead
func NewTSystemServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TSystemServiceClient {
  return &TSystemServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewTSystemServiceClient(c thrift.TClient) *TSystemServiceClient {
  return &TSystemServiceClient{
    c: c,
  }
}

// Parameters:
//  - TraceId
//  - KeyWords
func (p *TSystemServiceClient) SelectSystemConfigByKeyWords(ctx context.Context, traceId string, keyWords string) (r []*TSystemConfigResult_, err error) {
  var _args0 TSystemServiceSelectSystemConfigByKeyWordsArgs
  _args0.TraceId = traceId
  _args0.KeyWords = keyWords
  var _result1 TSystemServiceSelectSystemConfigByKeyWordsResult
  if err = p.c.Call(ctx, "selectSystemConfigByKeyWords", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - TraceId
//  - JsonValue
func (p *TSystemServiceClient) UpdateSystemConfigByJsonValue(ctx context.Context, traceId string, jsonValue string) (r bool, err error) {
  var _args2 TSystemServiceUpdateSystemConfigByJsonValueArgs
  _args2.TraceId = traceId
  _args2.JsonValue = jsonValue
  var _result3 TSystemServiceUpdateSystemConfigByJsonValueResult
  if err = p.c.Call(ctx, "updateSystemConfigByJsonValue", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

type TSystemServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler TSystemService
}

func (p *TSystemServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TSystemServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TSystemServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTSystemServiceProcessor(handler TSystemService) *TSystemServiceProcessor {

  self4 := &TSystemServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["selectSystemConfigByKeyWords"] = &tSystemServiceProcessorSelectSystemConfigByKeyWords{handler:handler}
  self4.processorMap["updateSystemConfigByJsonValue"] = &tSystemServiceProcessorUpdateSystemConfigByJsonValue{handler:handler}
return self4
}

func (p *TSystemServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x5

}

type tSystemServiceProcessorSelectSystemConfigByKeyWords struct {
  handler TSystemService
}

func (p *tSystemServiceProcessorSelectSystemConfigByKeyWords) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TSystemServiceSelectSystemConfigByKeyWordsArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("selectSystemConfigByKeyWords", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TSystemServiceSelectSystemConfigByKeyWordsResult{}
var retval []*TSystemConfigResult_
  var err2 error
  if retval, err2 = p.handler.SelectSystemConfigByKeyWords(ctx, args.TraceId, args.KeyWords); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing selectSystemConfigByKeyWords: " + err2.Error())
    oprot.WriteMessageBegin("selectSystemConfigByKeyWords", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("selectSystemConfigByKeyWords", thrift.REPLY, seqId); err2 != nil {
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

type tSystemServiceProcessorUpdateSystemConfigByJsonValue struct {
  handler TSystemService
}

func (p *tSystemServiceProcessorUpdateSystemConfigByJsonValue) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TSystemServiceUpdateSystemConfigByJsonValueArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("updateSystemConfigByJsonValue", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TSystemServiceUpdateSystemConfigByJsonValueResult{}
var retval bool
  var err2 error
  if retval, err2 = p.handler.UpdateSystemConfigByJsonValue(ctx, args.TraceId, args.JsonValue); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing updateSystemConfigByJsonValue: " + err2.Error())
    oprot.WriteMessageBegin("updateSystemConfigByJsonValue", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("updateSystemConfigByJsonValue", thrift.REPLY, seqId); err2 != nil {
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
//  - TraceId
//  - KeyWords
type TSystemServiceSelectSystemConfigByKeyWordsArgs struct {
  TraceId string `thrift:"traceId,1" db:"traceId" json:"traceId"`
  KeyWords string `thrift:"keyWords,2" db:"keyWords" json:"keyWords"`
}

func NewTSystemServiceSelectSystemConfigByKeyWordsArgs() *TSystemServiceSelectSystemConfigByKeyWordsArgs {
  return &TSystemServiceSelectSystemConfigByKeyWordsArgs{}
}


func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) GetTraceId() string {
  return p.TraceId
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) GetKeyWords() string {
  return p.KeyWords
}
func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.TraceId = v
}
  return nil
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.KeyWords = v
}
  return nil
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("selectSystemConfigByKeyWords_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("traceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:traceId: ", p), err) }
  if err := oprot.WriteString(string(p.TraceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.traceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:traceId: ", p), err) }
  return err
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("keyWords", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:keyWords: ", p), err) }
  if err := oprot.WriteString(string(p.KeyWords)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.keyWords (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:keyWords: ", p), err) }
  return err
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSystemServiceSelectSystemConfigByKeyWordsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TSystemServiceSelectSystemConfigByKeyWordsResult struct {
  Success []*TSystemConfigResult_ `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTSystemServiceSelectSystemConfigByKeyWordsResult() *TSystemServiceSelectSystemConfigByKeyWordsResult {
  return &TSystemServiceSelectSystemConfigByKeyWordsResult{}
}

var TSystemServiceSelectSystemConfigByKeyWordsResult_Success_DEFAULT []*TSystemConfigResult_

func (p *TSystemServiceSelectSystemConfigByKeyWordsResult) GetSuccess() []*TSystemConfigResult_ {
  return p.Success
}
func (p *TSystemServiceSelectSystemConfigByKeyWordsResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsResult) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.LIST {
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

func (p *TSystemServiceSelectSystemConfigByKeyWordsResult)  ReadField0(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*TSystemConfigResult_, 0, size)
  p.Success =  tSlice
  for i := 0; i < size; i ++ {
    _elem6 := &TSystemConfigResult_{}
    if err := _elem6.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem6), err)
    }
    p.Success = append(p.Success, _elem6)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("selectSystemConfigByKeyWords_result"); err != nil {
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

func (p *TSystemServiceSelectSystemConfigByKeyWordsResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.Success {
      if err := v.Write(oprot); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
      }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TSystemServiceSelectSystemConfigByKeyWordsResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSystemServiceSelectSystemConfigByKeyWordsResult(%+v)", *p)
}

// Attributes:
//  - TraceId
//  - JsonValue
type TSystemServiceUpdateSystemConfigByJsonValueArgs struct {
  TraceId string `thrift:"traceId,1" db:"traceId" json:"traceId"`
  JsonValue string `thrift:"jsonValue,2" db:"jsonValue" json:"jsonValue"`
}

func NewTSystemServiceUpdateSystemConfigByJsonValueArgs() *TSystemServiceUpdateSystemConfigByJsonValueArgs {
  return &TSystemServiceUpdateSystemConfigByJsonValueArgs{}
}


func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) GetTraceId() string {
  return p.TraceId
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) GetJsonValue() string {
  return p.JsonValue
}
func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) Read(iprot thrift.TProtocol) error {
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

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.TraceId = v
}
  return nil
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.JsonValue = v
}
  return nil
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("updateSystemConfigByJsonValue_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("traceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:traceId: ", p), err) }
  if err := oprot.WriteString(string(p.TraceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.traceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:traceId: ", p), err) }
  return err
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("jsonValue", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:jsonValue: ", p), err) }
  if err := oprot.WriteString(string(p.JsonValue)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.jsonValue (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:jsonValue: ", p), err) }
  return err
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSystemServiceUpdateSystemConfigByJsonValueArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TSystemServiceUpdateSystemConfigByJsonValueResult struct {
  Success *bool `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTSystemServiceUpdateSystemConfigByJsonValueResult() *TSystemServiceUpdateSystemConfigByJsonValueResult {
  return &TSystemServiceUpdateSystemConfigByJsonValueResult{}
}

var TSystemServiceUpdateSystemConfigByJsonValueResult_Success_DEFAULT bool
func (p *TSystemServiceUpdateSystemConfigByJsonValueResult) GetSuccess() bool {
  if !p.IsSetSuccess() {
    return TSystemServiceUpdateSystemConfigByJsonValueResult_Success_DEFAULT
  }
return *p.Success
}
func (p *TSystemServiceUpdateSystemConfigByJsonValueResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueResult) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.BOOL {
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

func (p *TSystemServiceUpdateSystemConfigByJsonValueResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("updateSystemConfigByJsonValue_result"); err != nil {
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

func (p *TSystemServiceUpdateSystemConfigByJsonValueResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TSystemServiceUpdateSystemConfigByJsonValueResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSystemServiceUpdateSystemConfigByJsonValueResult(%+v)", *p)
}


