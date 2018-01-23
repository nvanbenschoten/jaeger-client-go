// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - Healthy
//  - Msg
type HealthCheckRes struct {
  Healthy bool `thrift:"healthy,1" db:"healthy" json:"healthy"`
  Msg string `thrift:"msg,2" db:"msg" json:"msg"`
}

func NewHealthCheckRes() *HealthCheckRes {
  return &HealthCheckRes{}
}


func (p *HealthCheckRes) GetHealthy() bool {
  return p.Healthy
}

func (p *HealthCheckRes) GetMsg() string {
  return p.Msg
}
func (p *HealthCheckRes) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.BOOL {
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

func (p *HealthCheckRes)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Healthy = v
}
  return nil
}

func (p *HealthCheckRes)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *HealthCheckRes) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("HealthCheckRes"); err != nil {
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

func (p *HealthCheckRes) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("healthy", thrift.BOOL, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:healthy: ", p), err) }
  if err := oprot.WriteBool(bool(p.Healthy)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.healthy (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:healthy: ", p), err) }
  return err
}

func (p *HealthCheckRes) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:msg: ", p), err) }
  if err := oprot.WriteString(string(p.Msg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:msg: ", p), err) }
  return err
}

func (p *HealthCheckRes) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HealthCheckRes(%+v)", *p)
}

type Base interface {
  BaseCall(ctx context.Context) (err error)
}

type BaseClient struct {
  c thrift.TClient
}

// Deprecated: Use NewBase instead
func NewBaseClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BaseClient {
  return &BaseClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewBase instead
func NewBaseClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BaseClient {
  return &BaseClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewBaseClient(c thrift.TClient) *BaseClient {
  return &BaseClient{
    c: c,
  }
}

func (p *BaseClient) BaseCall(ctx context.Context) (err error) {
  var _args0 BaseBaseCallArgs
  var _result1 BaseBaseCallResult
  if err = p.c.Call(ctx, "BaseCall", &_args0, &_result1); err != nil {
    return
  }
  return nil
}

type BaseProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Base
}

func (p *BaseProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *BaseProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *BaseProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewBaseProcessor(handler Base) *BaseProcessor {

  self2 := &BaseProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["BaseCall"] = &baseProcessorBaseCall{handler:handler}
return self2
}

func (p *BaseProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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
  oprot.Flush()
  return false, x3

}

type baseProcessorBaseCall struct {
  handler Base
}

func (p *baseProcessorBaseCall) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := BaseBaseCallArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("BaseCall", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := BaseBaseCallResult{}
  var err2 error
  if err2 = p.handler.BaseCall(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing BaseCall: " + err2.Error())
    oprot.WriteMessageBegin("BaseCall", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("BaseCall", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type BaseBaseCallArgs struct {
}

func NewBaseBaseCallArgs() *BaseBaseCallArgs {
  return &BaseBaseCallArgs{}
}

func (p *BaseBaseCallArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *BaseBaseCallArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("BaseCall_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *BaseBaseCallArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("BaseBaseCallArgs(%+v)", *p)
}

type BaseBaseCallResult struct {
}

func NewBaseBaseCallResult() *BaseBaseCallResult {
  return &BaseBaseCallResult{}
}

func (p *BaseBaseCallResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *BaseBaseCallResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("BaseCall_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *BaseBaseCallResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("BaseBaseCallResult(%+v)", *p)
}


type First interface {
Base

  // Parameters:
  //  - Msg
  Echo(ctx context.Context, msg string) (r string, err error)
  Healthcheck(ctx context.Context) (r *HealthCheckRes, err error)
  AppError(ctx context.Context) (err error)
}

type FirstClient struct {
  c thrift.TClient
  *BaseClient
}

// Deprecated: Use NewFirst instead
func NewFirstClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FirstClient {
  return &FirstClient{BaseClient: NewBaseClientFactory(t, f)}}

// Deprecated: Use NewFirst instead
func NewFirstClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FirstClient {
  return &FirstClient{BaseClient: NewBaseClientProtocol(t, iprot, oprot)}
}

func NewFirstClient(c thrift.TClient) *FirstClient {
  return &FirstClient{
    c: c,
    BaseClient: NewBaseClient(c),
  }
}

// Parameters:
//  - Msg
func (p *FirstClient) Echo(ctx context.Context, msg string) (r string, err error) {
  var _args4 FirstEchoArgs
  _args4.Msg = msg
  var _result5 FirstEchoResult
  if err = p.c.Call(ctx, "Echo", &_args4, &_result5); err != nil {
    return
  }
  return _result5.GetSuccess(), nil
}

func (p *FirstClient) Healthcheck(ctx context.Context) (r *HealthCheckRes, err error) {
  var _args6 FirstHealthcheckArgs
  var _result7 FirstHealthcheckResult
  if err = p.c.Call(ctx, "Healthcheck", &_args6, &_result7); err != nil {
    return
  }
  return _result7.GetSuccess(), nil
}

func (p *FirstClient) AppError(ctx context.Context) (err error) {
  var _args8 FirstAppErrorArgs
  var _result9 FirstAppErrorResult
  if err = p.c.Call(ctx, "AppError", &_args8, &_result9); err != nil {
    return
  }
  return nil
}

type FirstProcessor struct {
  *BaseProcessor
}

func NewFirstProcessor(handler First) *FirstProcessor {
  self10 := &FirstProcessor{NewBaseProcessor(handler)}
  self10.AddToProcessorMap("Echo", &firstProcessorEcho{handler:handler})
  self10.AddToProcessorMap("Healthcheck", &firstProcessorHealthcheck{handler:handler})
  self10.AddToProcessorMap("AppError", &firstProcessorAppError{handler:handler})
  return self10
}

type firstProcessorEcho struct {
  handler First
}

func (p *firstProcessorEcho) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := FirstEchoArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Echo", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := FirstEchoResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Echo(ctx, args.Msg); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Echo: " + err2.Error())
    oprot.WriteMessageBegin("Echo", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Echo", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type firstProcessorHealthcheck struct {
  handler First
}

func (p *firstProcessorHealthcheck) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := FirstHealthcheckArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Healthcheck", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := FirstHealthcheckResult{}
var retval *HealthCheckRes
  var err2 error
  if retval, err2 = p.handler.Healthcheck(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Healthcheck: " + err2.Error())
    oprot.WriteMessageBegin("Healthcheck", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("Healthcheck", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type firstProcessorAppError struct {
  handler First
}

func (p *firstProcessorAppError) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := FirstAppErrorArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("AppError", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := FirstAppErrorResult{}
  var err2 error
  if err2 = p.handler.AppError(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing AppError: " + err2.Error())
    oprot.WriteMessageBegin("AppError", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("AppError", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Msg
type FirstEchoArgs struct {
  Msg string `thrift:"msg,1" db:"msg" json:"msg"`
}

func NewFirstEchoArgs() *FirstEchoArgs {
  return &FirstEchoArgs{}
}


func (p *FirstEchoArgs) GetMsg() string {
  return p.Msg
}
func (p *FirstEchoArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FirstEchoArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *FirstEchoArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Echo_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FirstEchoArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:msg: ", p), err) }
  if err := oprot.WriteString(string(p.Msg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:msg: ", p), err) }
  return err
}

func (p *FirstEchoArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FirstEchoArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FirstEchoResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFirstEchoResult() *FirstEchoResult {
  return &FirstEchoResult{}
}

var FirstEchoResult_Success_DEFAULT string
func (p *FirstEchoResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return FirstEchoResult_Success_DEFAULT
  }
return *p.Success
}
func (p *FirstEchoResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FirstEchoResult) Read(iprot thrift.TProtocol) error {
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

func (p *FirstEchoResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *FirstEchoResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Echo_result"); err != nil {
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

func (p *FirstEchoResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *FirstEchoResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FirstEchoResult(%+v)", *p)
}

type FirstHealthcheckArgs struct {
}

func NewFirstHealthcheckArgs() *FirstHealthcheckArgs {
  return &FirstHealthcheckArgs{}
}

func (p *FirstHealthcheckArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *FirstHealthcheckArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Healthcheck_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FirstHealthcheckArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FirstHealthcheckArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FirstHealthcheckResult struct {
  Success *HealthCheckRes `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFirstHealthcheckResult() *FirstHealthcheckResult {
  return &FirstHealthcheckResult{}
}

var FirstHealthcheckResult_Success_DEFAULT *HealthCheckRes
func (p *FirstHealthcheckResult) GetSuccess() *HealthCheckRes {
  if !p.IsSetSuccess() {
    return FirstHealthcheckResult_Success_DEFAULT
  }
return p.Success
}
func (p *FirstHealthcheckResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FirstHealthcheckResult) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.STRUCT {
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

func (p *FirstHealthcheckResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &HealthCheckRes{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *FirstHealthcheckResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Healthcheck_result"); err != nil {
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

func (p *FirstHealthcheckResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *FirstHealthcheckResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FirstHealthcheckResult(%+v)", *p)
}

type FirstAppErrorArgs struct {
}

func NewFirstAppErrorArgs() *FirstAppErrorArgs {
  return &FirstAppErrorArgs{}
}

func (p *FirstAppErrorArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *FirstAppErrorArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("AppError_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FirstAppErrorArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FirstAppErrorArgs(%+v)", *p)
}

type FirstAppErrorResult struct {
}

func NewFirstAppErrorResult() *FirstAppErrorResult {
  return &FirstAppErrorResult{}
}

func (p *FirstAppErrorResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *FirstAppErrorResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("AppError_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FirstAppErrorResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FirstAppErrorResult(%+v)", *p)
}


type Second interface {
  Test(ctx context.Context) (err error)
}

type SecondClient struct {
  c thrift.TClient
}

// Deprecated: Use NewSecond instead
func NewSecondClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SecondClient {
  return &SecondClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewSecond instead
func NewSecondClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SecondClient {
  return &SecondClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewSecondClient(c thrift.TClient) *SecondClient {
  return &SecondClient{
    c: c,
  }
}

func (p *SecondClient) Test(ctx context.Context) (err error) {
  var _args12 SecondTestArgs
  var _result13 SecondTestResult
  if err = p.c.Call(ctx, "Test", &_args12, &_result13); err != nil {
    return
  }
  return nil
}

type SecondProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Second
}

func (p *SecondProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *SecondProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *SecondProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewSecondProcessor(handler Second) *SecondProcessor {

  self14 := &SecondProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self14.processorMap["Test"] = &secondProcessorTest{handler:handler}
return self14
}

func (p *SecondProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x15 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x15.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x15

}

type secondProcessorTest struct {
  handler Second
}

func (p *secondProcessorTest) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := SecondTestArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Test", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := SecondTestResult{}
  var err2 error
  if err2 = p.handler.Test(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Test: " + err2.Error())
    oprot.WriteMessageBegin("Test", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("Test", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type SecondTestArgs struct {
}

func NewSecondTestArgs() *SecondTestArgs {
  return &SecondTestArgs{}
}

func (p *SecondTestArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *SecondTestArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Test_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SecondTestArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SecondTestArgs(%+v)", *p)
}

type SecondTestResult struct {
}

func NewSecondTestResult() *SecondTestResult {
  return &SecondTestResult{}
}

func (p *SecondTestResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *SecondTestResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Test_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SecondTestResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SecondTestResult(%+v)", *p)
}


