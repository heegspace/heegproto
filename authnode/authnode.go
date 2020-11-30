// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package authnode

import (
	"bytes"
	"context"
	"fmt"
	"reflect"

	"github.com/heegspace/heegproto/common"
	"github.com/heegspace/heegproto/rescode"
	"github.com/heegspace/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = rescode.GoUnusedProtection__
var _ = common.GoUnusedProtection__

// Attributes:
//  - Token
//  - Extra
type VerifyTokenReq struct {
	Token string            `thrift:"token,1" db:"token" json:"token"`
	Extra map[string]string `thrift:"extra,2" db:"extra" json:"extra"`
}

func NewVerifyTokenReq() *VerifyTokenReq {
	return &VerifyTokenReq{}
}

func (p *VerifyTokenReq) GetToken() string {
	return p.Token
}

func (p *VerifyTokenReq) GetExtra() map[string]string {
	return p.Extra
}
func (p *VerifyTokenReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			if fieldTypeId == thrift.MAP {
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

func (p *VerifyTokenReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *VerifyTokenReq) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val1 = v
		}
		p.Extra[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *VerifyTokenReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("verify_token_req"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyTokenReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *VerifyTokenReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:extra: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Extra {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:extra: ", p), err)
	}
	return err
}

func (p *VerifyTokenReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyTokenReq(%+v)", *p)
}

// Attributes:
//  - Rescode
//  - Extra
type VerifyTokenRes struct {
	Rescode rescode.Code      `thrift:"rescode,1" db:"rescode" json:"rescode"`
	Extra   map[string]string `thrift:"extra,2" db:"extra" json:"extra"`
}

func NewVerifyTokenRes() *VerifyTokenRes {
	return &VerifyTokenRes{}
}

func (p *VerifyTokenRes) GetRescode() rescode.Code {
	return p.Rescode
}

func (p *VerifyTokenRes) GetExtra() map[string]string {
	return p.Extra
}
func (p *VerifyTokenRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			if fieldTypeId == thrift.MAP {
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

func (p *VerifyTokenRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := rescode.Code(v)
		p.Rescode = temp
	}
	return nil
}

func (p *VerifyTokenRes) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key2 = v
		}
		var _val3 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val3 = v
		}
		p.Extra[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *VerifyTokenRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("verify_token_res"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyTokenRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("rescode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:rescode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Rescode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.rescode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:rescode: ", p), err)
	}
	return err
}

func (p *VerifyTokenRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:extra: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Extra {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:extra: ", p), err)
	}
	return err
}

func (p *VerifyTokenRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyTokenRes(%+v)", *p)
}

// Attributes:
//  - Token
//  - Extra
type AdminRoleReq struct {
	Token string            `thrift:"token,1" db:"token" json:"token"`
	Extra map[string]string `thrift:"extra,2" db:"extra" json:"extra"`
}

func NewAdminRoleReq() *AdminRoleReq {
	return &AdminRoleReq{}
}

func (p *AdminRoleReq) GetToken() string {
	return p.Token
}

func (p *AdminRoleReq) GetExtra() map[string]string {
	return p.Extra
}
func (p *AdminRoleReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			if fieldTypeId == thrift.MAP {
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

func (p *AdminRoleReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *AdminRoleReq) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key4 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key4 = v
		}
		var _val5 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val5 = v
		}
		p.Extra[_key4] = _val5
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *AdminRoleReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("admin_role_req"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AdminRoleReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *AdminRoleReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:extra: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Extra {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:extra: ", p), err)
	}
	return err
}

func (p *AdminRoleReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdminRoleReq(%+v)", *p)
}

// Attributes:
//  - Rescode
//  - Extra
type AdminRoleRes struct {
	Rescode rescode.Code      `thrift:"rescode,1" db:"rescode" json:"rescode"`
	Extra   map[string]string `thrift:"extra,2" db:"extra" json:"extra"`
}

func NewAdminRoleRes() *AdminRoleRes {
	return &AdminRoleRes{}
}

func (p *AdminRoleRes) GetRescode() rescode.Code {
	return p.Rescode
}

func (p *AdminRoleRes) GetExtra() map[string]string {
	return p.Extra
}
func (p *AdminRoleRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			if fieldTypeId == thrift.MAP {
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

func (p *AdminRoleRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := rescode.Code(v)
		p.Rescode = temp
	}
	return nil
}

func (p *AdminRoleRes) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key6 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key6 = v
		}
		var _val7 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val7 = v
		}
		p.Extra[_key6] = _val7
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *AdminRoleRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("admin_role_res"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AdminRoleRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("rescode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:rescode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Rescode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.rescode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:rescode: ", p), err)
	}
	return err
}

func (p *AdminRoleRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:extra: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Extra {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:extra: ", p), err)
	}
	return err
}

func (p *AdminRoleRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdminRoleRes(%+v)", *p)
}

// Attributes:
//  - Token
//  - Extra
type CoorRoleReq struct {
	Token string            `thrift:"token,1" db:"token" json:"token"`
	Extra map[string]string `thrift:"extra,2" db:"extra" json:"extra"`
}

func NewCoorRoleReq() *CoorRoleReq {
	return &CoorRoleReq{}
}

func (p *CoorRoleReq) GetToken() string {
	return p.Token
}

func (p *CoorRoleReq) GetExtra() map[string]string {
	return p.Extra
}
func (p *CoorRoleReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			if fieldTypeId == thrift.MAP {
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

func (p *CoorRoleReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *CoorRoleReq) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key8 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key8 = v
		}
		var _val9 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val9 = v
		}
		p.Extra[_key8] = _val9
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *CoorRoleReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("coor_role_req"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CoorRoleReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *CoorRoleReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:extra: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Extra {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:extra: ", p), err)
	}
	return err
}

func (p *CoorRoleReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CoorRoleReq(%+v)", *p)
}

// Attributes:
//  - Rescode
//  - Extra
type CoorRoleRes struct {
	Rescode rescode.Code      `thrift:"rescode,1" db:"rescode" json:"rescode"`
	Extra   map[string]string `thrift:"extra,2" db:"extra" json:"extra"`
}

func NewCoorRoleRes() *CoorRoleRes {
	return &CoorRoleRes{}
}

func (p *CoorRoleRes) GetRescode() rescode.Code {
	return p.Rescode
}

func (p *CoorRoleRes) GetExtra() map[string]string {
	return p.Extra
}
func (p *CoorRoleRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			if fieldTypeId == thrift.MAP {
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

func (p *CoorRoleRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := rescode.Code(v)
		p.Rescode = temp
	}
	return nil
}

func (p *CoorRoleRes) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key10 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key10 = v
		}
		var _val11 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val11 = v
		}
		p.Extra[_key10] = _val11
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *CoorRoleRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("coor_role_res"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CoorRoleRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("rescode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:rescode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Rescode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.rescode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:rescode: ", p), err)
	}
	return err
}

func (p *CoorRoleRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:extra: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Extra {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:extra: ", p), err)
	}
	return err
}

func (p *CoorRoleRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CoorRoleRes(%+v)", *p)
}

type AuthnodeService interface {
	// Parameters:
	//  - Req
	VerifyToken(ctx context.Context, req *VerifyTokenReq) (r *VerifyTokenRes, err error)
}

type AuthnodeServiceClient struct {
	c thrift.TClient
}

func NewAuthnodeServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AuthnodeServiceClient {
	return &AuthnodeServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAuthnodeServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AuthnodeServiceClient {
	return &AuthnodeServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAuthnodeServiceClient(c thrift.TClient) *AuthnodeServiceClient {
	return &AuthnodeServiceClient{
		c: c,
	}
}

func (p *AuthnodeServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Req
func (p *AuthnodeServiceClient) VerifyToken(ctx context.Context, req *VerifyTokenReq) (r *VerifyTokenRes, err error) {
	var _args12 AuthnodeServiceVerifyTokenArgs
	_args12.Req = req
	var _result13 AuthnodeServiceVerifyTokenResult
	if err = p.Client_().Call(ctx, "verify_token", &_args12, &_result13); err != nil {
		return
	}
	return _result13.GetSuccess(), nil
}

type AuthnodeServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AuthnodeService
}

func (p *AuthnodeServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AuthnodeServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AuthnodeServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAuthnodeServiceProcessor(handler AuthnodeService) *AuthnodeServiceProcessor {

	self14 := &AuthnodeServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self14.processorMap["verify_token"] = &authnodeServiceProcessorVerifyToken{handler: handler}
	return self14
}

func (p *AuthnodeServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x15 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x15.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x15

}

type authnodeServiceProcessorVerifyToken struct {
	handler AuthnodeService
}

func (p *authnodeServiceProcessorVerifyToken) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AuthnodeServiceVerifyTokenArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("verify_token", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AuthnodeServiceVerifyTokenResult{}
	var retval *VerifyTokenRes
	var err2 error
	if retval, err2 = p.handler.VerifyToken(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing verify_token: "+err2.Error())
		oprot.WriteMessageBegin("verify_token", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("verify_token", thrift.REPLY, seqId); err2 != nil {
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
//  - Req
type AuthnodeServiceVerifyTokenArgs struct {
	Req *VerifyTokenReq `thrift:"req,1" db:"req" json:"req"`
}

func NewAuthnodeServiceVerifyTokenArgs() *AuthnodeServiceVerifyTokenArgs {
	return &AuthnodeServiceVerifyTokenArgs{}
}

var AuthnodeServiceVerifyTokenArgs_Req_DEFAULT *VerifyTokenReq

func (p *AuthnodeServiceVerifyTokenArgs) GetReq() *VerifyTokenReq {
	if !p.IsSetReq() {
		return AuthnodeServiceVerifyTokenArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *AuthnodeServiceVerifyTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AuthnodeServiceVerifyTokenArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
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

func (p *AuthnodeServiceVerifyTokenArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = &VerifyTokenReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *AuthnodeServiceVerifyTokenArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("verify_token_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuthnodeServiceVerifyTokenArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *AuthnodeServiceVerifyTokenArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuthnodeServiceVerifyTokenArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AuthnodeServiceVerifyTokenResult struct {
	Success *VerifyTokenRes `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAuthnodeServiceVerifyTokenResult() *AuthnodeServiceVerifyTokenResult {
	return &AuthnodeServiceVerifyTokenResult{}
}

var AuthnodeServiceVerifyTokenResult_Success_DEFAULT *VerifyTokenRes

func (p *AuthnodeServiceVerifyTokenResult) GetSuccess() *VerifyTokenRes {
	if !p.IsSetSuccess() {
		return AuthnodeServiceVerifyTokenResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AuthnodeServiceVerifyTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AuthnodeServiceVerifyTokenResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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

func (p *AuthnodeServiceVerifyTokenResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &VerifyTokenRes{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AuthnodeServiceVerifyTokenResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("verify_token_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuthnodeServiceVerifyTokenResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AuthnodeServiceVerifyTokenResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuthnodeServiceVerifyTokenResult(%+v)", *p)
}
