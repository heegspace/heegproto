// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package searchnode

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

type SearchQuestionReq struct {
	Auth    *common.Authorize `thrift:"auth,1"`
	Uid     int64             `thrift:"uid,2"`
	Keyword string            `thrift:"keyword,3"`
	Page    int32             `thrift:"page,4"`
	Size    int32             `thrift:"size,5"`
	Extra   map[string]string `thrift:"extra,6"`
}

func NewSearchQuestionReq() *SearchQuestionReq {
	return &SearchQuestionReq{}
}

func (p *SearchQuestionReq) Read(iprot thrift.TProtocol) error {
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
		case 6:
			if err := p.readField6(iprot); err != nil {
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

func (p *SearchQuestionReq) readField1(iprot thrift.TProtocol) error {
	p.Auth = common.NewAuthorize()
	if err := p.Auth.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Auth)
	}
	return nil
}

func (p *SearchQuestionReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Uid = v
	}
	return nil
}

func (p *SearchQuestionReq) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Keyword = v
	}
	return nil
}

func (p *SearchQuestionReq) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 4: %s")
	} else {
		p.Page = v
	}
	return nil
}

func (p *SearchQuestionReq) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 5: %s")
	} else {
		p.Size = v
	}
	return nil
}

func (p *SearchQuestionReq) readField6(iprot thrift.TProtocol) error {
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

func (p *SearchQuestionReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search_question_req"); err != nil {
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
	if err := p.writeField6(oprot); err != nil {
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

func (p *SearchQuestionReq) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SearchQuestionReq) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *SearchQuestionReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("keyword", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:keyword: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Keyword)); err != nil {
		return fmt.Errorf("%T.keyword (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:keyword: %s", p, err)
	}
	return err
}

func (p *SearchQuestionReq) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("page", thrift.I32, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:page: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Page)); err != nil {
		return fmt.Errorf("%T.page (4) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:page: %s", p, err)
	}
	return err
}

func (p *SearchQuestionReq) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("size", thrift.I32, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:size: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Size)); err != nil {
		return fmt.Errorf("%T.size (5) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:size: %s", p, err)
	}
	return err
}

func (p *SearchQuestionReq) writeField6(oprot thrift.TProtocol) (err error) {
	if p.Extra != nil {
		if err := oprot.WriteFieldBegin("extra", thrift.MAP, 6); err != nil {
			return fmt.Errorf("%T write field begin error 6:extra: %s", p, err)
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
			return fmt.Errorf("%T write field end error 6:extra: %s", p, err)
		}
	}
	return err
}

func (p *SearchQuestionReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SearchQuestionReq(%+v)", *p)
}

type SearchQuestionRes struct {
	Rescode   rescode.Code             `thrift:"rescode,1"`
	Resmsg    string                   `thrift:"resmsg,2"`
	Timestamp float64                  `thrift:"timestamp,3"`
	Total     *common.SearchHitsTotal  `thrift:"total,4"`
	Hits      []*common.SearchHitsItem `thrift:"hits,5"`
	Extra     map[string]string        `thrift:"extra,6"`
}

func NewSearchQuestionRes() *SearchQuestionRes {
	return &SearchQuestionRes{
		Rescode: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *SearchQuestionRes) IsSetRescode() bool {
	return int64(p.Rescode) != math.MinInt32-1
}

func (p *SearchQuestionRes) Read(iprot thrift.TProtocol) error {
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
		case 6:
			if err := p.readField6(iprot); err != nil {
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

func (p *SearchQuestionRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Rescode = rescode.Code(v)
	}
	return nil
}

func (p *SearchQuestionRes) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Resmsg = v
	}
	return nil
}

func (p *SearchQuestionRes) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Timestamp = v
	}
	return nil
}

func (p *SearchQuestionRes) readField4(iprot thrift.TProtocol) error {
	p.Total = common.NewSearchHitsTotal()
	if err := p.Total.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Total)
	}
	return nil
}

func (p *SearchQuestionRes) readField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Hits = make([]*common.SearchHitsItem, 0, size)
	for i := 0; i < size; i++ {
		_elem2 := common.NewSearchHitsItem()
		if err := _elem2.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem2)
		}
		p.Hits = append(p.Hits, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *SearchQuestionRes) readField6(iprot thrift.TProtocol) error {
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

func (p *SearchQuestionRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search_question_res"); err != nil {
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
	if err := p.writeField6(oprot); err != nil {
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

func (p *SearchQuestionRes) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SearchQuestionRes) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *SearchQuestionRes) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timestamp", thrift.DOUBLE, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:timestamp: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.Timestamp)); err != nil {
		return fmt.Errorf("%T.timestamp (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:timestamp: %s", p, err)
	}
	return err
}

func (p *SearchQuestionRes) writeField4(oprot thrift.TProtocol) (err error) {
	if p.Total != nil {
		if err := oprot.WriteFieldBegin("total", thrift.STRUCT, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:total: %s", p, err)
		}
		if err := p.Total.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Total)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:total: %s", p, err)
		}
	}
	return err
}

func (p *SearchQuestionRes) writeField5(oprot thrift.TProtocol) (err error) {
	if p.Hits != nil {
		if err := oprot.WriteFieldBegin("hits", thrift.LIST, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:hits: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Hits)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Hits {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:hits: %s", p, err)
		}
	}
	return err
}

func (p *SearchQuestionRes) writeField6(oprot thrift.TProtocol) (err error) {
	if p.Extra != nil {
		if err := oprot.WriteFieldBegin("extra", thrift.MAP, 6); err != nil {
			return fmt.Errorf("%T write field begin error 6:extra: %s", p, err)
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
			return fmt.Errorf("%T write field end error 6:extra: %s", p, err)
		}
	}
	return err
}

func (p *SearchQuestionRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SearchQuestionRes(%+v)", *p)
}