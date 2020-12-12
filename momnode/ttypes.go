// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package momnode

import (
	"fmt"
	"math"

	"github.com/heegspace/heegproto/common"
	"github.com/heegspace/heegproto/rescode"
	"github.com/heegspace/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var _ = rescode.GoUnusedProtection__
var _ = common.GoUnusedProtection__
var GoUnusedProtection__ int

type MomentsCountReq struct {
	Auth *common.Authorize `thrift:"auth,1"`
	Uid  int64             `thrift:"uid,2"`
}

func NewMomentsCountReq() *MomentsCountReq {
	return &MomentsCountReq{}
}

func (p *MomentsCountReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *MomentsCountReq) readField1(iprot thrift.TProtocol) error {
	p.Auth = common.NewAuthorize()
	if err := p.Auth.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Auth)
	}
	return nil
}

func (p *MomentsCountReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Uid = v
	}
	return nil
}

func (p *MomentsCountReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("moments_count_req"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *MomentsCountReq) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Auth != nil {
		if err := oprot.WriteFieldBegin("auth", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:auth: %s", p, err)
		}
		if err := p.Auth.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Auth)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:auth: %s", p, err)
		}
	}
	return err
}

func (p *MomentsCountReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:uid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Uid)); err != nil {
		return fmt.Errorf("%T.uid (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:uid: %s", p, err)
	}
	return err
}

func (p *MomentsCountReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MomentsCountReq(%+v)", *p)
}

type MomentsCountRes struct {
	Rescode rescode.Code      `thrift:"rescode,1"`
	Resmsg  string            `thrift:"resmsg,2"`
	Count   int32             `thrift:"count,3"`
	Extra   map[string]string `thrift:"extra,4"`
}

func NewMomentsCountRes() *MomentsCountRes {
	return &MomentsCountRes{
		Rescode: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *MomentsCountRes) IsSetRescode() bool {
	return int64(p.Rescode) != math.MinInt32-1
}

func (p *MomentsCountRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *MomentsCountRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Rescode = rescode.Code(v)
	}
	return nil
}

func (p *MomentsCountRes) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Resmsg = v
	}
	return nil
}

func (p *MomentsCountRes) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Count = v
	}
	return nil
}

func (p *MomentsCountRes) readField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s")
	}
	p.Extra = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_val1 = v
		}
		p.Extra[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s")
	}
	return nil
}

func (p *MomentsCountRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("moments_count_res"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *MomentsCountRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetRescode() {
		if err := oprot.WriteFieldBegin("rescode", thrift.I32, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:rescode: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Rescode)); err != nil {
			return fmt.Errorf("%T.rescode (1) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:rescode: %s", p, err)
		}
	}
	return err
}

func (p *MomentsCountRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("resmsg", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:resmsg: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Resmsg)); err != nil {
		return fmt.Errorf("%T.resmsg (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:resmsg: %s", p, err)
	}
	return err
}

func (p *MomentsCountRes) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("count", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:count: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Count)); err != nil {
		return fmt.Errorf("%T.count (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:count: %s", p, err)
	}
	return err
}

func (p *MomentsCountRes) writeField4(oprot thrift.TProtocol) (err error) {
	if p.Extra != nil {
		if err := oprot.WriteFieldBegin("extra", thrift.MAP, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:extra: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
			return fmt.Errorf("error writing map begin: %s")
		}
		for k, v := range p.Extra {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:extra: %s", p, err)
		}
	}
	return err
}

func (p *MomentsCountRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MomentsCountRes(%+v)", *p)
}

type Extra struct {
	Path   string `thrift:"path,1"`
	TypeA1 int32  `thrift:"type,2"`
}

func NewExtra() *Extra {
	return &Extra{}
}

func (p *Extra) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Extra) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Path = v
	}
	return nil
}

func (p *Extra) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.TypeA1 = v
	}
	return nil
}

func (p *Extra) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("extra"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Extra) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("path", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:path: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Path)); err != nil {
		return fmt.Errorf("%T.path (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:path: %s", p, err)
	}
	return err
}

func (p *Extra) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("type", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:type: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.TypeA1)); err != nil {
		return fmt.Errorf("%T.type (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:type: %s", p, err)
	}
	return err
}

func (p *Extra) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Extra(%+v)", *p)
}

type Moments struct {
	Id       int32    `thrift:"id,1"`
	Mid      int64    `thrift:"mid,2"`
	Text     string   `thrift:"text,3"`
	Extra    []*Extra `thrift:"extra,4"`
	CreateAt int32    `thrift:"create_at,5"`
}

func NewMoments() *Moments {
	return &Moments{}
}

func (p *Moments) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Moments) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Id = v
	}
	return nil
}

func (p *Moments) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Mid = v
	}
	return nil
}

func (p *Moments) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Text = v
	}
	return nil
}

func (p *Moments) readField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Extra = make([]*Extra, 0, size)
	for i := 0; i < size; i++ {
		_elem2 := NewExtra()
		if err := _elem2.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem2)
		}
		p.Extra = append(p.Extra, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *Moments) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 5: %s")
	} else {
		p.CreateAt = v
	}
	return nil
}

func (p *Moments) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("moments"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Moments) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:id: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Id)); err != nil {
		return fmt.Errorf("%T.id (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:id: %s", p, err)
	}
	return err
}

func (p *Moments) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:mid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Mid)); err != nil {
		return fmt.Errorf("%T.mid (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:mid: %s", p, err)
	}
	return err
}

func (p *Moments) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("text", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:text: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Text)); err != nil {
		return fmt.Errorf("%T.text (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:text: %s", p, err)
	}
	return err
}

func (p *Moments) writeField4(oprot thrift.TProtocol) (err error) {
	if p.Extra != nil {
		if err := oprot.WriteFieldBegin("extra", thrift.LIST, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:extra: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Extra)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Extra {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:extra: %s", p, err)
		}
	}
	return err
}

func (p *Moments) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("create_at", thrift.I32, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:create_at: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.CreateAt)); err != nil {
		return fmt.Errorf("%T.create_at (5) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:create_at: %s", p, err)
	}
	return err
}

func (p *Moments) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Moments(%+v)", *p)
}

type AddMomentsReq struct {
	Auth    *common.Authorize `thrift:"auth,1"`
	Uid     int64             `thrift:"uid,2"`
	Moments *Moments          `thrift:"moments,3"`
	Extra   map[string]string `thrift:"extra,4"`
}

func NewAddMomentsReq() *AddMomentsReq {
	return &AddMomentsReq{}
}

func (p *AddMomentsReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddMomentsReq) readField1(iprot thrift.TProtocol) error {
	p.Auth = common.NewAuthorize()
	if err := p.Auth.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Auth)
	}
	return nil
}

func (p *AddMomentsReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Uid = v
	}
	return nil
}

func (p *AddMomentsReq) readField3(iprot thrift.TProtocol) error {
	p.Moments = NewMoments()
	if err := p.Moments.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Moments)
	}
	return nil
}

func (p *AddMomentsReq) readField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s")
	}
	p.Extra = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key3 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_key3 = v
		}
		var _val4 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_val4 = v
		}
		p.Extra[_key3] = _val4
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s")
	}
	return nil
}

func (p *AddMomentsReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_moments_req"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AddMomentsReq) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Auth != nil {
		if err := oprot.WriteFieldBegin("auth", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:auth: %s", p, err)
		}
		if err := p.Auth.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Auth)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:auth: %s", p, err)
		}
	}
	return err
}

func (p *AddMomentsReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:uid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Uid)); err != nil {
		return fmt.Errorf("%T.uid (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:uid: %s", p, err)
	}
	return err
}

func (p *AddMomentsReq) writeField3(oprot thrift.TProtocol) (err error) {
	if p.Moments != nil {
		if err := oprot.WriteFieldBegin("moments", thrift.STRUCT, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:moments: %s", p, err)
		}
		if err := p.Moments.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Moments)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:moments: %s", p, err)
		}
	}
	return err
}

func (p *AddMomentsReq) writeField4(oprot thrift.TProtocol) (err error) {
	if p.Extra != nil {
		if err := oprot.WriteFieldBegin("extra", thrift.MAP, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:extra: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
			return fmt.Errorf("error writing map begin: %s")
		}
		for k, v := range p.Extra {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:extra: %s", p, err)
		}
	}
	return err
}

func (p *AddMomentsReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddMomentsReq(%+v)", *p)
}

type AddMomentsRes struct {
	Rescode rescode.Code      `thrift:"rescode,1"`
	Resmsg  string            `thrift:"resmsg,2"`
	Extra   map[string]string `thrift:"extra,3"`
}

func NewAddMomentsRes() *AddMomentsRes {
	return &AddMomentsRes{
		Rescode: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *AddMomentsRes) IsSetRescode() bool {
	return int64(p.Rescode) != math.MinInt32-1
}

func (p *AddMomentsRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddMomentsRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Rescode = rescode.Code(v)
	}
	return nil
}

func (p *AddMomentsRes) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Resmsg = v
	}
	return nil
}

func (p *AddMomentsRes) readField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s")
	}
	p.Extra = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key5 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_key5 = v
		}
		var _val6 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_val6 = v
		}
		p.Extra[_key5] = _val6
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s")
	}
	return nil
}

func (p *AddMomentsRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_moments_res"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AddMomentsRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetRescode() {
		if err := oprot.WriteFieldBegin("rescode", thrift.I32, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:rescode: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Rescode)); err != nil {
			return fmt.Errorf("%T.rescode (1) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:rescode: %s", p, err)
		}
	}
	return err
}

func (p *AddMomentsRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("resmsg", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:resmsg: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Resmsg)); err != nil {
		return fmt.Errorf("%T.resmsg (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:resmsg: %s", p, err)
	}
	return err
}

func (p *AddMomentsRes) writeField3(oprot thrift.TProtocol) (err error) {
	if p.Extra != nil {
		if err := oprot.WriteFieldBegin("extra", thrift.MAP, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:extra: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
			return fmt.Errorf("error writing map begin: %s")
		}
		for k, v := range p.Extra {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:extra: %s", p, err)
		}
	}
	return err
}

func (p *AddMomentsRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddMomentsRes(%+v)", *p)
}

type MomentsListReq struct {
	Auth *common.Authorize `thrift:"auth,1"`
	Uid  int64             `thrift:"uid,2"`
	Page int32             `thrift:"page,3"`
	Size int32             `thrift:"size,4"`
}

func NewMomentsListReq() *MomentsListReq {
	return &MomentsListReq{}
}

func (p *MomentsListReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *MomentsListReq) readField1(iprot thrift.TProtocol) error {
	p.Auth = common.NewAuthorize()
	if err := p.Auth.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Auth)
	}
	return nil
}

func (p *MomentsListReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Uid = v
	}
	return nil
}

func (p *MomentsListReq) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Page = v
	}
	return nil
}

func (p *MomentsListReq) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 4: %s")
	} else {
		p.Size = v
	}
	return nil
}

func (p *MomentsListReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("moments_list_req"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *MomentsListReq) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Auth != nil {
		if err := oprot.WriteFieldBegin("auth", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:auth: %s", p, err)
		}
		if err := p.Auth.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Auth)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:auth: %s", p, err)
		}
	}
	return err
}

func (p *MomentsListReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:uid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Uid)); err != nil {
		return fmt.Errorf("%T.uid (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:uid: %s", p, err)
	}
	return err
}

func (p *MomentsListReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("page", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:page: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Page)); err != nil {
		return fmt.Errorf("%T.page (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:page: %s", p, err)
	}
	return err
}

func (p *MomentsListReq) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("size", thrift.I32, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:size: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Size)); err != nil {
		return fmt.Errorf("%T.size (4) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:size: %s", p, err)
	}
	return err
}

func (p *MomentsListReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MomentsListReq(%+v)", *p)
}

type MomentsListRes struct {
	Rescode rescode.Code `thrift:"rescode,1"`
	Resmsg  string       `thrift:"resmsg,2"`
	Data    []*Moments   `thrift:"data,3"`
}

func NewMomentsListRes() *MomentsListRes {
	return &MomentsListRes{
		Rescode: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *MomentsListRes) IsSetRescode() bool {
	return int64(p.Rescode) != math.MinInt32-1
}

func (p *MomentsListRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *MomentsListRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Rescode = rescode.Code(v)
	}
	return nil
}

func (p *MomentsListRes) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Resmsg = v
	}
	return nil
}

func (p *MomentsListRes) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Data = make([]*Moments, 0, size)
	for i := 0; i < size; i++ {
		_elem7 := NewMoments()
		if err := _elem7.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem7)
		}
		p.Data = append(p.Data, _elem7)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *MomentsListRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("moments_list_res"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *MomentsListRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetRescode() {
		if err := oprot.WriteFieldBegin("rescode", thrift.I32, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:rescode: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Rescode)); err != nil {
			return fmt.Errorf("%T.rescode (1) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:rescode: %s", p, err)
		}
	}
	return err
}

func (p *MomentsListRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("resmsg", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:resmsg: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Resmsg)); err != nil {
		return fmt.Errorf("%T.resmsg (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:resmsg: %s", p, err)
	}
	return err
}

func (p *MomentsListRes) writeField3(oprot thrift.TProtocol) (err error) {
	if p.Data != nil {
		if err := oprot.WriteFieldBegin("data", thrift.LIST, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:data: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Data)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Data {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:data: %s", p, err)
		}
	}
	return err
}

func (p *MomentsListRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MomentsListRes(%+v)", *p)
}