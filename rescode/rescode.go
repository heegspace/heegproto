// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rescode

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/heegspace/thrift"
	"reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type Code int64

const (
	Code_SUCCESS                  Code = 0
	Code_ERROR                    Code = 1
	Code_DB_ERROR                 Code = 2
	Code_AUTH_ERR                 Code = 3
	Code_MOBILE_ERR               Code = 4
	Code_EMAIL_ERR                Code = 5
	Code_PARAM_ERR                Code = 99
	Code_EXISTS                   Code = 100
	Code_IS_SELF                  Code = 101
	Code_NOT_EXISTS               Code = 400
	Code_NOT_DATA                 Code = 404
	Code_SEND_CODE_ERR            Code = 10000
	Code_CODE_ERROR               Code = 10001
	Code_CODE_EXPIRE              Code = 10002
	Code_CODE_RATE                Code = 10003
	Code_CODE_LIMIT               Code = 10004
	Code_CODE_TYPE_ERR            Code = 10005
	Code_JSON_MAR_ERR             Code = 10006
	Code_JSON_UNMAR_ERR           Code = 10007
	Code_CODE_NODE_ERROR          Code = 9000
	Code_CODE_NODE_NOINIT         Code = 9001
	Code_DATA_NODE_ERROR          Code = 9002
	Code_DATA_NODE_NOINIT         Code = 9003
	Code_DARTY_NODE_ERROR         Code = 9004
	Code_DARTY_NODE_NOINIT        Code = 9005
	Code_WECHAT_LOGIN_ERROR       Code = 30001
	Code_WECHAT_TOKEN_TIMEOUT     Code = 30002
	Code_WECHAT_REFRESH_ERROR     Code = 30003
	Code_WECHAT_USERINFO_ERR      Code = 30004
	Code_ALIPAY_LOGIN_ERROR       Code = 30031
	Code_ALIPAY_TOKEN_TIMEOUT     Code = 30032
	Code_ALIPAY_REFRESH_ERROR     Code = 30033
	Code_ALIPAY_USERINFO_ERR      Code = 30034
	Code_BAIDU_ACCESS_TOKEN_ERROR Code = 30061
	Code_BAIDU_ENTITY_ERROR       Code = 30062
	Code_BAIDU_IDIDENT_ERROR      Code = 30063
	Code_BAIDU_IDIDENT_FAIL       Code = 30084
)

func (p Code) String() string {
	switch p {
	case Code_SUCCESS:
		return "SUCCESS"
	case Code_ERROR:
		return "ERROR"
	case Code_DB_ERROR:
		return "DB_ERROR"
	case Code_AUTH_ERR:
		return "AUTH_ERR"
	case Code_MOBILE_ERR:
		return "MOBILE_ERR"
	case Code_EMAIL_ERR:
		return "EMAIL_ERR"
	case Code_PARAM_ERR:
		return "PARAM_ERR"
	case Code_EXISTS:
		return "EXISTS"
	case Code_IS_SELF:
		return "IS_SELF"
	case Code_NOT_EXISTS:
		return "NOT_EXISTS"
	case Code_NOT_DATA:
		return "NOT_DATA"
	case Code_SEND_CODE_ERR:
		return "SEND_CODE_ERR"
	case Code_CODE_ERROR:
		return "CODE_ERROR"
	case Code_CODE_EXPIRE:
		return "CODE_EXPIRE"
	case Code_CODE_RATE:
		return "CODE_RATE"
	case Code_CODE_LIMIT:
		return "CODE_LIMIT"
	case Code_CODE_TYPE_ERR:
		return "CODE_TYPE_ERR"
	case Code_JSON_MAR_ERR:
		return "JSON_MAR_ERR"
	case Code_JSON_UNMAR_ERR:
		return "JSON_UNMAR_ERR"
	case Code_CODE_NODE_ERROR:
		return "CODE_NODE_ERROR"
	case Code_CODE_NODE_NOINIT:
		return "CODE_NODE_NOINIT"
	case Code_DATA_NODE_ERROR:
		return "DATA_NODE_ERROR"
	case Code_DATA_NODE_NOINIT:
		return "DATA_NODE_NOINIT"
	case Code_DARTY_NODE_ERROR:
		return "DARTY_NODE_ERROR"
	case Code_DARTY_NODE_NOINIT:
		return "DARTY_NODE_NOINIT"
	case Code_WECHAT_LOGIN_ERROR:
		return "WECHAT_LOGIN_ERROR"
	case Code_WECHAT_TOKEN_TIMEOUT:
		return "WECHAT_TOKEN_TIMEOUT"
	case Code_WECHAT_REFRESH_ERROR:
		return "WECHAT_REFRESH_ERROR"
	case Code_WECHAT_USERINFO_ERR:
		return "WECHAT_USERINFO_ERR"
	case Code_ALIPAY_LOGIN_ERROR:
		return "ALIPAY_LOGIN_ERROR"
	case Code_ALIPAY_TOKEN_TIMEOUT:
		return "ALIPAY_TOKEN_TIMEOUT"
	case Code_ALIPAY_REFRESH_ERROR:
		return "ALIPAY_REFRESH_ERROR"
	case Code_ALIPAY_USERINFO_ERR:
		return "ALIPAY_USERINFO_ERR"
	case Code_BAIDU_ACCESS_TOKEN_ERROR:
		return "BAIDU_ACCESS_TOKEN_ERROR"
	case Code_BAIDU_ENTITY_ERROR:
		return "BAIDU_ENTITY_ERROR"
	case Code_BAIDU_IDIDENT_ERROR:
		return "BAIDU_IDIDENT_ERROR"
	case Code_BAIDU_IDIDENT_FAIL:
		return "BAIDU_IDIDENT_FAIL"
	}
	return "<UNSET>"
}

func CodeFromString(s string) (Code, error) {
	switch s {
	case "SUCCESS":
		return Code_SUCCESS, nil
	case "ERROR":
		return Code_ERROR, nil
	case "DB_ERROR":
		return Code_DB_ERROR, nil
	case "AUTH_ERR":
		return Code_AUTH_ERR, nil
	case "MOBILE_ERR":
		return Code_MOBILE_ERR, nil
	case "EMAIL_ERR":
		return Code_EMAIL_ERR, nil
	case "PARAM_ERR":
		return Code_PARAM_ERR, nil
	case "EXISTS":
		return Code_EXISTS, nil
	case "IS_SELF":
		return Code_IS_SELF, nil
	case "NOT_EXISTS":
		return Code_NOT_EXISTS, nil
	case "NOT_DATA":
		return Code_NOT_DATA, nil
	case "SEND_CODE_ERR":
		return Code_SEND_CODE_ERR, nil
	case "CODE_ERROR":
		return Code_CODE_ERROR, nil
	case "CODE_EXPIRE":
		return Code_CODE_EXPIRE, nil
	case "CODE_RATE":
		return Code_CODE_RATE, nil
	case "CODE_LIMIT":
		return Code_CODE_LIMIT, nil
	case "CODE_TYPE_ERR":
		return Code_CODE_TYPE_ERR, nil
	case "JSON_MAR_ERR":
		return Code_JSON_MAR_ERR, nil
	case "JSON_UNMAR_ERR":
		return Code_JSON_UNMAR_ERR, nil
	case "CODE_NODE_ERROR":
		return Code_CODE_NODE_ERROR, nil
	case "CODE_NODE_NOINIT":
		return Code_CODE_NODE_NOINIT, nil
	case "DATA_NODE_ERROR":
		return Code_DATA_NODE_ERROR, nil
	case "DATA_NODE_NOINIT":
		return Code_DATA_NODE_NOINIT, nil
	case "DARTY_NODE_ERROR":
		return Code_DARTY_NODE_ERROR, nil
	case "DARTY_NODE_NOINIT":
		return Code_DARTY_NODE_NOINIT, nil
	case "WECHAT_LOGIN_ERROR":
		return Code_WECHAT_LOGIN_ERROR, nil
	case "WECHAT_TOKEN_TIMEOUT":
		return Code_WECHAT_TOKEN_TIMEOUT, nil
	case "WECHAT_REFRESH_ERROR":
		return Code_WECHAT_REFRESH_ERROR, nil
	case "WECHAT_USERINFO_ERR":
		return Code_WECHAT_USERINFO_ERR, nil
	case "ALIPAY_LOGIN_ERROR":
		return Code_ALIPAY_LOGIN_ERROR, nil
	case "ALIPAY_TOKEN_TIMEOUT":
		return Code_ALIPAY_TOKEN_TIMEOUT, nil
	case "ALIPAY_REFRESH_ERROR":
		return Code_ALIPAY_REFRESH_ERROR, nil
	case "ALIPAY_USERINFO_ERR":
		return Code_ALIPAY_USERINFO_ERR, nil
	case "BAIDU_ACCESS_TOKEN_ERROR":
		return Code_BAIDU_ACCESS_TOKEN_ERROR, nil
	case "BAIDU_ENTITY_ERROR":
		return Code_BAIDU_ENTITY_ERROR, nil
	case "BAIDU_IDIDENT_ERROR":
		return Code_BAIDU_IDIDENT_ERROR, nil
	case "BAIDU_IDIDENT_FAIL":
		return Code_BAIDU_IDIDENT_FAIL, nil
	}
	return Code(0), fmt.Errorf("not a valid Code string")
}

func CodePtr(v Code) *Code { return &v }

func (p Code) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Code) UnmarshalText(text []byte) error {
	q, err := CodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *Code) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = Code(v)
	return nil
}

func (p *Code) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}
