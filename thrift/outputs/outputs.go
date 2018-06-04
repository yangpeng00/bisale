// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package outputs

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
//  - Token
type LoginOutput struct {
  Token string `thrift:"token,1" db:"token" json:"token"`
}

func NewLoginOutput() *LoginOutput {
  return &LoginOutput{}
}


func (p *LoginOutput) GetToken() string {
  return p.Token
}
func (p *LoginOutput) Read(iprot thrift.TProtocol) error {
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

func (p *LoginOutput)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Token = v
}
  return nil
}

func (p *LoginOutput) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("LoginOutput"); err != nil {
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

func (p *LoginOutput) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err) }
  if err := oprot.WriteString(string(p.Token)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err) }
  return err
}

func (p *LoginOutput) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoginOutput(%+v)", *p)
}

// Attributes:
//  - Token
type RegisterOutput struct {
  Token string `thrift:"token,1" db:"token" json:"token"`
}

func NewRegisterOutput() *RegisterOutput {
  return &RegisterOutput{}
}


func (p *RegisterOutput) GetToken() string {
  return p.Token
}
func (p *RegisterOutput) Read(iprot thrift.TProtocol) error {
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

func (p *RegisterOutput)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Token = v
}
  return nil
}

func (p *RegisterOutput) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("RegisterOutput"); err != nil {
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

func (p *RegisterOutput) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err) }
  if err := oprot.WriteString(string(p.Token)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err) }
  return err
}

func (p *RegisterOutput) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RegisterOutput(%+v)", *p)
}

// Attributes:
//  - MemberID
type CreateMemberOutput struct {
  MemberID string `thrift:"member_id,1" db:"member_id" json:"member_id"`
}

func NewCreateMemberOutput() *CreateMemberOutput {
  return &CreateMemberOutput{}
}


func (p *CreateMemberOutput) GetMemberID() string {
  return p.MemberID
}
func (p *CreateMemberOutput) Read(iprot thrift.TProtocol) error {
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

func (p *CreateMemberOutput)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.MemberID = v
}
  return nil
}

func (p *CreateMemberOutput) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("CreateMemberOutput"); err != nil {
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

func (p *CreateMemberOutput) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("member_id", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:member_id: ", p), err) }
  if err := oprot.WriteString(string(p.MemberID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.member_id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:member_id: ", p), err) }
  return err
}

func (p *CreateMemberOutput) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("CreateMemberOutput(%+v)", *p)
}

// Attributes:
//  - Valid
//  - MemberId
type JWTOutput struct {
  Valid bool `thrift:"Valid,1" db:"Valid" json:"Valid"`
  MemberId string `thrift:"MemberId,2" db:"MemberId" json:"MemberId"`
}

func NewJWTOutput() *JWTOutput {
  return &JWTOutput{}
}


func (p *JWTOutput) GetValid() bool {
  return p.Valid
}

func (p *JWTOutput) GetMemberId() string {
  return p.MemberId
}
func (p *JWTOutput) Read(iprot thrift.TProtocol) error {
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

func (p *JWTOutput)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Valid = v
}
  return nil
}

func (p *JWTOutput)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.MemberId = v
}
  return nil
}

func (p *JWTOutput) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("JWTOutput"); err != nil {
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

func (p *JWTOutput) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("Valid", thrift.BOOL, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Valid: ", p), err) }
  if err := oprot.WriteBool(bool(p.Valid)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.Valid (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Valid: ", p), err) }
  return err
}

func (p *JWTOutput) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("MemberId", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:MemberId: ", p), err) }
  if err := oprot.WriteString(string(p.MemberId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.MemberId (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:MemberId: ", p), err) }
  return err
}

func (p *JWTOutput) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("JWTOutput(%+v)", *p)
}

