// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package macipnode

import (
	"bytes"
	"context"
	"fmt"
	"github.com/heegspace/heegproto/common"
	"github.com/heegspace/heegproto/rescode"
	"github.com/heegspace/thrift"
	"reflect"
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
//  - Auth
//  - IP
//  - Extra
type IPToAddressReq struct {
	Auth  *common.Authorize `thrift:"auth,1" db:"auth" json:"auth"`
	IP    int64             `thrift:"ip,2" db:"ip" json:"ip"`
	Extra map[string]string `thrift:"extra,3" db:"extra" json:"extra"`
}

func NewIPToAddressReq() *IPToAddressReq {
	return &IPToAddressReq{}
}

var IPToAddressReq_Auth_DEFAULT *common.Authorize

func (p *IPToAddressReq) GetAuth() *common.Authorize {
	if !p.IsSetAuth() {
		return IPToAddressReq_Auth_DEFAULT
	}
	return p.Auth
}

func (p *IPToAddressReq) GetIP() int64 {
	return p.IP
}

func (p *IPToAddressReq) GetExtra() map[string]string {
	return p.Extra
}
func (p *IPToAddressReq) IsSetAuth() bool {
	return p.Auth != nil
}

func (p *IPToAddressReq) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.MAP {
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

func (p *IPToAddressReq) ReadField1(iprot thrift.TProtocol) error {
	p.Auth = &common.Authorize{}
	if err := p.Auth.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Auth), err)
	}
	return nil
}

func (p *IPToAddressReq) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *IPToAddressReq) ReadField3(iprot thrift.TProtocol) error {
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

func (p *IPToAddressReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ip_to_address_req"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
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

func (p *IPToAddressReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("auth", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:auth: ", p), err)
	}
	if err := p.Auth.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Auth), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:auth: ", p), err)
	}
	return err
}

func (p *IPToAddressReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ip: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ip: ", p), err)
	}
	return err
}

func (p *IPToAddressReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:extra: ", p), err)
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
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:extra: ", p), err)
	}
	return err
}

func (p *IPToAddressReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IPToAddressReq(%+v)", *p)
}

// Attributes:
//  - IP
//  - Country
//  - Province
//  - City
//  - Organization
//  - Isp
//  - CountryCode
type AddressItem struct {
	IP           int32  `thrift:"ip,1" db:"ip" json:"ip"`
	Country      string `thrift:"country,2" db:"country" json:"country"`
	Province     string `thrift:"province,3" db:"province" json:"province"`
	City         string `thrift:"city,4" db:"city" json:"city"`
	Organization string `thrift:"organization,5" db:"organization" json:"organization"`
	Isp          string `thrift:"isp,6" db:"isp" json:"isp"`
	CountryCode  string `thrift:"country_code,7" db:"country_code" json:"country_code"`
}

func NewAddressItem() *AddressItem {
	return &AddressItem{}
}

func (p *AddressItem) GetIP() int32 {
	return p.IP
}

func (p *AddressItem) GetCountry() string {
	return p.Country
}

func (p *AddressItem) GetProvince() string {
	return p.Province
}

func (p *AddressItem) GetCity() string {
	return p.City
}

func (p *AddressItem) GetOrganization() string {
	return p.Organization
}

func (p *AddressItem) GetIsp() string {
	return p.Isp
}

func (p *AddressItem) GetCountryCode() string {
	return p.CountryCode
}
func (p *AddressItem) Read(iprot thrift.TProtocol) error {
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
			if fieldTypeId == thrift.STRING {
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

func (p *AddressItem) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *AddressItem) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Country = v
	}
	return nil
}

func (p *AddressItem) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Province = v
	}
	return nil
}

func (p *AddressItem) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.City = v
	}
	return nil
}

func (p *AddressItem) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Organization = v
	}
	return nil
}

func (p *AddressItem) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Isp = v
	}
	return nil
}

func (p *AddressItem) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *AddressItem) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("address_item"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
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
		if err := p.writeField7(oprot); err != nil {
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

func (p *AddressItem) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ip: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ip: ", p), err)
	}
	return err
}

func (p *AddressItem) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("country", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:country: ", p), err)
	}
	if err := oprot.WriteString(string(p.Country)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.country (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:country: ", p), err)
	}
	return err
}

func (p *AddressItem) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("province", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:province: ", p), err)
	}
	if err := oprot.WriteString(string(p.Province)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.province (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:province: ", p), err)
	}
	return err
}

func (p *AddressItem) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("city", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:city: ", p), err)
	}
	if err := oprot.WriteString(string(p.City)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.city (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:city: ", p), err)
	}
	return err
}

func (p *AddressItem) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("organization", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:organization: ", p), err)
	}
	if err := oprot.WriteString(string(p.Organization)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.organization (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:organization: ", p), err)
	}
	return err
}

func (p *AddressItem) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("isp", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:isp: ", p), err)
	}
	if err := oprot.WriteString(string(p.Isp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.isp (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:isp: ", p), err)
	}
	return err
}

func (p *AddressItem) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("country_code", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:country_code: ", p), err)
	}
	if err := oprot.WriteString(string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.country_code (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:country_code: ", p), err)
	}
	return err
}

func (p *AddressItem) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddressItem(%+v)", *p)
}

// Attributes:
//  - Rescode
//  - Resmsg
//  - Address
//  - Extra
type IPToAddressRes struct {
	Rescode rescode.Code      `thrift:"rescode,1" db:"rescode" json:"rescode"`
	Resmsg  string            `thrift:"resmsg,2" db:"resmsg" json:"resmsg"`
	Address *AddressItem      `thrift:"address,3" db:"address" json:"address"`
	Extra   map[string]string `thrift:"extra,4" db:"extra" json:"extra"`
}

func NewIPToAddressRes() *IPToAddressRes {
	return &IPToAddressRes{}
}

func (p *IPToAddressRes) GetRescode() rescode.Code {
	return p.Rescode
}

func (p *IPToAddressRes) GetResmsg() string {
	return p.Resmsg
}

var IPToAddressRes_Address_DEFAULT *AddressItem

func (p *IPToAddressRes) GetAddress() *AddressItem {
	if !p.IsSetAddress() {
		return IPToAddressRes_Address_DEFAULT
	}
	return p.Address
}

func (p *IPToAddressRes) GetExtra() map[string]string {
	return p.Extra
}
func (p *IPToAddressRes) IsSetAddress() bool {
	return p.Address != nil
}

func (p *IPToAddressRes) Read(iprot thrift.TProtocol) error {
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
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.MAP {
				if err := p.ReadField4(iprot); err != nil {
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

func (p *IPToAddressRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := rescode.Code(v)
		p.Rescode = temp
	}
	return nil
}

func (p *IPToAddressRes) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Resmsg = v
	}
	return nil
}

func (p *IPToAddressRes) ReadField3(iprot thrift.TProtocol) error {
	p.Address = &AddressItem{}
	if err := p.Address.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Address), err)
	}
	return nil
}

func (p *IPToAddressRes) ReadField4(iprot thrift.TProtocol) error {
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

func (p *IPToAddressRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ip_to_address_res"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
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
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IPToAddressRes) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *IPToAddressRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("resmsg", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:resmsg: ", p), err)
	}
	if err := oprot.WriteString(string(p.Resmsg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.resmsg (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:resmsg: ", p), err)
	}
	return err
}

func (p *IPToAddressRes) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("address", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:address: ", p), err)
	}
	if err := p.Address.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Address), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:address: ", p), err)
	}
	return err
}

func (p *IPToAddressRes) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("extra", thrift.MAP, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:extra: ", p), err)
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
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:extra: ", p), err)
	}
	return err
}

func (p *IPToAddressRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IPToAddressRes(%+v)", *p)
}

type MacipnodeService interface {
	// Parameters:
	//  - Req
	IPToAddress(ctx context.Context, req *IPToAddressReq) (r *IPToAddressRes, err error)
}

type MacipnodeServiceClient struct {
	c thrift.TClient
}

func NewMacipnodeServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MacipnodeServiceClient {
	return &MacipnodeServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewMacipnodeServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MacipnodeServiceClient {
	return &MacipnodeServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewMacipnodeServiceClient(c thrift.TClient) *MacipnodeServiceClient {
	return &MacipnodeServiceClient{
		c: c,
	}
}

func (p *MacipnodeServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Req
func (p *MacipnodeServiceClient) IPToAddress(ctx context.Context, req *IPToAddressReq) (r *IPToAddressRes, err error) {
	var _args4 MacipnodeServiceIPToAddressArgs
	_args4.Req = req
	var _result5 MacipnodeServiceIPToAddressResult
	if err = p.Client_().Call(ctx, "ip_to_address", &_args4, &_result5); err != nil {
		return
	}
	return _result5.GetSuccess(), nil
}

type MacipnodeServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      MacipnodeService
}

func (p *MacipnodeServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *MacipnodeServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *MacipnodeServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewMacipnodeServiceProcessor(handler MacipnodeService) *MacipnodeServiceProcessor {

	self6 := &MacipnodeServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["ip_to_address"] = &macipnodeServiceProcessorIPToAddress{handler: handler}
	return self6
}

func (p *MacipnodeServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x7

}

type macipnodeServiceProcessorIPToAddress struct {
	handler MacipnodeService
}

func (p *macipnodeServiceProcessorIPToAddress) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MacipnodeServiceIPToAddressArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ip_to_address", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MacipnodeServiceIPToAddressResult{}
	var retval *IPToAddressRes
	var err2 error
	if retval, err2 = p.handler.IPToAddress(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ip_to_address: "+err2.Error())
		oprot.WriteMessageBegin("ip_to_address", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ip_to_address", thrift.REPLY, seqId); err2 != nil {
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
type MacipnodeServiceIPToAddressArgs struct {
	Req *IPToAddressReq `thrift:"req,1" db:"req" json:"req"`
}

func NewMacipnodeServiceIPToAddressArgs() *MacipnodeServiceIPToAddressArgs {
	return &MacipnodeServiceIPToAddressArgs{}
}

var MacipnodeServiceIPToAddressArgs_Req_DEFAULT *IPToAddressReq

func (p *MacipnodeServiceIPToAddressArgs) GetReq() *IPToAddressReq {
	if !p.IsSetReq() {
		return MacipnodeServiceIPToAddressArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *MacipnodeServiceIPToAddressArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MacipnodeServiceIPToAddressArgs) Read(iprot thrift.TProtocol) error {
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

func (p *MacipnodeServiceIPToAddressArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = &IPToAddressReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *MacipnodeServiceIPToAddressArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ip_to_address_args"); err != nil {
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

func (p *MacipnodeServiceIPToAddressArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *MacipnodeServiceIPToAddressArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MacipnodeServiceIPToAddressArgs(%+v)", *p)
}

// Attributes:
//  - Success
type MacipnodeServiceIPToAddressResult struct {
	Success *IPToAddressRes `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewMacipnodeServiceIPToAddressResult() *MacipnodeServiceIPToAddressResult {
	return &MacipnodeServiceIPToAddressResult{}
}

var MacipnodeServiceIPToAddressResult_Success_DEFAULT *IPToAddressRes

func (p *MacipnodeServiceIPToAddressResult) GetSuccess() *IPToAddressRes {
	if !p.IsSetSuccess() {
		return MacipnodeServiceIPToAddressResult_Success_DEFAULT
	}
	return p.Success
}
func (p *MacipnodeServiceIPToAddressResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MacipnodeServiceIPToAddressResult) Read(iprot thrift.TProtocol) error {
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

func (p *MacipnodeServiceIPToAddressResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &IPToAddressRes{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *MacipnodeServiceIPToAddressResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ip_to_address_result"); err != nil {
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

func (p *MacipnodeServiceIPToAddressResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *MacipnodeServiceIPToAddressResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MacipnodeServiceIPToAddressResult(%+v)", *p)
}
